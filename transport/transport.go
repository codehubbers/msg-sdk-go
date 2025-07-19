package transport

// Transport defines the behavior of a connection medium.
// transport types (TCP, WebSocket, etc.) will follow this contract.
import (
	"fmt"
	"net"
)

// please follow comments for further contributions.
// transport defines the behavior of a connection meedium.
type Transport interface {
	Dial(address string) error
	Send(data []byte) error
	Receive() ([]byte, error)
	Close() error
}

// below implememts trans. interface over TCP.
type TCPTransport struct {
	conn net.Conn
}

// Dial connects to a remote TCP peer at given address.
func (t *TCPTransport) Dial(address string) error {
	// connecting to the guy over TCP
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("TCP dial failed: %v", err)
	}
	t.conn = conn
	return nil
}

// Send writes data to TCP
func (t *TCPTransport) Send(data []byte) error {
	if t.conn == nil {
		return fmt.Errorf("no active TCP connection")
	}
	_, err := t.conn.Write(data)
	return err
}

// Receive reads data from the TCP
func (t *TCPTransport) Receive() ([]byte, error) {
	if t.conn == nil {
		return nil, fmt.Errorf("no active TCP connection")
	}

	buffer := make([]byte, 4096) // receive buffer size
	n, err := t.conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
}

// Close closes the TCP connection.
func (t *TCPTransport) Close() error {
	if t.conn == nil {
		return fmt.Errorf("no active TCP connection to close")
	}
	return t.conn.Close()
}
