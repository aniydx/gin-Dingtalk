package controller

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 1, "user info", "alex", 0)
}

func GetUserList(c *gin.Context) {
	ReturnError(c, 4004, "Erro")
}
