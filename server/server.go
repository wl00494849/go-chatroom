package server

import (
	"fmt"
	"net"

	"github.com/gorilla/websocket"
)

var chatServer *server

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

func ServerInit() {
	chatServer = NewServer()
	go chatServer.run()
}

func GetChatServer() *server {
	return chatServer
}

func (s *server) run() {
	for msg := range s.commands {
		switch msg.commandType {
		case "Room":
			s.joinRoom(msg.client, msg.message)
		case "Msg":
			s.sendMsg(msg.client, msg.message)
		case "Quit":
			s.quit(msg.client, msg.message)
		}
	}
}

func (s *server) NewClient(ws *websocket.Conn, name string) {
	ws.WriteJSON(&message{Msg: "New Client has Connection"})
	c := &client{
		conn:     ws,
		user:     name,
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
	c.room.broadcast(c, "歡迎"+c.user+"加入房間")
}
func (s *server) sendMsg(c *client, msg string) {
	fmt.Println(msg)
	if c.room == nil {
		c.conn.WriteJSON(&message{Msg: "you not join any room"})
	}
	c.room.broadcast(c, fmt.Sprintf("%s:%s", c.user, msg))
}

func (s *server) quit(c *client, msg string) {
	if c.room != nil {
		delete(c.room.mambers, c.conn.RemoteAddr())
		c.room.broadcast(c, c.user+":Left room")
	}
}
