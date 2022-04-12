package main

import (
	"flag"
	"go-chatroom/controller"
	"go-chatroom/server"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	var port string
	flag.StringVar(&port, "p", ":8080", "port")

	if p := os.Getenv("PORT"); len(p) != 0 {
		port = ":" + p
	}

	server.ServerInit()

	app := gin.Default()
	app.LoadHTMLGlob("view/*")
	app.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	app.GET("/ws", controller.CreateSocket)

	app.Run(port)
}
