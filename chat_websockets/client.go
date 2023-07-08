package main

import "github.com/gorilla/websocket"

type client struct {
	send  chan []byte
	conn  *websocket.Conn
	group *group
}

func (c *client) read() {
	defer c.conn.Close()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		c.group.msgs <- msg
	}
}

func (c *client) write() {
	defer c.conn.Close()
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}

	}
}
