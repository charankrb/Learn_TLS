package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
)

// handleConnection handles the communication with a client.
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	// Step 5: Server reads data from the client
	buf := make([]byte, 1024)
	for {
		// TLS Record Layer: Reads encrypted data from the client
		// TLS Header Example:
		// +------------+------------+----------------+
		// | ContentType | Version    | Length         |
		// +------------+------------+----------------+
		// | 0x17 (App) | 0x0303 (TLS 1.2) | 0x0040 (64 bytes) |
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected:", conn.RemoteAddr())
			} else {
				log.Println("Error reading from client:", err)
			}
			return
		}
		fmt.Printf("Step 6: Server received: %s\n", string(buf[:n]))

		// Step 7: Server echoes the data back to the client
		// TLS Record Layer: Sends encrypted data back to the client
		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Println("Error writing to client:", err)
			return
		}
		fmt.Println("Step 8: Server echoed the data back to the client.")
	}
}

// StartServer initializes the TLS server and listens for client connections.
func StartServer() {
	// Step 1: Load server certificate and private key
	// The certificate contains the server's public key and identity.
	// The private key is used to establish secure communication.
	// To generate the certificate and key, you can use the following OpenSSL commands:
	// 1. Generate a private key:
	//    openssl genrsa -out server.key 2048
	// 2. Create a certificate signing request (CSR):
	//    openssl req -new -key server.key -out server.csr
	// 3. Generate a self-signed certificate:
	//    openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
	// The resulting `server.crt` and `server.key` files are used here.
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Failed to load server certificate and key: %v", err)
	}
	fmt.Println("Step 1: Server loaded certificate and private key.")

	// Step 2: Configure TLS settings
	// TLS Configuration Example:
	// - Supported Cipher Suites
	// - Protocol Versions (e.g., TLS 1.2, TLS 1.3)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	fmt.Println("Step 2: Server configured TLS settings.")

	// Step 3: Start listening for TLS connections
	listener, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()
	fmt.Println("Step 3: Server is listening on port 8443...")

	// Step 4: Accept client connections and handle them concurrently
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		fmt.Println("Step 4: Server accepted a client connection.")
		go handleConnection(conn)
	}
}
