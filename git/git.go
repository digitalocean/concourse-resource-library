package git

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Client provides a common interface to perform Git operations
type Client struct {
	repository      *git.Repository
	accessToken     string
	sslVerification bool
	directory       string
	memory          bool
	depth           int
}

// AccessToken sets the access token used to query the API
func AccessToken(token string) func(*Client) error {
	return func(c *Client) error {
		return c.setAccessToken(token)
	}
}

// NewClient builds a new Git client
func NewClient(repository string, options ...func(*Client) error) (*Client, error) {
	c := &Client{
		depth:           250,
		sslVerification: true,
	}

	return c, nil
}

func (c *Client) setAccessToken(token string) error {
	c.accessToken = token

	return nil
}

// Clone a repository to directory
func (g *Client) Clone(url, reference string, depth int) (*git.Repository, error) {
	r, err := git.PlainClone(g.Directory, false, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: "user",
			Password: g.AccessToken,
		},
		ReferenceName: plumbing.ReferenceName(reference),
		Depth:         depth,
		Progress:      os.Stdout,
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}
