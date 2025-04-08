package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
)

func StartClient() {
	// Step 1: Load the server's certificate
	// We need to verify the server's identity to ensure we are communicating with the intended server.
	// To do this, we load the server's certificate into a certificate pool that will act as our trusted root CA.
	certPool := x509.NewCertPool()
	serverCert, err := ioutil.ReadFile("server.crt") // Path to the server's certificate
	if err != nil {
		log.Fatalf("Failed to read server certificate: %v", err)
	}
	if !certPool.AppendCertsFromPEM(serverCert) {
		log.Fatalf("Failed to append server certificate to the pool")
	}
	fmt.Println("Step 1: Loaded server's certificate into the trusted root CA pool.")

	// Step 2: Configure TLS settings
	// We configure the TLS settings to use the custom root CA pool created in Step 1.
	// This ensures that the client will trust the server's certificate only if it matches the one in the pool.

	config := &tls.Config{
		RootCAs:            certPool, // Use the custom root CA pool
		InsecureSkipVerify: true,     // Skip hostname verification (for testing only)
	}

	//config := &tls.Config{
	//	RootCAs: certPool, // Use the custom root CA pool
	//}
	fmt.Println("Step 2: Configured TLS settings with custom root CA pool.")

	// Step 3: Connect to the server
	// Establish a secure TLS connection to the server using the configured settings.
	// This step ensures that all communication between the client and server is encrypted.
	conn, err := tls.Dial("tcp", "localhost:8443", config)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	fmt.Println("Step 3: Client successfully connected to the server.")

	// Step 4: Send a message to the server
	// After establishing a secure connection, we send a message to the server.
	// This demonstrates how data can be securely transmitted over the TLS connection.
	message := "Hello, TLS Server!"
	fmt.Println("Step 4: Sending message to the server:", message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Println("Step 4: Client successfully sent the message.")

	// Step 5: Read the server's response
	// Finally, we read the server's response to verify that the communication is working as expected.
	// This step ensures that the server can process the client's message and respond appropriately.
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	fmt.Println("Step 5: Client received server response:", string(buf[:n]))
}
