package server

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)


type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() {
	log.Println("Server running on port:", s.port)
	err := http.ListenAndServe(":"+s.port, s.Routes())
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}