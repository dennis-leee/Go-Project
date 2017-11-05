package entity

import (
	"encoding/json"
	"io"
	"log"
)

type User struct {
	Username string
	Password string
	Mail     string
	Phone    string
}

type Users map[string]*User

func NewUsers() Users {
	return make(Users)
}

func (users Users) Lookup(username string) *User {
	return users[username]
}

func (users Users) Size() int {
	return len(users)
}

func (users Users) Add(user *User) bool {
	if user == nil {
		return false
	}
	u := users.Lookup(user.Username)
	if u != nil {
		return false
	}
	users[user.Username] = user
	return true
}

func (users Users) Remove(user *User) *User {
	if user == nil {
		return nil
	}
	u := users.Lookup(user.Username)
	if u == nil {
		return nil
	}
	delete(users, user.Username)
	return u
}

func (users Users) Slice() []*User {
	s := make([]*User, 0, len(users))
	for _, u := range users {
		s = append(s, u)
	}
	return s
}

func (users Users) Serialize(w io.Writer) {
	encoder := json.NewEncoder(w)
	for _, u := range users {
		encoder.Encode(u)
	}
}

func DeserializeUser(r io.Reader) (Users, error) {
	decoder := json.NewDecoder(r)
	users := make(Users)
	for {
		u := new(User)
		if err := decoder.Decode(u); err == io.EOF {
			return users, nil
		} else if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		users.Add(u)
	}
}
