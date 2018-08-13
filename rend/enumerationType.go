package rend

// EnumerationValues maps the enumeration item to its value
type EnumerationValues map[string]string

// EnumerationType represents a distinct enumeration type
type EnumerationType struct {
	TypeName    string
	Description string

	// Values maps the enumerations to their corresponding values
	Values EnumerationValues
}

// TypeCategory implements the AbstractType interface
func (t *EnumerationType) TypeCategory() TypeCategory {
	return Enumeration
}

// Name implements the AbstractType interface
func (t *EnumerationType) Name() string {
	return t.TypeName
}
