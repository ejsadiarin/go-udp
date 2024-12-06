// Go uses the net package for network programming.
// Familiarize yourself with the net package's ListenPacket, Dial, and PacketConn interfaces.

// A UDP server typically listens for packets on a specific address and port using ListenPacket. A client sends data using Dial or DialUDP.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"go-udp/client"
	"go-udp/server"
)

func main() {
	// get args to choose which to start

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What to start (1 - client, 2 - server): ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	choice = strings.TrimSpace(choice)
	switch choice {
	case "1":
		client.StartClient()
	case "2":
		server.StartServer()
	default:
		fmt.Println("Internal server error")
	}
}
