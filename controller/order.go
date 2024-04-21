package controller

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetOrderInfo(c *gin.Context) {
	ReturnSuccess(c, 1, "order info", "alex", 0)
}

func (o OrderController) GetOrderList(c *gin.Context) {
	cid := c.PostForm("cid")
	name := c.PostForm("name")
	//ReturnError(c, 4004, "order Error")
	ReturnSuccess(c, 2000, "GetOrderList", cid+" "+name, 0)
}
