package controller

import "github.com/gin-gonic/gin"

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 1, "user info", "GET Param id/name is: "+id+" "+name, 0)
}

func (u UserController) GetUserList(c *gin.Context) {
	ReturnError(c, 4004, "user Error")
}
