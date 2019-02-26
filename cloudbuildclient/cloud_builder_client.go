package cloudbuildclient

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
	cloudbuild "google.golang.org/api/cloudbuild/v1"
	"google.golang.org/api/googleapi"
)

type StatusCodeError struct {
	Message string
	Code    int
}

func (e StatusCodeError) Error() string {
	return e.Message
}

func LatestStatusWithContext(ctx context.Context, client *http.Client, project string) (string, *StatusCodeError) {
	svc, _ := cloudbuild.New(client)
	res, _err := svc.Projects.Builds.List(project).Do()

	if _err != nil {
		log.Println(_err)
		switch _err.(type) {
		case *googleapi.Error:
			return "", &StatusCodeError{Message: _err.Error(), Code: _err.(*googleapi.Error).Code}
		default:
			return "", &StatusCodeError{Message: _err.Error(), Code: 500}
		}

	}
	if len(res.Builds) < 1 {
		return "", &StatusCodeError{Message: "No Builds"}
	}
	return res.Builds[0].Status, nil

}

func DefaultGoogleClient(ctx context.Context) (*http.Client, error) {
	return google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
}
