package server

import (
	"net"
)

type room struct {
	mambers map[net.Addr]*client
	id      string
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, member := range r.mambers {
		if addr != sender.conn.RemoteAddr() {
			member.msg(msg)
		}
	}
}
