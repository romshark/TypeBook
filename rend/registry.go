package rend

import (
	"docbuilder/doc"
	"fmt"
	"reflect"
)

type TypeCategory int

const (
	_ TypeCategory = iota
	Scalar
	Enumeration
	Composite
)

type Registry struct {
	Types                 map[string]interface{}
	PredefinedScalarTypes map[string]doc.ScalarType
	PredefinedTypes       map[string]interface{}

	TotalScalarTypes      uint32
	TotalEnumerationTypes uint32
	TotalCompositeTypes   uint32
	TotalTypes            uint32
}

// build builds the registry verifying that predefined types are not redeclared
func (r *Registry) build(
	addError func(format string, a ...interface{}),
	document *doc.Document,
	predefinedScalarTypes map[string]doc.ScalarType,
) error {
	// Register predefined types
	r.PredefinedTypes = make(
		map[string]interface{},
		len(predefinedScalarTypes),
	)
	for typeName, meta := range predefinedScalarTypes {
		r.Types[typeName] = &meta
		r.PredefinedTypes[typeName] = &meta
	}

	checkRedeclaration := func(typeName string) error {
		// Check whether redaclaring a predefined type
		if _, isDeclared := r.PredefinedTypes[typeName]; isDeclared {
			addError(fmt.Sprintf(
				"Redeclaration of type '%s' "+
					"('%s' is a predefined scalar type)",
				typeName,
				typeName,
			))
			return nil
		}

		// Check whether redaclaring an already defined user type
		if meta, isDeclared := r.Types[typeName]; isDeclared {
			switch meta.(type) {
			case *doc.ScalarType:
				addError(fmt.Sprintf(
					"Redeclaration of type '%s' "+
						"('%s' is declared scalar type)",
					typeName,
					typeName,
				))
			case *doc.EnumType:
				addError(fmt.Sprintf(
					"Redeclaration of type '%s' "+
						"('%s' is declared enumeration type)",
					typeName,
					typeName,
				))
			case *doc.CompositeType:
				addError(fmt.Sprintf(
					"Redeclaration of type '%s' "+
						"('%s' is declared composite type)",
					typeName,
					typeName,
				))
			default:
				return fmt.Errorf(
					"Unexpected type '%s' during type redeclaration check",
					reflect.TypeOf(meta),
				)
			}
		}

		return nil
	}

	// Check for redeclaration of predefined types
	for typeName, meta := range document.ScalarTypes {
		if err := checkRedeclaration(typeName); err != nil {
			return err
		}
		r.Types[typeName] = &meta
	}
	for typeName, meta := range document.EnumerationTypes {
		if err := checkRedeclaration(typeName); err != nil {
			return err
		}
		r.Types[typeName] = &meta
	}
	for typeName, meta := range document.CompositeTypes {
		if err := checkRedeclaration(typeName); err != nil {
			return err
		}
		r.Types[typeName] = &meta
	}

	// Count the types
	r.TotalScalarTypes = uint32(
		len(predefinedScalarTypes) + len(document.ScalarTypes),
	)
	r.TotalEnumerationTypes = uint32(len(document.EnumerationTypes))
	r.TotalCompositeTypes = uint32(len(document.CompositeTypes))
	r.TotalTypes = r.TotalScalarTypes +
		r.TotalEnumerationTypes +
		r.TotalCompositeTypes

	return nil
}
