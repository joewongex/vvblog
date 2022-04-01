package main

import (
	"vvblog/config"
	"vvblog/middleware"
	"vvblog/route/admin"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors(), middleware.ErrHandler())
	admin.Setup(r)
	r.Run(config.App.Port)
}
