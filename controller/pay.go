package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type PayStruct struct{}

func (p PayStruct) GetPayList(c *gin.Context) {
	param := make(map[string]interface{})
	fmt.Println(param)
	err := c.BindJSON(&param)
	if err != nil {
		ReturnError(c, 400, "Error Pay")
		return
	}
	ReturnSuccess(c, 20001, "Success Pay", param, 1)
}
