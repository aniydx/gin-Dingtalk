package models

import (
	"goDingTalk/dao"
)

type User struct {
	Id       int
	Name     string
	Password string
}

func (u User) TableName() string {
	return "user"
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id=?", id).First(&user).Error
	return user, err
}

func CreateUserTest(id int, name string) (int, string, error) {
	user := User{Id: id, Name: name}
	err := dao.Db.Create(&user).Error
	return user.Id, user.Name, err
}

func UpdateUserTest(id int, name string) (int, string, error) {
	err := dao.Db.Model(&User{}).Where("id=?", id).Update("name", name).Error
	return id, name, err
}

func DeleteUserTest(id int) error {
	err := dao.Db.Delete(&User{}, id).Error
	return err
}
