package docker

import (
	"testing"

	_ "github.com/digitalocean/concourse-resource-library/log"
	. "github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
)

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
			option:      Authentication("username", "password"),
			key:         "Authentication",
			value:       "eyJ1c2VybmFtZSI6InVzZXJuYW1lIiwicGFzc3dvcmQiOiJwYXNzd29yZCJ9",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			c, err := NewClient(tc.option)

			switch {
			case tc.key == "Authentication":
				Expect(t, err).To(BeNil())
				Expect(t, c.authentication).To(Equal(tc.value))
			default:
				Expect(t, err).To(Not(BeNil()))
			}
		})
	}
}

func TestPull(t *testing.T) {
	c, err := NewClient()
	Expect(t, err).To(BeNil())

	err = c.Pull("golang:1.14-alpine3.12")
	Expect(t, err).To(BeNil())
}

func TestImage(t *testing.T) {
	c, err := NewClient()
	Expect(t, err).To(BeNil())

	img, err := c.Image("golang:1.14-alpine3.12")
	Expect(t, err).To(BeNil())

	Expect(t, img.ID).To(Equal("sha256:30df784d62066da7df8995f99da527fc5b9971b478699f1de4fd9f54a0f945cb"))
}
