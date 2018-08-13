package rend

// CompositeType represents a distinct composite type
type CompositeType struct {
	TypeName    string
	Description string
	Metadata    Metadata
}

// TypeCategory implements the AbstractType interface
func (t *CompositeType) TypeCategory() TypeCategory {
	return Composite
}

// Name implements the AbstractType and ComplexType interfaces
func (t *CompositeType) Name() string {
	return t.TypeName
}

// MetaInformation implements the ComplexType interface
func (t *CompositeType) MetaInformation() Metadata {
	return t.Metadata
}

// TotalMetadataFields returns the number of metadata fields
func (t *CompositeType) TotalMetadataFields() uint32 {
	return uint32(len(t.Metadata))
}
