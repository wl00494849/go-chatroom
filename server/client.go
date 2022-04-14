package server

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type message struct {
	Msg     string `json:"msg"`
	MsgType string `json:"msgType"`
	User    string `json:"user"`
}

type client struct {
	user     string
	conn     *websocket.Conn
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	defer c.conn.Close()
	for {
		msg := &message{}
		_, b, err := c.conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		json.Unmarshal(b, &msg)

		c.commands <- command{
			commandType: msg.MsgType,
			message:     msg.Msg,
			client:      c,
		}
	}
}

func (c *client) msg(msg string) {
	c.conn.WriteJSON(&message{Msg: msg})
}
