package transport

// Transport defines the behavior of a connection medium.
// transport types (TCP, WebSocket, etc.) will follow this contract.
type Transport interface {
	Dial(address string) error        // Connect to remote peer
	Send(data []byte) error           // Send encrypted data
	Receive() ([]byte, error)         // Read data 
	Close() error                     // Close the connection 
}
