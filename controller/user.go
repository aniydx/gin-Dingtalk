package controller

import "github.com/gin-gonic/gin"

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 1, "user info", "alex", 0)
}

func (u UserController) GetUserList(c *gin.Context) {
	ReturnError(c, 4004, "user Error")
}
