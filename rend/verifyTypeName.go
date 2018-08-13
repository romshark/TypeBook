package rend

import "regexp"

var typeNameRule = regexp.MustCompile("^[A-Z][a-zA-Z]+$")

// verifyTypeName returns an error if the given type name
// violates type name rules, otherwise returns nil
func (d *Document) verifyTypeName(
	typeName string,
	declarationLocation string,
) (errors ModelErrors) {
	// Verify type name
	if !typeNameRule.MatchString(typeName) {
		errors.AddErrIllegalTypeName(
			typeName,
			declarationLocation,
		)
	}
	return errors
}
