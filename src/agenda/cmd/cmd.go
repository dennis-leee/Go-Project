package cmd

import (
	"entity"
	"fmt"
	"log"
	"os"
	"strings"
)

func printWrongLoginState(action string, required bool) int {
	var s string
	if required {
		s = "login"
	} else {
		s = "logout"
	}
	util.PrintfErr("Action %s requires an %s state\n", action, s)
	return int(err.WrongLoginState)
}
  //当要打印不存在的会议时，返回错误 
func printMeetingDoesntExist(title string) int {
	util.PrintfErr("meeting doesn't exist: %s\n", title)
	return int(err.NoSuchMeeting)
}
   //判断会议发起人 
func printNotAHost(title string, user string) int {
	util.PrintfErr("meeting '%s' is not hosted by '%s'", title, user)
	return int(err.NotEnoughPrivilege)
}
   //当要打印的用户不存在时，返回错误 
func printUserDoesntExist(username string) int {
	util.PrintfErr("user doesn't exist: %s\n", username)
	return int(err.NoSuchUser)
}
    //判断要打印的时间的格式是否错误 
func printInvalidTimeFormat(time string) int {
	util.PrintfErr("invalid time format: %s\n", time)
	return int(err.InvalidTime)
}
   //用户登录 
func loadLogin(us entity.Users) *entity.User {
	u, e := model.LoadLogin(us)
	if e == err.OK {
		return u
	}
	log.Fatalf("something wrong with login file, error: %d", int(e))
	return nil
}

  //用户注册
func Register(user, pass, mail, phone string) int {
	users := model.LoadUsers()
	passhash := util.PrettyHash(pass)
	log.Printf("password hash for '%s': %s\n", pass, passhash)
	if !users.Add(&entity.User{
		Username: user,
		Password: passhash,
		Mail:     mail,
		Phone:    phone,
	}) {
		util.PrintfErr("there's another user with username %s\n", user)
		return int(err.DuplicateUser)
	}
	model.StoreUser(users)
	return 0
}

// 登陆的命令
func Login(user, pass string) int {
	users := model.LoadUsers()
	if loadLogin(users) != nil {
		return printWrongLoginState("login", false)
	}
	if model.Login(users, user, pass) != err.OK {
		util.PrintlnErr("Authentication Fail")
		return int(err.AuthenticateFail)
	}
	return 0
}

//退出的命令
func Logout() int {
	if model.Logout() {
		fmt.Printf("logout success\n")
	} else {
		return printWrongLoginState("logout", true)
	}
	return 0
}

// 打印所有的用户 
func ShowUsers() int {
	users := model.LoadUsers()
	if loadLogin(users) == nil {
		return printWrongLoginState("ShowUser", true)
	}
	fmt.Println("Username Email Phone")
	for _, u := range users {
		fmt.Printf("'%s' '%s' '%s'\n", u.Username, u.Mail, u.Phone)
	}
	return 0
}

// 删除当前的用户以及相关的内容 
func DeleteUser() int {
	users := model.LoadUsers()
	u := loadLogin(users)
	if u == nil {
		return printWrongLoginState("DeleteUser", true)
	}
	if users.Remove(u) == nil {
		log.Fatalln("Login user cannot be removed !?")
	}
	meetings := model.LoadMeetings(users)
	meetings.RemoveAll(u)
	model.StoreMeeting(meetings)
	model.StoreUser(users)
	return 0
}
//判断现在能不能主持会议 
func hostMeeting(ms *entity.Meetings, m *entity.Meeting) int {
	e := ms.Host(m)
	switch e {
	case err.InvalidTime:
		util.PrintlnErr("meeting should end later than start")
	case err.DuplicateMeeting:
		util.PrintfErr("there's another meeting with title: %s\n", m.Title)
	case err.TimeConflict:
		util.PrintlnErr("there are time conflict of some participants")
	case err.OK:
		model.StoreMeeting(ms)
		fmt.Println("meeting hosted")
	}
	return int(e)
}

//判断在特定时间能不能开始会议 
func HostMeeting(title string, parts []string, start, end string) int {
	users := model.LoadUsers()
	host := loadLogin(users)
	if host == nil {
		return printWrongLoginState("HostMeeting", true)
	}
	s, es := util.YMDParse(start)
	if es != nil {
		return printInvalidTimeFormat(start)
	}
	e, ee := util.YMDParse(end)
	if ee != nil {
		return printInvalidTimeFormat(end)
	}
	participants := entity.NewUsers()
	for _, p := range parts {
		u := users.Lookup(p)
		if u == nil {
			return printUserDoesntExist(p)
		}
		participants.Add(u)
	}
	meeting := &entity.Meeting{
		Title:        title,
		Host:         host,
		Participants: participants,
		Start:        s,
		End:          e,
	}
	meetings := model.LoadMeetings(users)
	return hostMeeting(meetings, meeting)
}

