package rend

import (
	"github.com/romshark/TypeBook/document"
)

// EntityRelationTypeName represents the name of an entity relation type
type EntityRelationTypeName struct {
	SourceType   string
	RelationType string
	TargetType   string
}

// String returns the name of the entity relation type as string
func (n *EntityRelationTypeName) String() string {
	return n.SourceType + "_" + n.RelationType + "_" + n.TargetType
}

// EntityRelationType represents a typed relation between two entities
type EntityRelationType struct {
	TypeName        EntityRelationTypeName
	Description     string
	Metadata        Metadata
	SourceTypeName  string
	TargetTypeName  string
	RelatedTypeName string
	SourceType      AbstractType
	TargetType      AbstractType
	RelatedType     AbstractType
	Direction       document.RelationDirection
}

// TypeCategory implements the AbstractType interface
func (t *EntityRelationType) TypeCategory() TypeCategory {
	return Relation
}

// TotalMetadataFields returns the number of metadata fields
func (t *EntityRelationType) TotalMetadataFields() uint32 {
	return uint32(len(t.Metadata))
}

// Name implements the AbstractType and ComplexType interfaces
func (t *EntityRelationType) Name() string {
	return t.TypeName.String()
}

// MetaInformation implements the ComplexType interface
func (t *EntityRelationType) MetaInformation() Metadata {
	return t.Metadata
}
