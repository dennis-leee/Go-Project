package cmd

import (
	"log"
)

type Err int

const (
	// OK表示成功 
	OK Err = 0
	//时间冲突错误发生在加入某个会议室与其他会议时间冲突 
	TimeConflict Err = 1
	// NoSuchMeeting表示该会议不存在 
	NoSuchMeeting Err = 2
	// NoSuchUser 表示该用户不存在 
	NoSuchUser Err = 3
	// DuplicateMeeting 表示要创建一个已经存在的会议 
	DuplicateMeeting Err = 4
	// DuplicateUser表示想注册一个已经存在的用户 
	DuplicateUser Err = 5
	// InvalidTime 表示时间存在错误 
	InvalidTime Err = 6
	// NoSuchFile 表示操作要找的文件不存在 
	NoSuchFile Err = 7
	// InconsistentState 表示状态有误 
	InconsistentState Err = 8
	// AuthenticateFail表示用户与密码不匹配 
	AuthenticateFail Err = 9
	// WrongLoginState 表示登录时状态有误 
	WrongLoginState Err = 10
	// NotEnoughPrivilege 表示进行了同样的操作 
	NotEnoughPrivilege Err = 11
)

func LogFatalIfError(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
