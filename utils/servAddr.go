package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type ServAddr struct {
	host string
	port string
}

func NewServerAddress() string {
	err := godotenv.Load("application.env")
	if err != nil {
		log.Println("error oppening env file.")
	}
	port := os.Getenv("server.port")
	addr := os.Getenv("server.address")
	sa := ServAddr{
		host: addr,
		port: port,
	}
	return sa.String()
}

func (s *ServAddr) String() string {
	conAddr := strings.Builder{}
	conAddr.WriteString(s.host)
	conAddr.WriteString(":")
	conAddr.WriteString(s.port)
	return conAddr.String()
}
