package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

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
	Sender    string `json:"sender,omitempty"` // 为空则不输出
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
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
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			cm.send(jsonMessage, conn)
		case conn := <-cm.unregister:
			if _, ok := cm.clients[conn]; ok {
				close(conn.send)
				delete(cm.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				cm.send(jsonMessage, conn)
			}
		case message := <-cm.broadcast:
			for conn := range cm.clients {
				conn.send <- message
			}
		}
	}
}

// 向所有已连接的 client 发送消息
func (cm *ClientManager) send(message []byte, client *Client) {
	for conn := range cm.clients {
		if conn != client {
			conn.send <- message
		}
	}
}

func (c *Client) read() {
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			return
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

// write 和 client 中有一个关闭 conn 和 发送 unregister 就可以了
func (c *Client) write() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func main() {
	fmt.Println("Start application...")
	go manager.start()

	http.HandleFunc("/ws", wsPage)
	http.ListenAndServe(":12345", nil)
}

func wsPage(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(w, r, nil)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	uid := uuid.NewV4()

	client := &Client{
		id:     uid.String(),
		socket: conn,
		send:   make(chan []byte),
	}
	manager.register <- client

	go client.read()
	go client.write()
}
