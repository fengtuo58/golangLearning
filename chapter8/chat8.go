package main

import (
	"fmt"
	"bufio"
	"net"
	"log"
)
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string//an outing message channel

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("broad caster panic", err)
		}
	}()
	for {
		select {
		case msg := <-messages:
			//Broadcast incomming meaasge to all
			//client's outgoing message channels
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	cn := make (chan string)
	go clientWriter(conn, cn)
	who := conn.RemoteAddr().String()
	cn <- "You are" + who
	messages <- who + "has arrived"
	entering <-cn

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <-cn
	messages <- who + ": " + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
