package tts

// Handles OpenAI API calls for text-to-speech

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const elevenLabsAPIURL = "https://api.elevenlabs.io/v1/text-to-speech/%s"

// Calls ElevenLabs API and returns MP3 data
func GenerateTTS(voiceID, text, apiKey string, stability, similarityBoost, style, speed float64, useSpeakerBoost bool) ([]byte, error) {

	url := fmt.Sprintf(elevenLabsAPIURL, voiceID)

	payload, err := json.Marshal(map[string]interface{}{
		"text":     text,
		"model_id": "eleven_multilingual_v2",
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("xi-api-key", apiKey)

	// Send request to Eleven Labs API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response (MP3 file content)
	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return audioData, nil
}
