package artifactory

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
