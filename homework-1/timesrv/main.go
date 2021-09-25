package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

type client chan<- string
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	cfg := net.ListenConfig{
		KeepAlive: time.Minute,
	}

	l, err := cfg.Listen(ctx, "tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	log.Println("I'm started!")

	go inputReader()
	go broadcaster(ctx)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			conn, err := l.Accept()
			if err != nil {
				log.Println(err)
				continue
			}

			clt := make(chan string)
			entering <- clt
			wg.Add(1)
			go handleConn(ctx, conn, clt, wg)
		}
	}()

	<-ctx.Done()

	log.Println("done")
	if err := l.Close(); err != nil {
		log.Println("Failed to close listener!")
	}
	wg.Wait()
	log.Println("exit")
}

// handleConn writes current time and messages from the channel to the client
// Inputs:
//    ctx  - context to close the goroutine
//    conn - connection to handle
//    messages - string channel with messages
//    wg - wait group is used to notify when goroutine is finished
func handleConn(ctx context.Context, conn net.Conn, messages <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	// start time ticker with 1 second period
	tck := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return

		// writes current time to the client
		case t := <-tck.C:
			fmt.Fprintf(conn, "now: %s\n", t)

		// writes a message to the client
		case msg := <-messages:
			fmt.Fprintf(conn, "%s\n", msg)

		default:
		}
	}
}

// inputReader reads stdin and puts user input to the messages channel
func inputReader() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		messages <- msg
	}
}

// broadcaster sends messages from the channel to all clients
// Inputs:
//   ctx - context to close the goroutine
func broadcaster(ctx context.Context) {
	clients := make(map[client]bool)
	for {
		select {
		case <-ctx.Done():
			return

		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)

		default:
		}
	}
}
