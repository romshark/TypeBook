package rend

import (
	"github.com/romshark/TypeBook/document"
)

// registerCompositeTypes registers new composite types.
// It will automatically set the type names as well as
// the names and types of the metadata fields.
// It won't register in case of type name collisions returning an error
func (d *Document) registerCompositeTypes(
	newTypes CompositeTypes,
) (errors ModelErrors) {
	for typeName, newType := range newTypes {
		newType.TypeName = typeName

		// Verify type name
		errors.Add(d.verifyTypeName(
			typeName,
			"composite type declaration",
		)...)
		if errors.HasErrors() {
			// Don't evaluate further in case of illegal name
			continue
		}

		errors.Add(d.verifyType(
			nil,
			typeName,
			"composite type declaration", // error location
		)...)
		if errors.HasErrors() {
			// Don't evaluate further in case of invalid type
			continue
		}

		// Verify metadata
		errors.Add(d.verifyMetadataIntegrity(
			nil,
			newType,
		)...)
	}

	// Return errors if any
	if errors.HasErrors() {
		return errors
	}

	// Successfully register the new types
	for typeName, newType := range newTypes {
		d.CompositeTypes[typeName] = newType
		d.Types[typeName] = newType
	}
	return nil
}

// RegisterCompositeTypes returns nil if all given composite types
// were registered in the document model, otherwise returns errors
func (d *Document) RegisterCompositeTypes(
	newTypes map[string]document.CompositeType,
) ModelErrors {
	newCompositeTypes := make(CompositeTypes, len(newTypes))
	for typeName, compositeType := range newTypes {
		// Parse composite metadata
		metadata := make(Metadata, len(compositeType.Metadata))
		for fieldName, field := range compositeType.Metadata {
			metadata[fieldName] = TypedField{
				Description: field.Description,
				Nullable:    field.Nullable,
				TypeName:    field.Type.Name,
				IsList:      field.Type.IsList,
				// Leave Name undefined, it will be set automatically
				// Leave Type undefined, ref will be set automatically
			}
		}

		newCompositeTypes[typeName] = &CompositeType{
			// Leave TypeName undefined, it will be set automatically
			Description: compositeType.Description,
			Metadata:    metadata,
		}
	}
	// Try to register the new composite types
	return d.registerCompositeTypes(newCompositeTypes)
}
