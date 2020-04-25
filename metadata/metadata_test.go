package metadata_test

import (
	"fmt"
	"testing"

	"github.com/digitalocean/concourse-resource-library/metadata"
	"github.com/poy/onpar"
	"github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
)

func TestAdd(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) expect.Expectation {
		return expect.New(t)
	})

	tests := []struct {
		description string
		name        string
		value       interface{}
	}{
		{
			description: "store string",
			name:        "ref",
			value:       "3SDA25F",
		},
		{
			description: "store integer",
			name:        "count",
			value:       3,
		},
		{
			description: "store bool",
			name:        "enabled",
			value:       true,
		},
	}

	m := metadata.Metadata{}

	for _, tc := range tests {
		o.Spec(tc.description, func(expect expect.Expectation) {
			m.Add(tc.name, tc.value)
			out := m.Get(tc.name)

			expect(out).To(Equal(fmt.Sprint(tc.value)))
		})
	}
}
