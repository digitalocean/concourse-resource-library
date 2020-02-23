package github

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-github/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// GithubClient for handling requests to the Github V3 and V4 APIs.
type GithubClient struct {
	V4         *githubv4.Client
	Repository string
	Owner      string
}

// NewGithubClient ...
func NewGithubClient(s *Source) (*GithubClient, error) {
	owner, repository, err := parseRepository(s.Repository)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	httpClient := http.Client{}

	// Skip SSL verification for self-signed certificates
	// source: https://github.com/google/go-github/pull/598#issuecomment-333039238
	if s.SkipSSLVerification {
		log.Println("disabling SSL verification")
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &httpClient)
	}

	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: s.AccessToken},
	))

	if s.PreviewSchema {
		log.Println("attaching preview schema transport to client")
		client.Transport = &PreviewSchemaTransport{
			oauthTransport: client.Transport,
		}
	}

	var v3 *github.Client
	if s.V3Endpoint != "" {
		endpoint, err := url.Parse(s.V3Endpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to parse v3 endpoint: %s", err)
		}
		v3, err = github.NewEnterpriseClient(endpoint.String(), endpoint.String(), client)
		if err != nil {
			return nil, err
		}
	} else {
		v3 = github.NewClient(client)
	}

	var v4 *githubv4.Client
	if s.V4Endpoint != "" {
		endpoint, err := url.Parse(s.V4Endpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to parse v4 endpoint: %s", err)
		}
		v4 = githubv4.NewEnterpriseClient(endpoint.String(), client)
		if err != nil {
			return nil, err
		}
	} else {
		v4 = githubv4.NewClient(client)
	}

	return &GithubClient{
		V3:         v3,
		V4:         v4,
		Owner:      owner,
		Repository: repository,
	}, nil
}
