package graderapiclient

import (
	"io"
	"net/http"
)

func (g *GraderApiClient) AuthenticatedPost(url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", g.AccessToken)

	return http.DefaultClient.Do(req)
}

func (g *GraderApiClient) AuthenticatedPut(url string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", g.AccessToken)

	return http.DefaultClient.Do(req)
}

func (g *GraderApiClient) AuthenticatedGet(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", g.AccessToken)

	return http.DefaultClient.Do(req)
}
