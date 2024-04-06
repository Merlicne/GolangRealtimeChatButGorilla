package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	ws_connection *websocket.Conn
	username      string
	room          *Room
}

func newClient(c *websocket.Conn, name string, room *Room) *Client {
	return &Client{
		ws_connection: c,
		username:      name,
		room:          room,
	}
}

func (c *Client) readMessage() {
	for {
		messageType, msg, err := c.ws_connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Println("Read error :", err)
			}
			break
		}

		log.Println(messageType)
		log.Println(c.username, ":", string(msg))
	}
}
