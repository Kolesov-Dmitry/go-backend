package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
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

// extractNickName extracts nickname from a line like "\nick nickname"
// Inputs:
//   line - string which holds new nickName
// Output:
//   - returns extracted nickName
func extractNickName(line string) string {
	parts := strings.Split(line, " ")
	if len(parts) == 2 {
		return parts[1]
	}

	return "anonymous"
}

// handleConn handles new client connection
// Inputs:
//    conn - the client connection
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	nickName := "anonymous"
	ch <- "You are " + nickName
	messages <- nickName + " has arrived"
	entering <- ch

	log.Println(nickName + " has arrived")

	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg := input.Text()

		// change nickname if received "\nick" command
		if strings.HasPrefix(msg, "\\nick") {
			newNickName := extractNickName(msg)
			if newNickName != nickName {
				messages <- nickName + ": changes nickname to " + newNickName
				nickName = newNickName
			}
		// quit if received "\q" command
		} else if strings.HasPrefix(msg, "\\q") {
			break
		// otherwise send the message
		} else {
			messages <- nickName + ": " + msg
		}
	}

	leaving <- ch
	messages <- nickName + " has left"
	conn.Close()
}

// clientWriter writes messages from the channel to the client
// Inputs:
//   conn - client connection
//   ch - channel to read messages from
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

// broadcaster puts messages from connected clients to the corresponding channel
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
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
