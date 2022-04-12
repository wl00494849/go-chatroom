package controller

import (
	"fmt"
	"go-chatroom/server"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var uparader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

func CreateSocket(ctx *gin.Context) {
	name := ctx.Query("name")
	fmt.Println(name)
	ws, err := uparader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(500, err)
	}
	server.GetChatServer().NewClient(ws, name)
}
