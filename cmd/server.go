package cmd

import (
	"asciiArtWeb/internal/delivery"
	"fmt"
	"log"
	"net/http"
)

func Server() {
	server := delivery.New()
	fmt.Print("Starting server at port 8081...\nhttp://localhost:8081/\n")

	if err := http.ListenAndServe(":8081", server.Route()); err != nil {
		log.Println(err)
		return
	}
}
