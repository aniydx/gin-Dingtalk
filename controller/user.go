package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goDingTalk/models"
	"strconv"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")

	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserTest(id)
	ReturnSuccess(c, 1, name, user, 0)
}

func (u UserController) CreateUser(c *gin.Context) {
	student := models.User{}
	err := c.ShouldBindJSON(&student)
	if err != nil {
		fmt.Println(err.Error())
	}

	id, name, _ := models.CreateUserTest(student.Id, student.Name)
	ReturnSuccess(c, 200, id, name, 10)
}

func (u UserController) GetUserList(c *gin.Context) {
	ReturnError(c, 4004, "user Error")
}

func (u UserController) UpdateUser(c *gin.Context) {
	student := models.User{}
	err := c.ShouldBindJSON(&student)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, name, _ := models.UpdateUserTest(student.Id, student.Name)
	ReturnSuccess(c, 200, id, name, 10)
}

func (u UserController) DeleteUser(c *gin.Context) {
	student := models.User{}
	err := c.ShouldBindJSON(&student)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = models.DeleteUserTest(student.Id)
	if err != nil {
		ReturnError(c, 200, err.Error())
	}
	ReturnSuccess(c, 200, "Delete msg success", student.Id, 10)
}

func (u UserController) GetUserListTest(c *gin.Context) {
	student := models.User{}
	err := c.ShouldBindJSON(&student)
	if err != nil {
		fmt.Println(err.Error())
	}
	s, err := models.GetUserListTest(student.Id)
	if err != nil {
		ReturnError(c, 4004, err.Error())
	}
	ReturnSuccess(c, 200, "List User msg success", s, 10)
}
