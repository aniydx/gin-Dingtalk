package controller

import (
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

func (u UserController) GetUserList(c *gin.Context) {
	ReturnError(c, 4004, "user Error")
}
