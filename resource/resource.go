package resource

import (
	"encoding/json"
	"os"
)

// Version communicated with Concourse.
type Version interface {
}

// Source is the configuration for the resource
type Source interface {
	Validate() error
}

// Parameters is the configuration for a resource step
type Parameters interface {
}

// CheckRequest is the data struct received from Concoruse by the resource check operation
type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

// GetRequest is the data struct received from Concoruse by the resource get operation
type GetRequest struct {
	Source  Source     `json:"source"`
	Params  Parameters `json:"params"`
	Version Version    `json:"version"`
}

// PutRequest is the data struct received from Concoruse by the resource put operation
type PutRequest struct {
	Source Source     `json:"source"`
	Params Parameters `json:"params"`
}

// Metadata output from get/put steps.
type Metadata []*MetadataField

// Add a MetadataField to the Metadata.
func (m *Metadata) Add(name, value string) {
	*m = append(*m, &MetadataField{Name: name, Value: value})
}

// MetadataField ...
type MetadataField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CheckResponse is the data struct returned to Concourse by the resource check operation
type CheckResponse []Version

// Len returns the number of versions in the response
func (r CheckResponse) Len() int {
	return len(r)
}

// Write will write the json response to stdout for Concourse to parse
func (r CheckResponse) Write() error {
	return json.NewEncoder(os.Stdout).Encode(r)
}

// GetPutResponse is the data struct returned to Concourse by the resource get & put operations
type GetPutResponse struct {
	Version  Version  `json:"version"`
	Metadata Metadata `json:"metadata,omitempty"`
}
