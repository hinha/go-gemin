package users

import (
	"net/http"
)

func RegistersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
