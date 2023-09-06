package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var (
	port = flag.Int("port", 8080, "Defaults to 8080")
)

func main() {
	flag.Parse()

	fmt.Printf("Connecting to port %d\n", *port)

	if err := connect(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:8080", nil)
	if err != nil {
		return err
	}
	defer c.Close(websocket.StatusInternalError, "Something went wrong")

	err = wsjson.Write(ctx, c, "hi from client")
	if err != nil {
		fmt.Println("Error writing to socket")
		return err
	}

	c.Close(websocket.StatusNormalClosure, "")
	return nil
}