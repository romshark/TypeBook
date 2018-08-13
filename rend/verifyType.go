package rend

import (
	"fmt"
	"reflect"
)

// verifyType returns an error if the given type name is either already
// reserved by another registered or forward declared type,
// or if the type name violates the type name rules,
// otherwise returns nil
func (d *Document) verifyType(
	forwardDeclared Types,
	typeName string,
	declarationLocation string,
) (errors ModelErrors) {
	determineName := func(ref AbstractType) string {
		switch ref.(type) {
		case *ScalarType:
			return "scalar"
		case *EnumerationType:
			return "enumeration"
		case *CompositeType:
			return "composite"
		case *EntityType:
			return "entity"
		case *EntityRelationType:
			return "relation"
		}
		panic(fmt.Errorf(
			"unexpected type '%s' during type verification",
			reflect.TypeOf(ref),
		))
	}

	// Verify existence
	forwardDeclaredRef, isForwardDeclared := forwardDeclared[typeName]
	registryRef, isDeclared := d.Types[typeName]

	// Check whether redeclaring an already defined user type
	if isDeclared {
		errors.AddErrTypeNameCollision(
			typeName,
			determineName(registryRef),
			declarationLocation,
		)
	} else if isForwardDeclared {
		errors.AddErrTypeNameCollision(
			typeName,
			determineName(forwardDeclaredRef),
			declarationLocation,
		)
	}

	return errors
}
