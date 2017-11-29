package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

/********全都由Dao层实现***********/
// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	return Save(u)
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	return FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	return FindByID(id)
}
