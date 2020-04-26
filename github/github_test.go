package github

import (
	"fmt"
	"testing"

	_ "github.com/digitalocean/concourse-resource-library/log"
	. "github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
)

func TestNewClientRepository(t *testing.T) {
	tests := []struct {
		description string
		repository  string
		expectError bool
	}{
		{
			description: "empty repository",
			repository:  "",
			expectError: true,
		},
		{
			description: "invalid repository",
			repository:  "invalid",
			expectError: true,
		},
		{
			description: "valid repository",
			repository:  "valid/repository",
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			c, err := NewClient(tc.repository)
			if tc.expectError {
				Expect(t, err).To(Not(BeNil()))
				return
			}

			Expect(t, err).To(BeNil())
			Expect(t, fmt.Sprintf("%s/%s", c.Owner, c.Repository)).To(Equal(tc.repository))
		})
	}
}

func TestNewClientOptions(t *testing.T) {
	tests := []struct {
		description string
		option      func(*Client) error
		key         string
		value       string
	}{
		{
			description: "nil pointer option",
			option:      nil,
			key:         "",
			value:       "",
		},
		{
			description: "set API Endpoint",
			option:      Endpoint("https://myghe.com/api/v9"),
			key:         "Endpoint",
			value:       "https://myghe.com/api/v9",
		},
		{
			description: "disable SSL Verification",
			option:      DisableSSLVerification(),
			key:         "DisableSSLVerification",
			value:       "",
		},
		{
			description: "set access token",
			option:      AccessToken("mytoken"),
			key:         "AccessToken",
			value:       "mytoken",
		},
		{
			description: "set Preview Schema",
			option:      PreviewSchema("accept-schemas"),
			key:         "PreviewSchema",
			value:       "accept-schemas",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			c, err := NewClient("some/repo", tc.option)

			switch {
			case tc.key == "Endpoint":
				Expect(t, err).To(BeNil())
				Expect(t, c.endpoint).To(Equal(tc.value))
			case tc.key == "DisableSSLVerification":
				Expect(t, err).To(BeNil())
				Expect(t, c.sslVerification).To(BeFalse())
			case tc.key == "AccessToken":
				Expect(t, err).To(BeNil())
				Expect(t, c.accessToken).To(Equal(tc.value))
			case tc.key == "PreviewSchema":
				Expect(t, err).To(BeNil())
				Expect(t, c.previewSchema).To(Equal(tc.value))
			default:
				Expect(t, err).To(Not(BeNil()))
			}
		})
	}
}
