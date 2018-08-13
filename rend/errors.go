package rend

import (
	"fmt"
)

// ErrorCode represents an error code
type ErrorCode string

const (
	ErrIllegalTypeName   ErrorCode = "ErrIllegalTypeName"
	ErrUndefinedType     ErrorCode = "ErrUndefinedType"
	ErrTypeNameCollision ErrorCode = "ErrTypeNameCollision"
	ErrEntityNesting     ErrorCode = "ErrEntityNesting"
	ErrInappropriateType ErrorCode = "ErrInappropriateType"
)

// ModelErr represents a document model error
type ModelErr struct {
	Code     ErrorCode
	Message  string
	Location string
}

// Error implements the standard Go error interface
func (r *ModelErr) Error() string {
	return fmt.Sprintf("%s: %s in %s", r.Code, r.Message, r.Location)
}

// ModelErrors represents a list of model errors
type ModelErrors []ModelErr

// Add adds an error to the error list
// initializing the error list if necessary
func (errs *ModelErrors) Add(errors ...ModelErr) {
	if len(errors) < 1 {
		return
	}
	if *errs == nil {
		*errs = errors
	} else {
		*errs = append(*errs, errors...)
	}
}

// HasErrors returns true if the list contains any errors,
// otherwise returns false
func (errs *ModelErrors) HasErrors() bool {
	return len(*errs) > 0
}

// AddErrIllegalTypeName adds a new illegal type name error
// indicating that a type name violates the type name rules
func (errs *ModelErrors) AddErrIllegalTypeName(
	typeName,
	errLocation string,
) {
	errs.Add(ModelErr{
		Code:     ErrIllegalTypeName,
		Message:  fmt.Sprintf("illegal type name: '%s'", typeName),
		Location: errLocation,
	})
}

// AddErrEntityNesting adds a entity nesting error
func (errs *ModelErrors) AddErrEntityNesting(
	nestedEntityTypeName,
	containerEntityTypeName,
	fieldName string,
) {
	errs.Add(ModelErr{
		Code: ErrEntityNesting,
		Message: fmt.Sprintf(
			"illegal nesting of entity types ('%s' in '%s')",
			nestedEntityTypeName,
			containerEntityTypeName,
		),
		Location: fmt.Sprintf(
			"field '%s' of type '%s'",
			fieldName,
			containerEntityTypeName,
		),
	})
}

// AddErrRelationTypeAsField adds an error
// indicating that a relation type was used to define a field
func (errs *ModelErrors) AddErrRelationTypeAsField(
	relationTypeName,
	containerTypeName,
	fieldName string,
) {
	errs.Add(ModelErr{
		Code: ErrEntityNesting,
		Message: fmt.Sprintf(
			"illegal use of relation type '%s' for field definition",
			relationTypeName,
		),
		Location: fmt.Sprintf(
			"field '%s' of type '%s'",
			fieldName,
			containerTypeName,
		),
	})
}

// AddErrTypeNameCollision adds a new type name collision error
// indicating a type name redeclaration attempt
func (errs *ModelErrors) AddErrTypeNameCollision(
	redeclaredTypeName string,
	typeCategory string,
	errLocation string,
) {
	errs.Add(ModelErr{
		Code: ErrTypeNameCollision,
		Message: fmt.Sprintf(
			"redeclaration of %s type '%s'",
			typeCategory,
			redeclaredTypeName,
		),
		Location: errLocation,
	})
}

// AddErrUndefinedType adds a new undefined type error
// indicating that a referenced type is undefined
func (errs *ModelErrors) AddErrUndefinedType(
	undefinedTypeName string,
	errLocation string,
) {
	errs.Add(ModelErr{
		Code: ErrUndefinedType,
		Message: fmt.Sprintf(
			"undefined type '%s'",
			undefinedTypeName,
		),
		Location: errLocation,
	})
}

// AddErrUndefinedTypeInMetaField adds a new undefined type error
// indicating that an undefined type was used in a metadata field declaration
func (errs *ModelErrors) AddErrUndefinedTypeInMetaField(
	originType ComplexType,
	fieldName,
	undefinedTypeName string,
) {
	errs.AddErrUndefinedType(undefinedTypeName, fmt.Sprintf(
		"field '%s' of type '%s'",
		fieldName,
		originType.Name(),
	))
}

// AddErrUndefinedTypeAsRelation adds a new undefined type error
// indicating that a type referenced by a relation is
func (errs *ModelErrors) AddErrUndefinedTypeAsRelation(
	originType *EntityType,
	relationName,
	undefinedTypeName string,
) {
	errs.AddErrUndefinedType(undefinedTypeName, fmt.Sprintf(
		"relation '%s' of type '%s'",
		relationName,
		originType.Name(),
	))
}

// AddErrInappropriateType adds a new inappropriate type error
// indicating that an inappropriate type was used in a certain situation
func (errs *ModelErrors) AddErrInappropriateType(
	undefinedTypeName string,
	errLocation string,
) {
	errs.AddErrUndefinedType(undefinedTypeName, errLocation)
}
