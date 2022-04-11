package server

import "github.com/gorilla/websocket"

type command struct {
	commandType string
	client      *client
	message     string
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

func (s *server) NewClient(ws *websocket.Conn, name string) {
	ws.WriteJSON("New Client has Connection")
	c := &client{
		conn:     ws,
		user:     "AnyOne",
		commands: s.commands,
	}
	c.readInput()
}
