package rend

import (
	"fmt"
)

// verifyRelationType returns errors if the given type (relatedType)
// can't be related to, otherwise returns nil
func verifyRelationType(
	originType *EntityType,
	relationName,
	relationTypeName string,
	relatedType AbstractType,
) (errors ModelErrors) {
	switch relatedType.(type) {
	case *EntityType:
		// Can't use entity types as field types!
		errors.AddErrEntityNesting(
			relationTypeName,
			originType.Name(),
			relationName,
		)
	case *EntityRelationType:
		// Can't use relation types as field types!
		errors.AddErrRelationTypeAsField(
			relationTypeName,
			originType.Name(),
			relationName,
		)
	case nil:
		panic(fmt.Errorf(
			"Unexpected nil type reference "+
				"during metadata field '%s' type verification "+
				" of type %s, expected a reference to type '%s'",
			relationName,
			originType.Name(),
			relationTypeName,
		))
	}
	return errors
}

// verifyRelatedType returns errors if the given type (relatedType)
// can't be related to, otherwise returns nil
func verifyRelatedType(
	originType *EntityType,
	relationName,
	relationTypeName string,
	relatedType AbstractType,
) (errors ModelErrors) {
	switch relatedType.(type) {
	case *EntityType:
		// Can't use entity types as field types!
		errors.AddErrEntityNesting(
			relationTypeName,
			originType.Name(),
			relationName,
		)
	case *EntityRelationType:
		// Can't use relation types as field types!
		errors.AddErrRelationTypeAsField(
			relationTypeName,
			originType.Name(),
			relationName,
		)
	case nil:
		panic(fmt.Errorf(
			"Unexpected nil type reference "+
				"during metadata field '%s' type verification "+
				" of type %s, expected a reference to type '%s'",
			relationName,
			originType.Name(),
			relationTypeName,
		))
	}
	return errors
}

// verifyRelationIntegrity verifies the integrity of an entity relation type.
//
// forwardDeclared represents any forward-declared types
// that are not yet registered in the document model
func (d *Document) verifyRelationIntegrity(
	forwardDeclared Types,
	originType *EntityType,
	relationName string,
	relation *EntityRelationType,
) (errors ModelErrors) {
	if relation == nil {
		panic(fmt.Errorf("missing relation type object"))
	}

	// Verify type name
	errors.Add(d.verifyTypeName(
		relation.TypeName.RelationType,
		fmt.Sprintf(
			"relation '%s' of entity type '%s'",
			relationName,
			originType.Name(),
		),
	)...)
	if errors.HasErrors() {
		// Don't evaluate further in case of illegal name
		return errors
	}

	relationTypeName := relation.TypeName.String()

	// Check type
	errors.Add(d.verifyType(
		forwardDeclared,
		relationTypeName,
		fmt.Sprintf(
			"relation '%s' of entity type '%s'",
			relationName,
			originType.Name(),
		),
	)...)
	if errors.HasErrors() {
		// Don't continue in case of illegal name
		return errors
	}

	// Check metadata fields
	errors.Add(d.verifyMetadataIntegrity(forwardDeclared, relation)...)

	// Verify source type
	sourceTypeRegistry, isDeclared := d.Types[relation.SourceTypeName]
	sourceTypeForwardDeclared, isForwardDeclared := forwardDeclared[relation.SourceTypeName]
	if !isDeclared && !isForwardDeclared {
		errors.AddErrUndefinedType(
			relation.SourceTypeName,
			fmt.Sprintf(
				"source type of relation '%s' of entity type '%s'",
				relationName,
				originType.Name(),
			),
		)
		return errors
	} else if isDeclared {
		// Check source type
		switch sourceTypeRegistry.(type) {
		case *EntityType:
		default:
			errors.AddErrInappropriateType(
				relation.SourceTypeName,
				fmt.Sprintf(
					"source type of relation '%s' of entity type '%s'",
					relationName,
					originType.Name(),
				),
			)
			return
		}
	} else if isForwardDeclared {
		// Check source type
		switch sourceTypeForwardDeclared.(type) {
		case *EntityType:
		default:
			errors.AddErrInappropriateType(
				relation.SourceTypeName,
				fmt.Sprintf(
					"source type of relation '%s' of entity type '%s'",
					relationName,
					originType.Name(),
				),
			)
			return
		}
	}

	// Verify target type
	targetTypeRegistry, isDeclared := d.Types[relation.TargetTypeName]
	targetTypeForwardDeclared, isForwardDeclared := forwardDeclared[relation.TargetTypeName]
	if !isDeclared && !isForwardDeclared {
		errors.AddErrUndefinedType(
			relation.TargetTypeName,
			fmt.Sprintf(
				"target type of relation '%s' of entity type '%s'",
				relationName,
				originType.Name(),
			),
		)
		return errors
	} else if isDeclared {
		// Check target type
		switch targetTypeRegistry.(type) {
		case *EntityType:
		default:
			errors.AddErrInappropriateType(
				relation.TargetTypeName,
				fmt.Sprintf(
					"target type of relation '%s' of entity type '%s'",
					relationName,
					originType.Name(),
				),
			)
			return
		}
	} else if isForwardDeclared {
		// Check target type
		switch targetTypeForwardDeclared.(type) {
		case *EntityType:
		default:
			errors.AddErrInappropriateType(
				relation.TargetTypeName,
				fmt.Sprintf(
					"target type of relation '%s' of entity type '%s'",
					relationName,
					originType.Name(),
				),
			)
			return
		}
	}

	return errors
}
