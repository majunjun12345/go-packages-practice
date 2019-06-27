package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender, omitempty"`
	Recipient string `json:"recipient, omitempty"`
	Content   string `json:"content, omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (cm *ClientManager) start() {
	for {
		select {
		case conn := <-cm.register:
			cm.clients[conn] = true

		}
	}
}

func main() {

}
