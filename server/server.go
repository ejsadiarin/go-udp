package server

import (
	"fmt"
	"log"
	"net"
	"time"
)

func StartServer() {
	/***** server *****/
	// --> fixed port
	// 1. establish fixed port
	// 2. establish listener
	// 3. receive request from UDP socket/port
	// 4. send response to UDP socket/port
	srvaddr, err := net.ResolveUDPAddr("udp", ":8000")
	if err != nil {
		fmt.Printf("Error: %v", err)
		log.Fatal(err)
	}
	listener, err := net.ListenUDP("udp", srvaddr)
	if err != nil {
		fmt.Printf("Error: %v", err)
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Printf("Listening on port: %d\n", 8000)
	buf := make([]byte, 1024) // create a buffer with max size 1024 bytes for holding data
	for {
		// receive bytes
		data_size, addr, err := listener.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("Error: %v", err)
			log.Fatal(err)
		}
		fmt.Printf("Received: %s from %s\n", string(buf[:data_size]), addr)
		// send response to UDP socket/port (still to itself (the server))
		_, err = listener.WriteToUDP([]byte(fmt.Sprintf("Received %v\n", time.Now().String())), addr)
		if err != nil {
			fmt.Printf("Error: %v", err)
			log.Fatal(err)
		}
		// send data to client who sent the recent message
		listener.WriteToUDPAddrPort([]byte("Sent from server"), addr.AddrPort())
	}
}
