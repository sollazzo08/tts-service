package main

import (
	"log"

	"github.com/sollazzo08/tts-service/config"
	"github.com/sollazzo08/tts-service/internal/api"
)

func main()	{

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}


}
