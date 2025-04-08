package main

import (
	"time"
)

/*
Client                                Server

	|                                     |
	| ----------- ClientHello ----------> |  (Step 1: Client sends supported cipher suites, TLS version)
	|                                     |
	| <---------- ServerHello ----------- |  (Step 2: Server responds with selected cipher suite, TLS version)
	|                                     |
	| <------- Certificate (Server) ----- |  (Step 3: Server sends its certificate)
	|                                     |
	| -------- Key Exchange (Client) ---> |  (Step 4: Client sends key exchange data)
	|                                     |
	| <------ Finished (Server) --------- |  (Step 5: Server confirms handshake completion)
	|                                     |
	| -------- Finished (Client) -------> |  (Step 6: Client confirms handshake completion)
	|                                     |
	| <---- Encrypted Application Data -->|  (Step 7: Secure communication begins)
*/
func main() {
	// Step 1: Start the server in a separate goroutine
	go func() {
		StartServer()
	}()

	// Step 2: Wait for the server to start
	time.Sleep(2 * time.Second)

	// Step 3: Start the client
	StartClient()
}
