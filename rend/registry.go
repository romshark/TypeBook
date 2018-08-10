package rend

import (
	"fmt"
	"reflect"

	"github.com/romshark/TypeBook/document"
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
	PredefinedScalarTypes map[string]document.ScalarType
	PredefinedTypes       map[string]interface{}

	TotalScalarTypes      uint32
	TotalEnumerationTypes uint32
	TotalCompositeTypes   uint32
	TotalTypes            uint32
}

// build builds the registry verifying that predefined types are not redeclared
func (r *Registry) build(
	addError func(format string, a ...interface{}),
	doc *document.Document,
	predefinedScalarTypes map[string]document.ScalarType,
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
			case *document.ScalarType:
				addError(fmt.Sprintf(
					"Redeclaration of type '%s' "+
						"('%s' is declared scalar type)",
					typeName,
					typeName,
				))
			case *document.EnumType:
				addError(fmt.Sprintf(
					"Redeclaration of type '%s' "+
						"('%s' is declared enumeration type)",
					typeName,
					typeName,
				))
			case *document.CompositeType:
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
	for typeName, meta := range doc.ScalarTypes {
		if err := checkRedeclaration(typeName); err != nil {
			return err
		}
		r.Types[typeName] = &meta
	}
	for typeName, meta := range doc.EnumerationTypes {
		if err := checkRedeclaration(typeName); err != nil {
			return err
		}
		r.Types[typeName] = &meta
	}
	for typeName, meta := range doc.CompositeTypes {
		if err := checkRedeclaration(typeName); err != nil {
			return err
		}
		r.Types[typeName] = &meta
	}

	// Count the types
	r.TotalScalarTypes = uint32(
		len(predefinedScalarTypes) + len(doc.ScalarTypes),
	)
	r.TotalEnumerationTypes = uint32(len(doc.EnumerationTypes))
	r.TotalCompositeTypes = uint32(len(doc.CompositeTypes))
	r.TotalTypes = r.TotalScalarTypes +
		r.TotalEnumerationTypes +
		r.TotalCompositeTypes

	return nil
}
