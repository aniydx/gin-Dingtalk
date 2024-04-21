package controller

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 1, "order info", "alex", 0)
}

func (o OrderController) GetUserList(c *gin.Context) {
	ReturnError(c, 4004, "order Error")
}
