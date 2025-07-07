package server

import "net/http"

func HandleQueueEnter(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 404, "Not implemented", nil)
}
