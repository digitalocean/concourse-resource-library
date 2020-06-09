package artifactory

import (
	"bufio"
	"os"
	"strings"

	"github.com/jfrog/jfrog-client-go/artifactory/buildinfo"
)

// Property stores an artifact metadata property field / value pair
type Property struct {
	Name  string
	Value string
}

// Properties stores a slice of Property's
type Properties []Property

// PropertySeparator is the standard character used by artifactory
const PropertySeparator = ";"

func (props Properties) String() string {
	var s string

	for _, p := range props {
		s += p.Name + "=" + p.Value + PropertySeparator
	}

	return s
}

// FromFile reads a simple kv file into Properties
func (props *Properties) FromFile(p string) error {
	f, err := os.Open(p)
	defer f.Close()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		kv := strings.SplitN(scanner.Text(), "=", 2)
		if len(kv) == 2 {
			*props = append(*props, Property{Name: kv[0], Value: strings.Trim(kv[1], `'"`)})
		}
	}

	return nil
}

// Env converts Properties to Artifactory buildinfo.Env
func (props Properties) Env() buildinfo.Env {
	e := buildinfo.Env{}

	for _, p := range props {
		e[p.Name] = p.Value
	}

	return e
}
