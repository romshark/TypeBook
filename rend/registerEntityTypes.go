package rend

import (
	"github.com/romshark/TypeBook/document"
)

// registerEntityTypes registers new entity types.
// It will automatically set the type names,
// the names and types of the metadata fields and
// the names and types of the relation metadata fields.
// It won't register in case of type name collisions returning an error
func (d *Document) registerEntityTypes(
	newEntityTypes EntityTypes,
) (errors ModelErrors) {
	forwardDeclared := make(Types, len(newEntityTypes))

	// Verify entity types verifying their type names and metadata
	for entityTypeName, newType := range newEntityTypes {
		newType.TypeName = entityTypeName

		// Verify type name
		errors.Add(d.verifyTypeName(
			entityTypeName,
			"entity type declaration",
		)...)
		if errors.HasErrors() {
			// Don't evaluate further in case of illegal name
			continue
		}

		// Verify type
		errors.Add(d.verifyType(
			nil,
			entityTypeName,
			"entity type declaration", // error location
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

		forwardDeclared[entityTypeName] = newType
	}

	// Verify the integrity of all entity relations
	for _, entityType := range newEntityTypes {
		for relationName, relation := range entityType.Relations {
			// Verify all relations before committing any changes to the model
			errs := d.verifyRelationIntegrity(
				forwardDeclared,
				entityType,
				relationName,
				relation,
			)
			if errs.HasErrors() {
				errors.Add(errs...)
				continue
			}

			// Set source type reference
			if sourceTypeDeclaredRef, sourceIsDeclared :=
				d.Types[relation.SourceTypeName]; sourceIsDeclared {
				relation.SourceType = sourceTypeDeclaredRef
			} else {
				relation.SourceType = forwardDeclared[relation.SourceTypeName]
			}

			// Set target type reference
			if targetTypeDeclaredRef, targetIsDeclared :=
				d.Types[relation.TargetTypeName]; targetIsDeclared {
				relation.TargetType = targetTypeDeclaredRef
			} else {
				relation.TargetType = forwardDeclared[relation.TargetTypeName]
			}

			// Set related type reference
			if relatedTypeDeclaredRef, relatedIsDeclared :=
				d.Types[relation.RelatedTypeName]; relatedIsDeclared {
				relation.RelatedType = relatedTypeDeclaredRef
			} else {
				relation.RelatedType = forwardDeclared[relation.RelatedTypeName]
			}

			entityType.Relations[relationName] = relation
		}
	}

	// Don't commit any changes to the model if the registration
	// of new entity types and their relations failed
	if errors.HasErrors() {
		return errors
	}

	// Successfully register the new types together with their relations
	for typeName, newEntityType := range newEntityTypes {
		d.EntityTypes[typeName] = newEntityType
		d.Types[typeName] = newEntityType
		for _, relationType := range newEntityType.Relations {
			name := relationType.TypeName.String()
			d.Relations[name] = relationType
			d.Types[name] = relationType
		}
	}
	return nil
}

// RegisterEntityTypes returns nil if all given entity types
// were registered in the document model, otherwise returns errors
func (d *Document) RegisterEntityTypes(
	newTypes map[string]document.EntityType,
) ModelErrors {
	// Prepare entity types for registration
	newEntityTypes := make(EntityTypes, len(newTypes))
	for typeName, entityType := range newTypes {
		// Parse entity metadata
		metadata := make(Metadata, len(entityType.Metadata))
		for fieldName, field := range entityType.Metadata {
			metadata[fieldName] = TypedField{
				Description: field.Description,
				Nullable:    field.Nullable,
				TypeName:    field.Type.Name,
				IsList:      field.Type.IsList,
				// Leave Name undefined, it will be set automatically
				// Leave Type undefined, ref will be set automatically
			}
		}

		// Parse entity relations
		relations := make(
			map[string]*EntityRelationType,
			len(entityType.Relations),
		)
		for relationName, relation := range entityType.Relations {
			// Parse relation metadata
			relationMetadata := make(Metadata, len(relation.Metadata))
			for fieldName, field := range relation.Metadata {
				relationMetadata[fieldName] = TypedField{
					Description: field.Description,
					Nullable:    field.Nullable,
					TypeName:    field.Type.Name,
					IsList:      field.Type.IsList,
					// Leave Name undefined, it will be set automatically
					// Leave Type undefined, ref will be set automatically
				}
			}

			var sourceTypeName, targetTypeName string
			var relationTypeName EntityRelationTypeName
			if relation.Direction == document.OutboundRelation {
				// (this type) -> relatedType
				relationTypeName = EntityRelationTypeName{
					SourceType:   typeName,
					RelationType: relation.Type,
					TargetType:   relation.RelatedType,
				}
				sourceTypeName = typeName
				targetTypeName = relation.RelatedType
			} else {
				// (this type) <- relatedType
				relationTypeName = EntityRelationTypeName{
					SourceType:   relation.RelatedType,
					RelationType: relation.Type,
					TargetType:   typeName,
				}
				sourceTypeName = relation.RelatedType
				targetTypeName = typeName
			}
			relations[relationName] = &EntityRelationType{
				Description:     relation.Description,
				Metadata:        relationMetadata,
				Direction:       relation.Direction,
				SourceTypeName:  sourceTypeName,
				TargetTypeName:  targetTypeName,
				RelatedTypeName: relation.RelatedType,
				TypeName:        relationTypeName,
				// Leave SourceType undefined, ref will be set automatically
				// Leave TargetType undefined, ref will be set automatically
				// Leave RelatedType undefined, ref will be set automatically
				// Leave Type undefined, ref will be set automatically
			}
		}

		newEntityTypes[typeName] = &EntityType{
			// Leave TypeName undefined, it will be set automatically
			Description: entityType.Description,
			Metadata:    metadata,
			Relations:   relations,
		}
	}
	// Try to register the new entity types
	return d.registerEntityTypes(newEntityTypes)
}
