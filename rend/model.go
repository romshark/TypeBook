package rend

import (
	"docbuilder/doc"
	"fmt"
	"time"
)

type Model struct {
	Doc             *doc.Document
	Registry        Registry
	Build           string
	RendererVersion string
}

// NewModel initializes a new document model based on a document template.
// Redeclaration of predefined types is checked.
// References to undefined types are checked
func NewModel(
	document *doc.Document,
	options *ModelInitOptions,
) (*Model, []string, *ModelInitStats, error) {
	if document == nil {
		return nil, nil, nil, fmt.Errorf("Missing document template")
	}

	if options == nil {
		options = DefaultModelInitOptions()
	}

	// Determine predefined types
	//TODO: make predefined types configurable
	predefinedScalarTypes := map[string]doc.ScalarType{
		"Bool": doc.ScalarType{
			Description: "Boolean value that's either true or false",
		},
		"Number": doc.ScalarType{
			Description: "A signed floating point number",
		},
		"String": doc.ScalarType{
			Description: "A UTF8 encoded text value",
		},
	}

	// Create a document model instance
	model := &Model{
		Doc:             document,
		Build:           time.Now().UTC().Format(time.RFC3339),
		RendererVersion: rendererVersion,
		Registry: Registry{
			PredefinedScalarTypes: predefinedScalarTypes,
			Types: make(
				map[string]interface{},
				len(predefinedScalarTypes)+
					len(document.ScalarTypes)+
					len(document.EnumerationTypes)+
					len(document.CompositeTypes),
			),
		},
	}

	errorMessages := make([]string, 0, 8)
	addError := func(format string, a ...interface{}) {
		errorMessages = append(errorMessages, fmt.Sprintf(format, a...))
	}

	if err := model.Registry.build(
		addError,
		document,
		predefinedScalarTypes,
	); err != nil {
		return nil, nil, nil, fmt.Errorf("Couldn't build registry: %s", err)
	}

	// Verify null references, ensure that all referenced objects exist
	if options.CheckReferences {
		for typeName, compositeType := range document.CompositeTypes {
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
