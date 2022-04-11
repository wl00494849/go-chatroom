package controller

import (
	"go-chatroom/server"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var servers = server.NewServer()

func SocketPage(ctx *gin.Context) {
	ctx.HTML(200, "./view/index.html", nil)
}

func CreateSocket(ctx *gin.Context) {
	name := ctx.Query("name")
	ws, err := websocket.Upgrade(ctx.Writer, ctx.Request, nil, 1024, 1024)
	if err != nil {
		ctx.JSON(500, err)
	}
	servers.NewClient(ws, name)
}
