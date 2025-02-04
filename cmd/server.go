package main

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/VxVxN/powserverclient/pkg/pow"
)

type Server struct {
	rand   *rand.Rand
	quotes []string
}

func NewServer() *Server {
	var quotes = []string{
		"The pessimist sees difficulty in every opportunity. The optimist sees opportunity in every difficulty.",
		"Don’t let yesterday take up too much of today.",
		"You learn more from failure than from success. Don’t let it stop you. Failure builds character.",
		"If you are working on something that you really care about, you don’t have to be pushed. The vision pulls you.",
		"Experience is a hard teacher because she gives the test first, the lesson afterward.",
		"To know how much there is to know is the beginning of learning to live.",
	}
	return &Server{
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		quotes: quotes,
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()
	log.Println("Server is listening on port 8081...")

	server := NewServer()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go server.handleConnection(conn)
	}
}

func (server *Server) generateChallenge() string {
	return strconv.Itoa(server.rand.Intn(999999))
}

func (server *Server) getRandomQuote() string {
	return server.quotes[server.rand.Intn(len(server.quotes))]
}

func (server *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	challenge := server.generateChallenge()
	powData := pow.PerformPoW(challenge)

	_, err := conn.Write([]byte(challenge + "\n"))
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		response := scanner.Text()
		if response == powData {
			_, err = conn.Write([]byte("Correct nonce! Here’s your quote: " + server.getRandomQuote() + "\n"))
			break
		} else {
			conn.Write([]byte("Incorrect nonce. Try again.\n"))
		}
	}
}
