package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	Clients ClientList
	rooms   RoomList
	sync.RWMutex
}

func newManager() *Manager {
	return &Manager{
		Clients: make(ClientList),
		rooms:   make(RoomList),
	}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	room_name := r.URL.Query().Get("room")
	log.Println("- A user", username, "has been connected")

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("cannot upgrade", err)
		return
	}
	room := m.getRoom(room_name)
	client := newClient(conn, username, room)
	m.addClient(client)
	room.userConnect(client)

	go client.readMessage()
}

func (m *Manager) getRoom(room_name string) *Room {
	if room, OK := m.rooms[room_name]; !OK {
		room = newRoom(room_name)
		m.rooms[room_name] = room
		return room
	}
	return m.rooms[room_name]
}

func (m *Manager) addClient(c *Client) {
	m.Lock()
	defer m.Unlock()
	m.Clients[c] = true
}

func (m *Manager) removeClient(c *Client) {
	m.Lock()
	m.Unlock()

	if _, OK := m.Clients[c]; OK {
		room := c.room
		room.userDisconnect(c)
		delete(m.Clients, c)
		c.ws_connection.Close()
	}
}
