package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
  ELEVENLABS_API_KEY string
}

func LoadConfig() (*Config, error) {

  err := godotenv.Load()
  // if error is not empty
  if err != nil {
    log.Fatal("error loading .env file")
  }


  elevenLabsKey := os.Getenv("ELEVENLABS_API_KEY")

  if elevenLabsKey == "" {
    return nil, fmt.Errorf("error loading env variable: ELEVENLABS_API_KEY")
  }

  return &Config{
    ELEVENLABS_API_KEY: elevenLabsKey,
  }, nil
}
