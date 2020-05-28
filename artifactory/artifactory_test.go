package artifactory

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
			option:      Endpoint("https://myart.com/api/v9"),
			key:         "Endpoint",
			value:       "https://myart.com/api/v9",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			c, err := NewClient(tc.option)

			switch {
			case tc.key == "Endpoint":
				Expect(t, err).To(BeNil())
				Expect(t, c.endpoint).To(Equal(tc.value))
			default:
				Expect(t, err).To(Not(BeNil()))
			}
		})
	}
}
