package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/VxVxN/powserverclient/pkg/pow"
)

func main() {
	conn, err := net.Dial("tcp", "server:8081")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	challenge, _ := reader.ReadString('\n')
	challenge = strings.TrimSuffix(challenge, "\n")
	log.Printf("Received challenge: %s", challenge)

	nonce := pow.PerformPoW(challenge)
	log.Printf("Finded nonce: %s", nonce)
	fmt.Fprintf(conn, "%s\n", nonce) // Send nonce to server

	response, _ := reader.ReadString('\n')
	log.Print(response)
}
