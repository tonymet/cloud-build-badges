package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
	cloudbuild "google.golang.org/api/cloudbuild/v1"
)

func latestStatusWithContext(ctx context.Context, client *http.Client) ([]byte, error) {
	svc, err := cloudbuild.New(client)
	res, err := svc.Projects.Builds.List("fifth-coral-473").Do()
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	if len(res.Builds) < 1 {
		return []byte{}, errors.New("No Builds")
	}
	return json.Marshal(StatusToShieldJson(res.Builds[0].Status))

}

func defaultGoogleClient(ctx context.Context) (*http.Client, error) {
	return google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
}

/*
func main() {
	// Use oauth2.NoContext if there isn't a good context to pass in.
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		log.Panic(err)
		return
	}
	var status string
	if status, err = latestStatusWithContext(&ctx, client); err != nil {
		log.Panic(err)
		return
	}
	log.Printf("Status = %s", status)
}

*/
