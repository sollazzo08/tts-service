package api

import (
	"net/http"

	"github.com/sollazzo08/tts-service/internal/tts"
)

func TTSHandler(ttsService *tts.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		// query := r.URL.Query()
		// voiceID := query.Get("voice_id")


		w.Header().Set("Content-Type", "application/json")
	}

}
