package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhirajsahh/student-api/internal/config"
	"github.com/dhirajsahh/student-api/internal/http/handlers/student"
)

func main() {

	//load config
	cfg := config.MustLoad()

	//setup router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("server started")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
