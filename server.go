package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/tonymet/cloud-build-badges/cloudbuildclient"
	"github.com/tonymet/cloud-build-badges/shields"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	currentProject := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ctx := appengine.WithContext(r.Context(), r)
	defaultClient, err := cloudbuildclient.DefaultGoogleClient(ctx)
	if err != nil {
		http.Error(w, "Unable to get client", 404)
		return
	}
	r.ParseForm()
	if projectParam, ok := r.Form["project"]; ok && len(projectParam) > 0 {
		currentProject = projectParam[0]
	}
	badge := shields.New()
	badge.SetLabel(currentProject)
	w.Header().Set("Content-Type", "application/json")
	latestStatus, statusErr := cloudbuildclient.LatestStatusWithContext(ctx, defaultClient, currentProject)
	if statusErr != nil {
		switch statusErr.Code {
		case 403:
			latestStatus = "Permission Denied"
			w.WriteHeader(403)
		default:
			latestStatus = "Unknown Error"
			w.WriteHeader(404)
		}
	}
	jsonResponse, err := json.Marshal(badge.FromStatus(latestStatus))
	if err != nil {
		http.Error(w, "Unable to encode json", 500)
		return
	}
	w.Write(jsonResponse)
}
