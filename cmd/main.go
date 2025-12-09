package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	serv := server.NewServer(logger)

	if err := serv.Start(); err != nil {
		logger.Fatal("Server error:", err)
	}
}
