package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sollazzo08/tts-service/internal/tts"
)

// GenerateSpeech handles incoming TTS requests using net/http
func GenerateSpeech(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request
	var request struct {
		Text            string  `json:"text"`
		VoiceID         string  `json:"voice_id"`
		Stability       float64 `json:"stability,omitempty"`
		SimilarityBoost float64 `json:"similarity_boost,omitempty"`
		Style           float64 `json:"style,omitempty"`
		Speed           float64 `json:"speed,omitempty"`
		UseSpeakerBoost bool    `json:"use_speaker_boost,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Set default values if not provided
	if request.Stability == 0 {
		request.Stability = 0.5
	}
	if request.SimilarityBoost == 0 {
		request.SimilarityBoost = 0.8
	}
	if request.Style == 0 {
		request.Style = 0.7 // More expressive speech
	}
	if request.Speed == 0 {
		request.Speed = 1
	}

	// Get API key
	apiKey := os.Getenv("ELEVENLABS_API_KEY")
	if apiKey == "" {
		http.Error(w, "Missing Eleven Labs API Key", http.StatusInternalServerError)
		return
	}

	// Call TTS function
	audioData, err := tts.GenerateTTS(request.VoiceID, request.Text, apiKey, request.Stability, request.SimilarityBoost, request.Style, request.Speed, request.UseSpeakerBoost)
	if err != nil {
		http.Error(w, "TTS generation failed", http.StatusInternalServerError)
		return
	}

	// Send MP3 response
	w.Header().Set("Content-Type", "audio/mpeg")
	w.WriteHeader(http.StatusOK)
	w.Write(audioData)
}
