package artifactory

import (
	"encoding/json"

	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
)

// Artifact wraps Artifactory metadata for an artifact
type Artifact struct {
	File utils.FileInfo
	Item utils.ResultItem
}

// JSON returns an encoded byte slice of an Artifact
func (a *Artifact) JSON() ([]byte, error) {
	return json.Marshal(a)
}
