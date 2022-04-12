package controller

import (
	"go-chatroom/server"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func CreateSocket(ctx *gin.Context) {
	name := ctx.Query("name")
	uparader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	ws, err := uparader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(500, err)
	}
	server.GetChatServer().NewClient(ws, name)
}
