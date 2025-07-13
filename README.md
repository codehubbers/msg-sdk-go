# msg-sdk-go

A modular Go SDK for secure, end-to-end encrypted messaging over the command line. Built on the [Noise Protocol Framework](https://noiseprotocol.org/), the SDK enables encrypted communication using Noise handshakes and supports both TCP and WebSocket transports. It provides a flexible foundation for building secure peer-to-peer or client-server messaging tools in Go.

## Project Status

This project is in early development. The initial scaffolding, interface definitions, and architecture layout are currently being implemented.

## Overview

This SDK provides the core components for building encrypted messaging clients:

- Noise-based encrypted sessions using `Noise_XX_25519_AESGCM_SHA256`
- Abstraction over multiple transport layers (TCP and WebSocket)
- Keypair and peer identity management
- Session lifecycle and message exchange logic
- A command-line interface (CLI) for interacting with the SDK

## Key Features

- Noise protocol handshake (mutual authentication and encryption)
- Peer-to-peer and client-server support
- Abstract transport layer with interchangeable backends (TCP, WebSocket)
- CLI commands for initiating sessions, sending, and receiving encrypted messages
- In-memory or file-based session persistence
- Modular, extensible architecture to support future protocols (e.g., QUIC)

## Architecture

```mermaid
sequenceDiagram
    participant CLI
    participant SDK_Client as SDK Core API
    participant Identity
    participant Session
    participant Transport
    CLI->>SDK_Client: Dial("tcp", "peer.server.com:8080")
    SDK_Client->>Identity: GetLocalKeypair()
    Identity-->>SDK_Client: localKeypair
    SDK_Client->>Identity: GetPeerPublicKey("peer.server.com:8080")
    Identity-->>SDK_Client: remotePublicKey
    SDK_Client->>Transport: DialTCP("peer.server.com:8080")
    Transport-->>SDK_Client: rawConnection
    SDK_Client->>Session: NewInitiatorSession(rawConnection, localKeypair, remotePublicKey)
    Session->>Session: Perform Noise Handshake (XX pattern)
    Note over Session: Writes handshake message 1 to transport
    Session->>Transport: Write(handshakeMsg1)
    Transport-->>Session: Read(handshakeMsg2)
    Note over Session: Reads handshake message 2 from transport
    Session->>Transport: Write(handshakeMsg3)
    Note over Session: Writes handshake message 3 to transport<br>Handshake complete. Cipherstates established.
    Session-->>SDK_Client: secureSession
    SDK_Client-->>CLI: Session established successfully



# Getting Started

## Requirements
- Go 1.21 or higher
- Git

## Clone the Repository

```bash
git clone https://github.com/JumaOchi/msg-sdk-go.git
cd msg-sdk-go
```

## Build the Project
```bash
go build
```

## Run the App
```bash
./msg-sdk-go
```

## Project Structure
```
/
├── main.go        // Entry point
├── go.mod         // Go module definition
├── transport/     // Transport abstraction (TCP, WebSocket)
├── identity/      // Keypair generation and storage
├── cli/           // CLI logic and command handling
└── README.md      // Project documentation
```

## Roadmap
- Implement TCP transport
- Establish Noise_XX session setup
- Add WebSocket transport
- Add CLI command for connect
- Add CLI command for send
- Add CLI command for receive
- Add optional file-based session persistence

## Contributing
Contributions are welcome. To contribute:
1. Fork the repository
2. Create a new feature branch
3. Open a pull request with a clear description

Please follow the existing file structure and keep commits clear and focused.

## License
To be added. MIT or Apache-2.0 recommended.

## Maintainers
Developed and maintained by [https://codehubbers.com/](https://codehubbers.com/)

GitHub: [github.com/codehubbers](https://github.com/codehubbers)