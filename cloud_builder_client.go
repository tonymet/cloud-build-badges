package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
	cloudbuild "google.golang.org/api/cloudbuild/v1"
)

func latestStatusWithContext(ctx context.Context, client *http.Client) (string, error) {
	svc, err := cloudbuild.New(client)
	res, err := svc.Projects.Builds.List(os.Getenv("GOOGLE_CLOUD_PROJECT")).Do()
	if err != nil {
		log.Println(err)
		return "", err
	}
	if len(res.Builds) < 1 {
		return "", errors.New("No Builds")
	}
	return res.Builds[0].Status, nil

}

func defaultGoogleClient(ctx context.Context) (*http.Client, error) {
	return google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
}
