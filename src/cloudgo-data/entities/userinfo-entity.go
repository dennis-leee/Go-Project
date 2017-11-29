package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int `xorm:"pk" xorm:"autoincr"` //主键,递增
	UserName   string
	DepartName string
	CreatedAt  time.Time `xorm:"created"` //自动更新时间
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	return &u
}
