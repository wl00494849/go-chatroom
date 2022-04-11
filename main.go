package main

import (
	"go-chatroom/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", controller.SocketPage)
	app.GET("/ws", controller.CreateSocket)

}
