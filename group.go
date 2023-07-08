package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type group struct {
	msgs    chan []byte
	clients map[*client]bool
	add     chan *client
	remove  chan *client
}

func newGroup() *group {
	return &group{
		msgs:    make(chan []byte),
		clients: make(map[*client]bool),
		add:     make(chan *client),
		remove:  make(chan *client),
	}
}

func (g *group) run() {
	for {
		select {
		case client := <-g.add:
			g.clients[client] = true
		case client := <-g.remove:
			delete(g.clients, client)
			close(client.send)
		case msg := <-g.msgs:
			for client := range g.clients {
				client.send <- msg
			}
		}
	}
}

const (
	readSize  = 1024
	writeSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: readSize, WriteBufferSize: writeSize}

func (g *group) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP err: ", err)
		return
	}
	client := &client{
		send:  make(chan []byte, writeSize),
		conn:  conn,
		group: g,
	}
	g.add <- client
	go client.write()
	client.read()
}
