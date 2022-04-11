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
	commands chan<- string
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

		}
	}
}

func (c *client) msg(msg string) {
	c.conn.WriteJSON(msg)
}
