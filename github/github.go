package github

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Client for handling requests to the Github GraphQL API
type Client struct {
	client     *githubv4.Client
	Repository string
	Owner      string
}

// NewClient ...
func NewClient(s *Source) (*Client, error) {
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

	ghClient, err := getClient(s.V4Endpoint, client)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:     ghClient,
		Owner:      owner,
		Repository: repository,
	}, nil
}

func parseRepository(s string) (string, string, error) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return "", "", errors.New("malformed repository")
	}
	return parts[0], parts[1], nil
}

func getClient(uri string, client *http.Client) (*githubv4.Client, error) {
	if uri == "" {
		return githubv4.NewClient(client), nil
	}

	endpoint, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to parse v4 endpoint: %s", err)
	}

	ghClient := githubv4.NewEnterpriseClient(endpoint.String(), client)

	return ghClient, nil
}

// PreviewSchemaTransport is used to access GraphQL schema's hidden behind an Accept header by GitHub
type PreviewSchemaTransport struct {
	oauthTransport http.RoundTripper
}

// RoundTrip appends the Accept header and then executes the parent RoundTrip Transport
func (t *PreviewSchemaTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Println("setting accept header for timelineItems & files connections preview schemas")
	r.Header.Add("Accept", "application/vnd.github.starfire-preview+json, application/vnd.github.ocelot-preview+json")

	return t.oauthTransport.RoundTrip(r)
}
