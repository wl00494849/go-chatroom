package server

import (
	"net"
)

type room struct {
	mambers map[net.Addr]*client
	id      string
}

func (r *room) broadcast(sender *client, msg string) {
	for _, member := range r.mambers {
		member.msg(msg)
	}
}
