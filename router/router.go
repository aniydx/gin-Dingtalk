package router

import (
	"github.com/gin-gonic/gin"
	"goDingTalk/controller"
)

func Router() *gin.Engine {

	r := gin.Default()
	//// 默认路由
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "GET / Success", "data": "index.html"})
	//})

	// 路由组 user
	user := r.Group("user")
	{
		user.GET("/userinfo/:id/:name", controller.UserController{}.GetUserInfo)
		user.GET("/userlist", controller.UserController{}.GetUserList)
	}

	// 路由组 order
	order := r.Group("order")
	{
		order.GET("/orderlist", controller.OrderController{}.GetUserList)
		order.GET("/orderinfo", controller.OrderController{}.GetUserInfo)
	}
	return r
}
