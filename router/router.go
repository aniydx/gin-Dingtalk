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

	// 路由组v1
	v1Group := r.Group("v1")
	{
		v1Group.GET("/userinfo", controller.UserController{}.GetUserInfo)
		v1Group.GET("/userlist", controller.UserController{}.GetUserList)
		v1Group.GET("/orderlist", controller.OrderController{}.GetUserList)
		v1Group.GET("/orderinfo", controller.OrderController{}.GetUserInfo)
	}
	return r
}
