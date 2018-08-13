package rend

// EntityType represents a distinct entity type
type EntityType struct {
	TypeName    string
	Description string
	Metadata    Metadata
	Relations   Relations
}

// TypeCategory implements the AbstractType interface
func (t *EntityType) TypeCategory() TypeCategory {
	return Entity
}

// Name implements the AbstractType and ComplexType interfaces
func (t *EntityType) Name() string {
	return t.TypeName
}

// MetaInformation implements the ComplexType interface
func (t *EntityType) MetaInformation() Metadata {
	return t.Metadata
}

// TotalMetadataFields returns the number of metadata fields
func (t *EntityType) TotalMetadataFields() uint32 {
	return uint32(len(t.Metadata))
}
