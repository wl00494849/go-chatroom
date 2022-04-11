package server

import (
	"net"

	"github.com/gorilla/websocket"
)

type command struct {
	commandType string
	client      *client
	message     string
}
type server struct {
	rooms    map[string]*room
	commands chan command
}

func NewServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for msg := range s.commands {
		switch msg.commandType {
		case "Room":
			s.joinRoom(msg.client, msg.message)
		case "Msg":
			s.msg(msg.client, msg.message)
		case "Quit":
			s.quit(msg.client, msg.message)
		}
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

func (s *server) joinRoom(c *client, id string) {
	r, ok := s.rooms[id]
	if !ok {
		r = &room{
			id:      id,
			mambers: make(map[net.Addr]*client),
		}
		s.rooms[id] = r
	}
	r.mambers[c.conn.RemoteAddr()] = c
	c.room = r
}
func (s *server) msg(c *client, msg string) {

}

func (s *server) quit(c *client, msg string) {
	if c.room != nil {
		delete(c.room.mambers, c.conn.RemoteAddr())
		c.room.broadcast(c, c.user+":Left room")
	}
}
