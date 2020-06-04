package git

import (
	"fmt"
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
func NewClient(options ...func(*Client) error) (*Client, error) {
	c := &Client{
		depth:           250,
		sslVerification: true,
	}

	for _, option := range options {
		if option == nil {
			return nil, fmt.Errorf("option is nil pointer")
		}

		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) setAccessToken(token string) error {
	c.accessToken = token

	return nil
}

// Clone a repository to directory
func (c *Client) Clone(url, reference, directory string, depth int) (*git.Repository, error) {
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: url,
		Auth: &http.BasicAuth{
			Username: "user",
			Password: c.accessToken,
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

// Open a repository from disk
func (c *Client) Open(directory string) (*git.Repository, error) {
	r, err := git.PlainOpen(directory)
	if err != nil {
		return nil, err
	}

	return r, nil
}
