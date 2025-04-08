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