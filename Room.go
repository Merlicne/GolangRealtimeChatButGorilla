package main

import "sync"

type RoomList map[string]*Room

type Room struct {
	room_name      string
	messageHistory MessageList
	Clients        ClientList
	sync.RWMutex
}

func newRoom(name string) *Room {
	return &Room{
		room_name: name,
		Clients:   make(ClientList),
	}
}

func (r *Room) userConnect(c *Client) {
	r.Lock()
	defer r.Unlock()
	r.Clients[c] = true
}

func (r *Room) userDisconnect(c *Client) {
	r.Lock()
	defer r.Unlock()

	if _, OK := r.Clients[c]; OK {
		delete(r.Clients, c)
	}
}
