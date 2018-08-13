package rend

// ScalarType represents a distinct scalar type
type ScalarType struct {
	TypeName    string
	Description string
}

// TypeCategory implements the AbstractType interface
func (t *ScalarType) TypeCategory() TypeCategory {
	return Scalar
}

// Name implements the AbstractType interface
func (t *ScalarType) Name() string {
	return t.TypeName
}
