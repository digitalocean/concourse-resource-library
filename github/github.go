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
	client          *githubv4.Client
	endpoint        string
	accessToken     string
	sslVerification bool
	previewSchema   string
	Repository      string
	Owner           string
}

// Endpoint is used for accessing GitHub Enterprise
func Endpoint(endpoint string) func(*Client) error {
	return func(c *Client) error {
		return c.setEndpoint(endpoint)
	}
}

// DisableSSLVerification does exactly that
func DisableSSLVerification() func(*Client) error {
	return func(c *Client) error {
		return c.disableSSLVerification()
	}
}

// AccessToken sets the access token used to query the API
func AccessToken(token string) func(*Client) error {
	return func(c *Client) error {
		return c.setAccessToken(token)
	}
}

// PreviewSchema sets the Accept header to access preview schemas of the GitHub API, multiple schemas can be accessed via comma separated string
func PreviewSchema(schema string) func(*Client) error {
	return func(c *Client) error {
		return c.setPreviewSchema(schema)
	}
}

// NewClient builds the client used to access the API
func NewClient(repository string, options ...func(*Client) error) (*Client, error) {
	owner, repository, err := parseRepository(repository)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	httpClient := http.Client{}

	c := &Client{
		sslVerification: true,
		Owner:           owner,
		Repository:      repository,
	}

	log.Println("owner:", c.Owner)
	log.Println("repository:", c.Repository)

	for _, option := range options {
		if option == nil {
			return nil, fmt.Errorf("option is nil pointer")
		}

		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	// Skip SSL verification for self-signed certificates
	// source: https://github.com/google/go-github/pull/598#issuecomment-333039238
	if c.sslVerification {
		log.Println("disabling SSL verification")
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &httpClient)
	}

	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.accessToken},
	))

	if c.previewSchema != "" {
		log.Println("attaching preview schema transport to client")
		tc.Transport = &PreviewSchemaTransport{
			oauthTransport: tc.Transport,
		}
	}

	c.client, err = getClient(c.endpoint, tc)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) setEndpoint(endpoint string) error {
	c.endpoint = endpoint

	return nil
}

func (c *Client) setAccessToken(token string) error {
	c.accessToken = token

	return nil
}

func (c *Client) disableSSLVerification() error {
	c.sslVerification = false

	return nil
}

func (c *Client) setPreviewSchema(schema string) error {
	c.previewSchema = schema

	return nil
}

func parseRepository(s string) (string, string, error) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		log.Println("invalid repository config:", s)
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
	schema         string
}

// RoundTrip appends the Accept header and then executes the parent RoundTrip Transport
func (t *PreviewSchemaTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Println("setting Accept header to enable preview schemans", t.schema)
	r.Header.Add("Accept", t.schema)

	return t.oauthTransport.RoundTrip(r)
}

// Query sends a query to the GitHub API
func (c *Client) Query(q interface{}, vars map[string]interface{}) error {
	err := c.client.Query(context.TODO(), q, vars)
	if err != nil {
		return err
	}

	return nil
}
