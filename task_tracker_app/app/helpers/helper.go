package helpers

import (
	"encoding/json"
	"net/http"
)

func ParseRequest(r *http.Request, payload any) error {
	return json.NewDecoder(r.Body).Decode(payload)
}

func SendResponse(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(payload)
}

func SendErrorResponse(w http.ResponseWriter, status int, error error) {
	SendResponse(w, status, map[string]string{"error": error.Error()})
}

func GetTokenFromRequest(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func GetUserIDFromContext(r *http.Request) int {
	return r.Context().Value("userID").(int)
}