// 取消该用户主持的某个会议 
func CancelMeeting(title string) int {
	users := model.LoadUsers()
	host := loadLogin(users)
	if host == nil {
		return printWrongLoginState("CancelMeeting", true)
	}
	meetings := model.LoadMeetings(users)
	meeting := meetings.Lookup(title)
	if meeting == nil {
		return printMeetingDoesntExist(title)
	}
	if meeting.Host.Username != host.Username {
		return printNotAHost(title, host.Username)
	}
	meetings.Cancel(title)
	model.StoreMeeting(meetings)
	return 0
}



// 添加某个主持的会议的成员 
func AddParticipant(title string, username string) int {
	users := model.LoadUsers()
	meetings := model.LoadMeetings(users)
	host := loadLogin(users)
	if host == nil {
		return printWrongLoginState("AddParticipant", true)
	}
	meeting := meetings.Lookup(title)
	if meeting == nil {
		return printMeetingDoesntExist(title)
	}
	if meeting.Host.Username != host.Username {
		return printNotAHost(title, host.Username)
	}
	addedUser := users.Lookup(username)
	if addedUser == nil {
		return printUserDoesntExist(username)
	}
	e := meetings.Add(title, addedUser)
	switch e {
	case err.NoSuchMeeting:
		log.Fatalln("meeting should be checked before??")
	case err.DuplicateUser:
		fmt.Fprintf(
			os.Stderr, "user '%s' is already a participant of meeting '%s'",
			username, title)
	case err.TimeConflict:
		fmt.Fprintf(os.Stderr,
			"there is time conflict of user '%s' and meeting '%s'",
			username, title)
	case err.OK:
		model.StoreMeeting(meetings)
	default:
		log.Fatalf("unexpected error: %d\n", int(e))
	}
	return int(e)
}
//打印会议信息 
func printMeeting(m *entity.Meeting) {
	parts := make([]string, 0)
	for _, u := range m.Participants {
		parts = append(parts, u.Username)
	}
	fmt.Printf("title: %s\n\thost: %s\n\ttime: %s to %s\n\tparticipants: %s\n",
		m.Title, m.Host.Username,
		util.YMDFormat(m.Start), util.YMDFormat(m.End),
		strings.Join(parts, ", "))
}

// 移除某个主持的会议的成员 
func RemoveParticipant(title string, username string) int {
	users := model.LoadUsers()
	host := loadLogin(users)
	if host == nil {
		return printWrongLoginState("RemoveParticipant", true)
	}
	meetings := model.LoadMeetings(users)
	meeting := meetings.Lookup(title)
	if meeting == nil {
		return printMeetingDoesntExist(title)
	}
	if meeting.Host.Username != host.Username {
		return printNotAHost(title, host.Username)
	}
	user := users.Lookup(username)
	if user == nil {
		return printUserDoesntExist(username)
	}
	e := meetings.Remove(title, user)
	if e != err.OK {
		log.Fatalf("error should be all checked! %d", e)
	}
	model.StoreMeeting(meetings)
	return 0
}

// 退出参加的某个会议 
func QuitMeeting(title string) int {
	users := model.LoadUsers()
	meetings := model.LoadMeetings(users)
	user := loadLogin(users)
	e := meetings.Remove(title, user)
	switch e {
	case err.NoSuchMeeting:
		printMeetingDoesntExist(title)
	case err.NoSuchUser:
		fmt.Fprintf(
			os.Stderr, "user '%s' is not a participant of meeting '%s'\n",
			user.Username, title)
	case err.OK:
		model.StoreMeeting(meetings)
	}
	return int(e)
}

// 清除当前用户的所有会议 
func ClearMeetings() int {
	users := model.LoadUsers()
	host := loadLogin(users)
	if host == nil {
		return printWrongLoginState("ClearMeeting", true)
	}
	meetings := model.LoadMeetings(users)
	meetings.CancelAll(host)
	model.StoreMeeting(meetings)
	return 0
}


// 按照也定时间段来查询回忆 
func QueryMeeting(start, end string) int {
	users := model.LoadUsers()
	user := loadLogin(users)
	if user == nil {
		return printWrongLoginState("QueryMeeting", true)
	}
	s, es := util.YMDParse(start)
	if es != nil {
		return printInvalidTimeFormat(start)
	}
	e, ee := util.YMDParse(end)
	if ee != nil {
		return printInvalidTimeFormat(end)
	}
	if e.Before(s) {
		util.PrintfErr("invalid interval %s - %s", start, end)
		return int(err.InvalidTime)
	}
	meetings := model.LoadMeetings(users)
	for _, m := range meetings.Related(user.Username) {
		if e.After(m.Start) && s.Before(m.End) {
			printMeeting(m)
		}
	}
	return 0
}
