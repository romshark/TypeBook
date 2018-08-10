package rend

import (
	"fmt"
	"time"

	"github.com/romshark/TypeBook/document"
)

type Model struct {
	Doc             *document.Document
	Registry        Registry
	Build           string
	RendererVersion string
}

// NewModel initializes a new document model based on a document template.
// Redeclaration of predefined types is checked.
// References to undefined types are checked
func NewModel(
	doc *document.Document,
	options *ModelInitOptions,
) (*Model, []string, *ModelInitStats, error) {
	if doc == nil {
		return nil, nil, nil, fmt.Errorf("Missing document template")
	}

	if options == nil {
		options = DefaultModelInitOptions()
	}

	// Determine predefined types
	//TODO: make predefined types configurable
	predefinedScalarTypes := map[string]document.ScalarType{
		"Bool": document.ScalarType{
			Description: "Boolean value that's either true or false",
		},
		"Number": document.ScalarType{
			Description: "A signed floating point number",
		},
		"String": document.ScalarType{
			Description: "A UTF8 encoded text value",
		},
	}

	// Create a document model instance
	model := &Model{
		Doc:             doc,
		Build:           time.Now().UTC().Format(time.RFC3339),
		RendererVersion: rendererVersion,
		Registry: Registry{
			PredefinedScalarTypes: predefinedScalarTypes,
			Types: make(
				map[string]interface{},
				len(predefinedScalarTypes)+
					len(doc.ScalarTypes)+
					len(doc.EnumerationTypes)+
					len(doc.CompositeTypes),
			),
		},
	}

	errorMessages := make([]string, 0, 8)
	addError := func(format string, a ...interface{}) {
		errorMessages = append(errorMessages, fmt.Sprintf(format, a...))
	}

	if err := model.Registry.build(
		addError,
		doc,
		predefinedScalarTypes,
	); err != nil {
		return nil, nil, nil, fmt.Errorf("Couldn't build registry: %s", err)
	}

	// Verify null references, ensure that all referenced objects exist
	if options.CheckReferences {
		for typeName, compositeType := range doc.CompositeTypes {
			for fieldName, field := range compositeType.Metadata {
				_, exists := model.Registry.Types[field.Type.Name]
				if !exists {
					addError(
						"Undefined type '%s' in %s > %s",
						field.Type,
						typeName,
						fieldName,
					)
				}
			}
		}
	}

	return model, errorMessages, &ModelInitStats{}, nil
}
