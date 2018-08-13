package rend

import (
	"fmt"
	"time"

	"github.com/romshark/TypeBook/document"
)

// NewModel initializes a new document model based on a document template
func NewModel(
	doc *document.Document,
) (
	model *Document,
	errors ModelErrors,
	stats *ModelInitStats,
	err error,
) {
	if doc == nil {
		return nil, nil, nil, fmt.Errorf("missing document template")
	}

	// Create a document model instance
	model, err = NewDocument(
		time.Now().UTC(),
		doc.Title,
		doc.Author,
		doc.Version,
		doc.Description,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	// Try to register the new scalar types
	for typeName, scalarType := range doc.ScalarTypes {
		errors.Add(model.RegisterScalarType(
			typeName,
			scalarType.Description,
		)...)
	}

	// Try to register the new enumeration types
	for typeName, enumerationType := range doc.EnumerationTypes {
		errors.Add(model.RegisterEnumerationType(
			typeName,
			enumerationType.Description,
			enumerationType.Values,
		)...)
	}

	errors.Add(model.RegisterCompositeTypes(doc.CompositeTypes)...)
	errors.Add(model.RegisterEntityTypes(doc.EntityTypes)...)

	stats = &ModelInitStats{}
	return model, errors, stats, nil
}
