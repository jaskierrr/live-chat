package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/coder/websocket"
)

type room struct {
	clients map[*client]struct{}
	join    chan *client
	leave   chan *client
	forward chan []byte
}

func newRoom() *room {
	return &room{
		clients: make(map[*client]struct{}),
		join:    make(chan *client),
		leave:   make(chan *client),
		forward: make(chan []byte),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = struct{}{}
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.recive)
		case msg := <-r.forward:
			for client := range r.clients {
				client.recive <- msg
			}
		}
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatalf("cant connect to WS: %v", err)
		return
	}

	defer conn.Close(websocket.StatusNormalClosure, "normal close")

	client := &client{
		socket: conn,
		recive: make(chan []byte),
		room:   r,
	}

	fmt.Printf("client join %v\n",client.room)

	r.join <- client
	go client.read(ctx)
	client.write(ctx)
}
