package entities

// Save .
func Save(u *UserInfo) error {
	//使用事务处理
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	//插入数据
	_, err = session.Table("userinfo").Insert(u)
	//出错则回滚
	if err != nil {
		session.Rollback()
		return err
	}
	//提交修改
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

// FindAll .
func FindAll() []UserInfo {
	everyone := make([]UserInfo, 0)
	err := engine.Table("userinfo").Find(&everyone)
	checkErr(err)
	return everyone
}

// FindByID .
func FindByID(id int) *UserInfo {
	user := new(UserInfo)
	_, err := engine.Table("userinfo").Where("uid=?", id).Get(user)
	checkErr(err)
	return user
}
