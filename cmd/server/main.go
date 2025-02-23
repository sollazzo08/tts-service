package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sollazzo08/tts-service/config"
	"github.com/sollazzo08/tts-service/internal/api"
)

func main() {

	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	http.HandleFunc("/generate-tts", api.GenerateSpeech)

	// Start the server
	port := "8080"
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
