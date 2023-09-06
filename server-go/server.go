package main

import (
	"flag"
	"fmt"
)

var (
	port = flag.Int("port", 8080, "Defaults to 8080")
)

func main() {
	flag.Parse()

	fmt.Printf("Listening on port %d\n", *port)
}

func listen() {

}