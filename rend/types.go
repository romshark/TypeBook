package rend

import (
	"time"
)

// AbstractType can represent any scalar-, enumeration-, composite-, entity-
// relation type
type AbstractType interface {
	TypeCategory() TypeCategory
	Name() string
}

// ComplexType represents either a composite-, entity- or relation type
type ComplexType interface {
	// Name returns the type name
	Name() string

	// MetaInformation returns the typed metadata fields
	MetaInformation() Metadata

	// Returns the total number of metadata fields
	TotalMetadataFields() uint32
}

// TypedField represents a typed field of a composite or entity type
type TypedField struct {
	Name        string
	Description string

	// Nullable indicates whether or not this field is nullable
	Nullable bool

	// IsList indicates whether or not this field is list of the specified type
	IsList bool

	// Type references the type of this field
	Type     AbstractType
	TypeName string
}

// Metadata maps the field names to a metadata field
type Metadata map[string]TypedField

// Relations maps a related relation name to the relation metadata
type Relations map[string]*EntityRelationType

// DocumentMetadata represents document metadata
type DocumentMetadata struct {
	Title           string
	Author          string
	Version         string
	Description     string
	Build           time.Time
	RendererVersion string
}

// ScalarTypes maps type names to distinct scalar types
type ScalarTypes map[string]*ScalarType

// EnumerationTypes maps type names to distinct enumeration types
type EnumerationTypes map[string]*EnumerationType

// CompositeTypes maps type names to distinct composite types
type CompositeTypes map[string]*CompositeType

// EntityTypes maps type names to distinct entity types
type EntityTypes map[string]*EntityType

// EntityRelationTypes maps entity relation type names to lists of relations
type EntityRelationTypes map[string]*EntityRelationType

// Types maps all document types by type name
type Types map[string]AbstractType

// Document represents a document model
type Document struct {
	Metadata         DocumentMetadata
	ScalarTypes      ScalarTypes
	EnumerationTypes EnumerationTypes
	CompositeTypes   CompositeTypes
	EntityTypes      EntityTypes
	Relations        EntityRelationTypes
	Types            Types
}

func NewDocument(
	currentTime time.Time,
	title string,
	author string,
	version string,
	description string,
) (*Document, error) {
	return &Document{
		Metadata: DocumentMetadata{
			Title:           title,
			Author:          author,
			Version:         version,
			Description:     description,
			Build:           time.Unix(currentTime.Unix(), 0).UTC(),
			RendererVersion: rendererVersion,
		},
		ScalarTypes:      make(ScalarTypes),
		EnumerationTypes: make(EnumerationTypes),
		CompositeTypes:   make(CompositeTypes),
		EntityTypes:      make(EntityTypes),
		Relations:        make(EntityRelationTypes),
		Types:            make(Types),
	}, nil
}

// IsTypeDefined returns true if a type with the given name is defined
func (d *Document) IsTypeDefined(typeName string) bool {
	_, isDefined := d.Types[typeName]
	return isDefined
}

// TotalRelations returns the total count of defined relations
func (d *Document) TotalRelations() uint32 {
	return uint32(len(d.Relations))
}

// TotalScalarTypes returns the total count of defined scalar types
func (d *Document) TotalScalarTypes() uint32 {
	return uint32(len(d.ScalarTypes))
}

// TotalEnumerationTypes returns the total count of defined enumeration types
func (d *Document) TotalEnumerationTypes() uint32 {
	return uint32(len(d.EnumerationTypes))
}

// TotalCompositeTypes returns the total count of defined composite types
func (d *Document) TotalCompositeTypes() uint32 {
	return uint32(len(d.CompositeTypes))
}

// TotalEntityTypes returns the total count of defined entity types
func (d *Document) TotalEntityTypes() uint32 {
	return uint32(len(d.EntityTypes))
}

// TotalTypes returns the total count of defined types
func (d *Document) TotalTypes() uint32 {
	return uint32(len(d.Types))
}
