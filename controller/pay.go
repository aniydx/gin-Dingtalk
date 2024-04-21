package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Search struct {
	Cid  int    `json:"cid"`
	Name string `json:"name"`
}

type PayStruct struct{}

func (p PayStruct) GetPayList(c *gin.Context) {
	param := make(map[string]interface{})
	fmt.Println(param)
	err := c.BindJSON(&param)
	if err != nil {
		ReturnError(c, 200, "Error Pay: "+string(err.Error()))
		return
	}
	ReturnSuccess(c, 20001, "Success Pay", param, 1)
}

func (p PayStruct) GetPayInfo(c *gin.Context) {
	fmt.Println("--->>>")
	search := &Search{}
	err := c.BindJSON(&search)
	if err != nil {
		ReturnError(c, 200, "Get pay info err: "+err.Error())
	}
	ReturnSuccess(c, 200, "Get pay info success", search, 1)
}
