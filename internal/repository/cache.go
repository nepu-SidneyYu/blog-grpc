package repository

type CodeCache interface {
	SetPhoneCode(feild string, codekey string, code string, expire int64) error
	GetPhoneCode(feild string, codekey string) (string, error)
	SetEmailCode(feild string, codekey string, code string, expire int64) error
	GetEmailCode(feild string, codekey string) (string, error)
}

type UserNameCache interface {
	SetUserName(name string) error
	IsUserNameExist(name string) bool
}
