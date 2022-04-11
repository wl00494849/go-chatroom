package controller

import (
	"go-chatroom/server"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var socketMap = make(map[string]*websocket.Conn)

func SocketPage(ctx *gin.Context) {
	ctx.HTML(200, "./view/index.html", nil)
}

func CreateSocket(ctx *gin.Context) {
	ws, err := websocket.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header(), 1024, 1024)
	if err != nil {
		ctx.JSON(500, err)
	}

	server.NewClient(ws)
}
