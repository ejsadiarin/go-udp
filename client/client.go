package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func StartClient() {
	/***** client *****/
	// --> A client sends data using Dial or DialUDP.
	// --> dynamic, ephemeral port that the OS applies automatically
	// 1. create client socket
	// 2. read input from keyboard (as message)
	// 3. configure and send to destination server (server name, port)

	addr, err := net.ResolveUDPAddr("udp", "localhost:8000")
	if err != nil {
		fmt.Printf("Error: %v", err)
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Printf("\nConnection error: %v", err)
		log.Fatal(err)
	}
	defer conn.Close()
	// send to msg to the udp server
	buf := make([]byte, 1024)
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Enter input (from client): ")

		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("\nError reading input: %v", err)
			log.Fatal(err)
		}
		msg = strings.TrimSpace(msg)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("\nError: %v", err)
			log.Fatal(err)
		}
		// read response from server
		data_size, srvaddr, err := conn.ReadFromUDPAddrPort(buf)
		if err != nil {
			fmt.Printf("\nError: %v", err)
			log.Fatal(err)
		}
		fmt.Printf("From server %s response: %s\n", srvaddr, buf[:data_size])
	}
}
