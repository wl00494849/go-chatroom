package server

import (
	"github.com/gorilla/websocket"
)

type message struct {
	msg     string
	msgType string
}

type client struct {
	user     string
	conn     *websocket.Conn
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		var msg message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			panic(err)
		}

		switch msg.msgType {
		case "Room":
			c.commands <- command{
				commandType: msg.msgType,
				message:     msg.msg,
				client:      c,
			}
		case "Msg":
			c.commands <- command{
				commandType: msg.msgType,
				message:     msg.msg,
				client:      c,
			}
		case "Quit":
			c.commands <- command{
				commandType: msg.msgType,
				message:     msg.msg,
				client:      c,
			}
		}
	}
}

func (c *client) msg(msg string) {
	c.conn.WriteJSON(msg)
}
