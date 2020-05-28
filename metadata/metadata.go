package metadata

import (
	"encoding/json"
	"fmt"
)

// Metadata output from get/put steps
type Metadata []*Field

// Field data struct
type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Add a Field to the Metadata
func (m *Metadata) Add(name string, value interface{}) {
	*m = append(*m, &Field{Name: name, Value: fmt.Sprintf("%s", value)})
}

// Get a Field of the Metadata
func (m *Metadata) Get(name string) string {
	for _, v := range *m {
		if v.Name == name {
			return v.Value
		}
	}

	return ""
}

// JSON returns an encoded byte slice of an Artifact
func (m *Metadata) JSON() ([]byte, error) {
	return json.Marshal(m)
}
