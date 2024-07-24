package mysql

import (
	"time"

	"github.com/nepu-SidneyYu/blog-grpc/internal/model"
)

type User struct {
	// 定义User结构体
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserByName(name string) (*model.UserAuth, error) {
	var userauth model.UserAuth
	tx := _db.Model(&model.UserAuth{}).Where("username = ?", name).First(&userauth)
	// 根据id查询用户信息
	// 返回一个User结构体和一个错误信息
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &userauth, nil
}
func (u *User) GetUserByPhone(phone string) (*model.UserAuth, error) {
	var userauth model.UserAuth
	tx := _db.Model(&model.UserAuth{}).Where("phone = ?", phone).First(&userauth)
	// 根据id查询用户信息
	// 返回一个User结构体和一个错误信息
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &userauth, nil
}
func (u *User) SetUser(phone, password string) error {
	tx := _db.Model(&model.UserAuth{}).Create(&model.UserAuth{Phone: phone, Password: password, CreatedAt: time.Now().UnixMilli()})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *User) SetUserName(phone, name string) error {
	tx := _db.Model(&model.UserAuth{}).Where("phone = ?", phone).Updates(map[string]interface{}{
		"username":   name,
		"updated_at": time.Now().UnixMilli()})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *User) BindEmail(phone, email string) error {
	tx := _db.Model(&model.UserAuth{}).Where("phone = ?", phone).Updates(map[string]interface{}{"email": email, "updated_at": time.Now().UnixMilli()})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func (u *User)GetUserByEmail(email string) (*model.UserAuth, error){
	var userauth model.UserAuth
	tx := _db.Model(&model.UserAuth{}).Where("email = ?", email).First(&userauth)
	// 根据id查询用户信息
	// 返回一个User结构体和一个错误信息
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &userauth, nil
}
