package main

import (
	"fmt"
	"msg-sdk-go/identity"
	"msg-sdk-go/transport"
	"time"
)

// entry point for the CLI SDK.
func main() {
	fmt.Println("Noise CLI SDK starting up...")

	// Run a test *TCP
	testTCPTransport()
}

// below func establishes a TCP connection and sends/receives a message
func testTCPTransport() {
	tcp := &transport.TCPTransport{}

	err := tcp.Dial("localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer tcp.Close()

	fmt.Println("Connected to -P 8080")

	msg := []byte("Hello from SDK test!\n")
	err = tcp.Send(msg)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	kp, _ := identity.GenerateKeypair()
	fmt.Printf("Public Key: %x\n", kp.PublicKey)

	// Wait a bit if you’re testing against a local echo server *need to be optimised
	time.Sleep(1 * time.Second)

	reply, err := tcp.Receive()
	if err != nil {
		fmt.Println("Error receiving message:", err)
		return
	}

	fmt.Println("Received:", string(reply))
}
