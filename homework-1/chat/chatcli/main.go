package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Try to use the following commands:")
	fmt.Println("    \\nick - change your nickname")
	fmt.Println("    \\q    - quit the chat")

	defer conn.Close()
	go func() {
		io.Copy(os.Stdout, conn)
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// read Stdin until user send \q command or press ctrl+z
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Fprintln(conn, msg)

			if strings.HasPrefix(msg, "\\q") {
				stop()
				break
			}
		}
	}()

	<-ctx.Done()

	fmt.Println("quit")
}
