package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/coder/websocket"
)

type client struct {
	socket *websocket.Conn
	recive chan []byte
	room   *room
}

type message struct {
	Text string `json:"message"`
}

func (c *client) read(ctx context.Context) {
	defer c.socket.Close(websocket.StatusNormalClosure, "normal close connection")

	for {
		_, msg, err := c.socket.Read(ctx)
		if err != nil {
			log.Printf("cant read from socket: %v", err)
			return
		}

		c.room.forward <- msg
	}

}

func (c *client) write(ctx context.Context) {
	defer c.socket.Close(websocket.StatusNormalClosure, "normal close connection")

	for msg := range c.recive {
		value := message{}
		json.Unmarshal(msg, &value)
		err := c.socket.Write(ctx, websocket.MessageText, []byte(value.Text))
		if err != nil {
			log.Printf("cant write in socket: %v", err)
			return
		}
	}
}
