package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.WithContext(r.Context(), r)
	defaultClient, err := defaultGoogleClient(ctx)
	if err != nil {
		http.Error(w, "Unable to get client", 404)
		return
	}
	latestStatus, err := latestStatusWithContext(ctx, defaultClient)
	if err != nil {
		http.Error(w, "Unable to get status", 404)
		return
	}
	jsonResponse, err := json.Marshal(StatusToShieldJson(latestStatus))
	if err != nil {
		http.Error(w, "Unable to encode json", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
