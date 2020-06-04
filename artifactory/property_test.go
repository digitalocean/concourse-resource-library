package artifactory

import (
	"testing"

	_ "github.com/digitalocean/concourse-resource-library/log"
	. "github.com/poy/onpar/expect"
	. "github.com/poy/onpar/matchers"
)

func TestProperties(t *testing.T) {
	tests := []struct {
		description string
		props       Properties
		expected    string
	}{
		{
			description: "empty",
			props:       Properties{},
			expected:    "",
		},
		{
			description: "single",
			props: Properties{
				Property{Name: "myname", Value: "myvalue"},
			},
			expected: "myname=myvalue;",
		},
		{
			description: "multi",
			props: Properties{
				Property{Name: "myname", Value: "myvalue"},
				Property{Name: "myname1", Value: "myvalue1"},
			},
			expected: "myname=myvalue;myname1=myvalue1;",
		},
		{
			description: "build info",
			props: Properties{
				Property{Name: "build.name", Value: "my-pipeline"},
				Property{Name: "build.number", Value: "1"},
			},
			expected: "build.name=my-pipeline;build.number=1;",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			out := tc.props.String()

			Expect(t, out).To(Equal(tc.expected))
		})
	}
}
