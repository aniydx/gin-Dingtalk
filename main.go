package main

import (
	"goDingTalk/router"
)

func main() {
	r := router.Router()
	if err := r.Run(":8080"); err != nil {
		return
	}
}
