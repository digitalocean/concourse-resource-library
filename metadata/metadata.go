package metadata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Metadata output from get/put steps
type Metadata []*Field

// Field data struct
type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	JSON  bool   `json:"-"`
}

// FileName returns filename for metadata file
func (f *Field) FileName() string {
	if f.JSON {
		return f.Name + ".json"
	}

	return f.Name
}

// Content returns filename for metadata file
func (f *Field) Content() []byte {
	return []byte(f.Value)
}

// Add a Field to the Metadata
func (m *Metadata) Add(name string, value interface{}) {
	*m = append(*m, &Field{Name: name, Value: fmt.Sprintf("%s", value)})
}

// AddJSON marshals data, then adds it to the Metadata
func (m *Metadata) AddJSON(name string, value interface{}) {
	data, _ := json.Marshal(value)
	*m = append(*m, &Field{Name: name, Value: string(data), JSON: true})
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

// ToFiles writes metadata to files on disk as `resource/{Name}` with content {Value}
func (m Metadata) ToFiles(path string) error {
	for _, d := range m {
		if err := ioutil.WriteFile(filepath.Join(path, d.FileName()), d.Content(), 0644); err != nil {
			return fmt.Errorf("failed to write metadata file %s: %s", d.FileName(), err)
		}
	}

	return nil
}
