package main

type group struct {
	msgs    chan []byte
	clients map[*client]bool
	add     chan *client
	remove  chan *client
}
