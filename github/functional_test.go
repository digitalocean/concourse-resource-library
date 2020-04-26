package github_test

import (
	"testing"

	"github.com/digitalocean/concourse-resource-library/github"
	. "github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		description string
		repository  string
	}{
		{
			description: "store string",
			repository:  "owner/repo",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			c, err := github.NewClient(tc.repository)
			Expect(t, err).To(BeNil())

			Expect(t, c).To(Not(BeNil()))
		})
	}

}
