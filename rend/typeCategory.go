package rend

import (
	"fmt"
)

// TypeCategory represents a type category
type TypeCategory uint16

const (
	_ TypeCategory = iota

	// Scalar represents scalar types
	Scalar

	// Enumeration represents enumeration types
	Enumeration

	// Composite represents composite types
	Composite

	// Entity represents entity types
	Entity

	// Relation represents entity relation types
	Relation
)

// String stringifies the value
func (tc TypeCategory) String() string {
	switch tc {
	case Scalar:
		return "scalar"
	case Enumeration:
		return "enumeration"
	case Composite:
		return "composite"
	case Entity:
		return "entity"
	case Relation:
		return "relation"
	}
	panic(fmt.Errorf("couldn't stringify invalid TypeCategory value: %d", tc))
}
