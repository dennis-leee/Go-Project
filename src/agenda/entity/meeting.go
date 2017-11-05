package entity

import (
	"cmd"
	"encoding/json"
	"io"
	"log"
	"time"
)

type Meeting struct {
	Title string
	Host *User
	Participants Users
	Start time.Time
	End time.Time
}

type Meetings struct {
	meetings map[string]*Meeting
	relation map[string]map[string]*Meeting
}

type meetingSerialized struct {
	Title        string
	Host         string
	Participants []string
	Start        string
	End          string
}

func NewMeetings() *Meetings {
	return &Meetings{
		meetings: make(map[string]*Meeting),
		relation: make(map[string]map[string]*Meeting),
	}
}

func (ms *Meetings) Slice() []*Meeting {
	a := make([]*Meeting, 0, len(ms.meetings))
	for _, m := range ms.meetings {
		a = append(a, m)
	}
	return a
}

func (ms *Meetings) Has(title string) bool {
	return ms.meetings[title] != nil
}

func (ms *Meetings) Related(user string) map[string]*Meeting {
	related := ms.relation[user]
	if related == nil {
		return make(map[string]*Meeting)
	}
	return related
}

func (ms *Meetings) addRelatedMeeting(u *User, m *Meeting) {
	meetings := ms.relation[u.Username]
	if meetings == nil {
		ms.relation[u.Username] = map[string]*Meeting{m.Title: m}
	} else {
		meetings[m.Title] = m
	}
}

func (ms *Meetings) Lookup(title string) *Meeting {
	return ms.meetings[title]
}

func (ms *Meetings) host(m *Meeting) {
	ms.meetings[m.Title] = m
	for _, u := range m.Participants {
		ms.addRelatedMeeting(u, m)
	}
	ms.addRelatedMeeting(m.Host, m)
}

func (ms *Meetings) Host(m *Meeting) err.Err {
	if !m.End.After(m.Start) {
		return err.InvalidTime
	}
	if ms.Has(m.Title) {
		return err.DuplicateMeeting
	}
	for _, u := range m.Participants {
		for _, um := range ms.Related(u.Username) {
			if util.Overlapped(m.Start, m.End, um.Start, um.End) {
				log.Printf(
					"time conflict for '%s': %s to %s\n", u.Username,
					util.YMDFormat(um.Start), util.YMDFormat(um.End))
				return err.TimeConflict
			}
		}
	}
	ms.host(m)
	return err.OK
}

func (ms *Meetings) Serialize(w io.Writer) {
	encoder := json.NewEncoder(w)
	for _, m := range ms.meetings {
		encoder.Encode(toSerialized(m))
	}
}

func toSerialized(m *Meeting) *meetingSerialized {
	parts := make([]string, 0, len(m.Participants))
	for name := range m.Participants {
		parts = append(parts, name)
	}
	return &meetingSerialized{
		Title:        m.Title,
		Host:         m.Host.Username,
		Participants: parts,
		Start:        util.YMDFormat(m.Start),
		End:          util.YMDFormat(m.End),
	}
}

func fromSerialized(m *meetingSerialized, users Users) *Meeting {
	start, _ := util.YMDParse(m.Start)
	end, _ := util.YMDParse(m.End)
	host := users.Lookup(m.Host)
	parts := NewUsers()
	for _, p := range m.Participants {
		u := users.Lookup(p)
		parts.Add(u)
	}
	return &Meeting{
		Host:         host,
		Start:        start,
		End:          end,
		Participants: parts,
		Title:        m.Title,
	}
}

func DeserializeMeeting(r io.Reader, users Users) (*Meetings, error) {
	decoder := json.NewDecoder(r)
	ms := NewMeetings()
	for {
		m := new(meetingSerialized)
		if err := decoder.Decode(m); err == io.EOF {
			return ms, nil
		} else if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		ms.host(fromSerialized(m, users))
	}
}

func (ms *Meetings) cancel(m *Meeting) {
	for _, u := range m.Participants {
		delete(ms.relation[u.Username], m.Title)
	}
	delete(ms.meetings, m.Title)
}

func (ms *Meetings) Cancel(title string) err.Err {
	m := ms.meetings[title]
	if m == nil {
		return err.NoSuchMeeting
	}
	ms.cancel(m)
	return err.OK
}

func (ms *Meetings) CancelAll(host *User) {
	for _, m := range ms.Related(host.Username) {
		if m.Host.Username == host.Username {
			ms.cancel(m)
		}
	}
}

func (ms *Meetings) remove(m *Meeting, user *User) {
	m.Participants.Remove(user)
	if user.Username != m.Host.Username || m.Participants.Size() == 0 {
		delete(ms.relation[user.Username], m.Title)
	}
	if m.Participants.Size() == 0 {
		delete(ms.meetings, m.Title)
	}
}

func (ms *Meetings) Remove(title string, user *User) err.Err {
	m := ms.meetings[title]
	if m == nil {
		return err.NoSuchMeeting
	}
	if m.Participants.Lookup(user.Username) == nil {
		return err.NoSuchUser
	}
	ms.remove(m, user)
	return err.OK
}

func (ms *Meetings) RemoveAll(user *User) err.Err {
	ms.CancelAll(user)
	related := make([]*Meeting, 0)
	for _, m := range ms.Related(user.Username) {
		related = append(related, m)
	}
	for _, m := range related {
		ms.remove(m, user)
	}
	return err.OK
}

func (ms *Meetings) Add(title string, user *User) err.Err {
	m := ms.meetings[title]
	if m == nil {
		return err.NoSuchMeeting
	}
	if m.Participants.Lookup(user.Username) != nil {
		return err.DuplicateUser
	}
	for _, um := range ms.Related(user.Username) {
		if util.Overlapped(um.Start, um.End, m.Start, m.End) {
			return err.TimeConflict
		}
	}
	m.Participants.Add(user)
	ms.addRelatedMeeting(user, m)
	return err.OK
}