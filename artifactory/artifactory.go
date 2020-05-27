package artifactory

import (
	"errors"
	"fmt"
	"log"

	"github.com/jfrog/jfrog-client-go/artifactory"
	rtAuth "github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/auth"
	"github.com/jfrog/jfrog-client-go/config"
)

// Client for handling requests to the Artifactory AQL API
type Client struct {
	client      *artifactory.ArtifactoryServicesManager
	endpoint    string
	apiKey      string
	user        string
	password    string
	accessToken string
	certPath    string
	certKeyPath string
	sshKeyPath  string
}

// Endpoint is used for accessing Artifactory
func Endpoint(endpoint string) func(*Client) error {
	return func(c *Client) error {
		return c.setEndpoint(endpoint)
	}
}

// Authentication sets credentials to be used by the client
func Authentication(user, password, apiKey, accessToken string) func(*Client) error {
	return func(c *Client) error {
		if user != "" {
			err := c.setUser(user)
			if err != nil {
				return err
			}
		}

		switch {
		case user != "" && password != "":
			return c.setPassword(password)
		case apiKey != "":
			return c.setAPIKey(apiKey)
		case accessToken != "":
			return c.setAPIKey(apiKey)
		}

		return errors.New("invalid authentication configuration")
	}
}

// NewClient builds the client used to access the API
func NewClient(options ...func(*Client) error) (*Client, error) {
	c := &Client{}

	for _, option := range options {
		if option == nil {
			return nil, fmt.Errorf("option is nil pointer")
		}

		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	dets := rtDetails(c)

	sc, err := config.NewConfigBuilder().SetServiceDetails(dets).Build()
	if err != nil {
		return nil, err
	}

	c.client, err = artifactory.New(&dets, sc)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) setEndpoint(v string) error {
	c.endpoint = v

	return nil
}

func (c *Client) setUser(v string) error {
	c.user = v

	return nil
}

func (c *Client) setAccessToken(v string) error {
	c.accessToken = v

	return nil
}

func (c *Client) setPassword(v string) error {
	c.password = v

	return nil
}

func (c *Client) setAPIKey(v string) error {
	c.apiKey = v

	return nil
}

func rtDetails(c *Client) auth.ServiceDetails {
	rtDetails := rtAuth.NewArtifactoryDetails()
	rtDetails.SetUrl(c.endpoint)

	if c.accessToken != "" {
		rtDetails.SetAccessToken(c.accessToken)
	}

	if c.sshKeyPath != "" {
		rtDetails.SetSshKeyPath(c.sshKeyPath)
	}

	if c.apiKey != "" {
		rtDetails.SetApiKey(c.apiKey)
	}

	if c.user != "" {
		rtDetails.SetUser(c.user)
	}

	if c.password != "" {
		rtDetails.SetPassword(c.password)
	}

	if c.certPath != "" {
		rtDetails.SetClientCertPath(c.certPath)
	}

	if c.certKeyPath != "" {
		rtDetails.SetClientCertKeyPath(c.certKeyPath)
	}

	return rtDetails
}

// AQL returns the results of an AQL request
func (c *Client) AQL(aql string) ([]byte, error) {
	data, err := c.client.Aql(aql)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}
