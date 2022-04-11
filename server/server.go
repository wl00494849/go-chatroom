package server

import "github.com/gorilla/websocket"

type command struct {
	commandType string
	client      *client
}
type server struct {
	rooms    map[string]*room
	commands chan string
}

func NewServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan string),
	}
}

func NewClient(ws *websocket.Conn) {
	ws.WriteJSON(&message{user: "admin", msg: "New Client has Connection"})
	c := &client{conn: ws}
	c.readInput()
}
