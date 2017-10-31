package cmd

import (
	"encoding/json"
	"entity"
	"log"
	"os"
)

type login struct {
	Username, Password string
}

const rewritePerm = os.O_WRONLY | os.O_CREATE | os.O_TRUNC

func AgendaDir() string {
	home, present := os.LookupEnv("HOME")
	if !present {
		home = "."
	}
	return home + "/.agenda/"
}

func EnsureAgendaDir() {
	os.Mkdir(AgendaDir(), 0755)
}

func LoginFile() string {
	return AgendaDir() + "login"
}

func UserFile() string {
	return AgendaDir() + "user"
}

func MeetingFile() string {
	return AgendaDir() + "meeting"
}

func LogFile() string {
	return AgendaDir() + "log"
}

func openFileRewrite(path string) (*os.File, error) {
	return os.OpenFile(path, rewritePerm, 0644)
}

func LoadUsers() entity.Users {
	file, e := os.Open(UserFile())
	if e != nil {
		log.Printf("There is no user file. %s\n", e.Error())
		return entity.NewUsers()
	}
	us, e := entity.DeserializeUser(file)
	err.LogFatalIfError(e)
	return us
}

func LoadMeetings(users entity.Users) *entity.Meetings {
	file, e := os.Open(MeetingFile())
	if e != nil {
		log.Printf("There is no meeting file. %s\n", e.Error())
		return entity.NewMeetings()
	}
	ms, e := entity.DeserializeMeeting(file, users)
	return ms
}


func loadLoginFile() *login {
	file, e := os.Open(LoginFile())
	if e != nil {
		return nil
	}
	l := new(login)
	e = json.NewDecoder(file).Decode(l)
	if e != nil {
		Logout()
		return nil
	}
	return l
}

func writeLoginFile(user, pass string) {
	file, e := openFileRewrite(LoginFile())
	err.LogFatalIfError(e)
	json.NewEncoder(file).Encode(login{user, pass})
}

func validPassword(u *entity.User, pass string) bool {
	hash := util.PrettyHash(pass)
	log.Printf("pass '%s': '%s', to '%s'", pass, hash, u.Password)
	return hash == u.Password
}

func LoadLogin(users entity.Users) (*entity.User, err.Err) {
	l := loadLoginFile()
	if l == nil {
		log.Println("not found")
		return nil, err.OK
	}
	u := users.Lookup(l.Username)
	if u == nil {
		log.Printf("not found: %s\n", l.Username)
		return nil, err.NoSuchUser
	}
	if !validPassword(u, l.Password) {
		log.Printf("invalid")
		return nil, err.InconsistentState
	}
	log.Printf("login loaded: %s:%s\n", l.Username, l.Password)
	return u, err.OK
}

func Logout() bool {
	return os.Remove(LoginFile()) == nil
}

func Login(users entity.Users, user, pass string) err.Err {
	u := users.Lookup(user)
	if u == nil {
		log.Printf("There is no user named: %s\n", user)
		return err.NoSuchUser
	}
	if validPassword(u, pass) {
		writeLoginFile(user, pass)
		log.Printf("login success")
		return err.OK
	}
	log.Printf("Invalid password")
	return err.AuthenticateFail
}

func StoreUser(users entity.Users) {
	file, e := openFileRewrite(UserFile())
	err.LogFatalIfError(e)
	users.Serialize(file)
}

func StoreMeeting(meetings *entity.Meetings) {
	file, e := openFileRewrite(MeetingFile())
	err.LogFatalIfError(e)
	meetings.Serialize(file)
}

