package rend

// registerScalarType registers a new scalar type
func (d *Document) registerScalarType(
	forwardDeclared Types,
	typeName,
	description string,
) (errors ModelErrors) {
	// Verify type name
	errors.Add(d.verifyTypeName(
		typeName,
		"scalar type declaration",
	)...)
	if errors.HasErrors() {
		// Don't evaluate further in case of illegal name
		return errors
	}

	// verify type
	errors.Add(d.verifyType(
		forwardDeclared,
		typeName,
		"scalar type declaration", // error location
	)...)
	if errors.HasErrors() {
		// Don't evaluate further in case of invalid type
		return errors
	}

	newType := &ScalarType{
		TypeName:    typeName,
		Description: description,
	}

	// Successfully register the new type
	d.ScalarTypes[newType.TypeName] = newType
	d.Types[newType.TypeName] = newType
	return nil
}

// RegisterScalarType registers a new scalar type
func (d *Document) RegisterScalarType(
	typeName,
	description string,
) ModelErrors {
	return d.registerScalarType(nil, typeName, description)
}
