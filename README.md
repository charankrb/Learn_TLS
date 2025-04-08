# Learn TLS with Go

This project demonstrates a simple implementation of a TLS client-server communication in Go. It includes a server that listens for secure connections and a client that connects to the server, sends a message, and receives a response. The project is designed to help you understand the TLS handshake process and secure communication.

## Features
- TLS handshake implementation.
- Secure communication between client and server.
- Custom certificate validation.
- Reference to TLS 1.3 handshake steps.

## Prerequisites
- Go installed on your system. [Download Go](https://go.dev/dl/)
- OpenSSL installed for certificate generation. [Download OpenSSL](https://slproweb.com/products/Win32OpenSSL.html)

## Certificate Creation
To create the required certificates for the server:
1. Create an OpenSSL configuration file (`openssl.cnf`) with the following content:
   ```plaintext
   [req]
   default_bits       = 2048
   distinguished_name = req_distinguished_name
   req_extensions     = req_ext
   x509_extensions    = v3_req
   prompt             = no

   [req_distinguished_name]
   C  = US
   ST = State
   L  = City
   O  = Organization
   OU = Organizational Unit
   CN = localhost

   [req_ext]
   subjectAltName = @alt_names

   [v3_req]
   subjectAltName = @alt_names

   [alt_names]
   DNS.1 = localhost

Generate the private key:
```
openssl genrsa -out server.key 2048
```
Generate the self-signed certificate:
```
openssl req -x509 -new -nodes -key server.key -sha256 -days 365 -out server.crt -config openssl.cnf
```
RUN without building 
```
go run main.go server.go client.go
```
Sample output 
```                              
Step 1: Server loaded certificate and private key.
Step 2: Server configured TLS settings.
Step 3: Server is listening on port 8443...
Step 1: Loaded server's certificate into the trusted root CA pool.
Step 2: Configured TLS settings with custom root CA pool.
Step 4: Server accepted a client connection.
Client connected: [::1]:60587
Step 3: Client successfully connected to the server.
Step 4: Sending message to the server: Hello, TLS Server!
Step 4: Client successfully sent the message.
Step 6: Server received: Hello, TLS Server!
Step 8: Server echoed the data back to the client.
Step 5: Client received server response: Hello, TLS Server!
