package rend

import "fmt"

// verifyMetaFieldType returns errors if the given type (ref)
// can't be used as a metadata field type, otherwise returns nil
func verifyMetaFieldType(
	originTypeName,
	fieldName,
	fieldTypeName string,
	ref AbstractType,
) (errors ModelErrors) {
	switch ref.(type) {
	case *EntityType:
		// Can't use entity types as field types!
		errors.AddErrEntityNesting(
			fieldTypeName,
			originTypeName,
			fieldName,
		)
	case *EntityRelationType:
		// Can't use relation types as field types!
		errors.AddErrRelationTypeAsField(
			fieldTypeName,
			originTypeName,
			fieldName,
		)
	case nil:
		panic(fmt.Errorf(
			"Unexpected nil type reference "+
				"during metadata field '%s' type verification "+
				" of type %s, expected a reference to type '%s'",
			fieldName,
			originTypeName,
			fieldTypeName,
		))
	}
	return errors
}

// verifyMetadataIntegrity helps verifying integrity of metadata fields
// of composite-, entity- or relation types.
// It verifies the existence of referenced types and ensures
// that entity- and relation types are not used as metadata field types.
//
// forwardDeclared represents any forward-declared types
// that are not yet registered in the document model
func (d *Document) verifyMetadataIntegrity(
	forwardDeclared Types,
	origin ComplexType,
) (errors ModelErrors) {
	if origin == nil {
		panic(fmt.Errorf("missing origin type"))
	}

	metadata := origin.MetaInformation()

	for fieldName, field := range metadata {
		// Check whether the type of the field is declared
		// in either the registry or the list of new yet unregistered types
		registryRef, isDeclared := d.Types[field.TypeName]

		var forwardDeclaredRef AbstractType
		isForwardDeclared := false
		if len(forwardDeclared) > 0 {
			forwardDeclaredRef, isForwardDeclared = forwardDeclared[field.TypeName]
		}

		if isDeclared {
			if errs := verifyMetaFieldType(
				origin.Name(),
				fieldName,
				field.TypeName,
				registryRef,
			); errs != nil {
				errors.Add(errs...)
				continue
			}
		} else if isForwardDeclared {
			if errs := verifyMetaFieldType(
				origin.Name(),
				fieldName,
				field.TypeName,
				forwardDeclaredRef,
			); errs != nil {
				errors.Add(errs...)
				continue
			}
		} else {
			// Referenced type is undefined
			errors.AddErrUndefinedTypeInMetaField(
				origin,         // origin type
				fieldName,      // field name
				field.TypeName, // undefined type
			)
			continue
		}

		// Link the type reference
		var typeReference AbstractType
		if isDeclared {
			typeReference = registryRef
		} else {
			typeReference = forwardDeclaredRef
		}

		// Reinitialize field
		metadata[fieldName] = TypedField{
			Name:        fieldName,
			Type:        typeReference,
			Description: field.Description,
			Nullable:    field.Nullable,
			IsList:      field.IsList,
			TypeName:    field.TypeName,
		}
	}

	return errors
}
