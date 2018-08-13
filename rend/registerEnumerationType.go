package rend

// registerEnumerationType registers a new enumeration type
func (d *Document) registerEnumerationType(
	forwardDeclared Types,
	typeName,
	description string,
	values map[string]string,
) (errors ModelErrors) {
	// Verify type name
	errors.Add(d.verifyTypeName(
		typeName,
		"enumeration type declaration",
	)...)
	if errors.HasErrors() {
		// Don't evaluate further in case of illegal name
		return errors
	}

	// Verify type
	errors.Add(d.verifyType(
		forwardDeclared,
		typeName,
		"enumeration type declaration", // error location
	)...)
	if errors.HasErrors() {
		// Don't evaluate further in case of invalid type
		return errors
	}

	newType := &EnumerationType{
		TypeName:    typeName,
		Description: description,
		Values:      values,
	}

	// Successfully register the new type
	d.EnumerationTypes[newType.TypeName] = newType
	d.Types[newType.TypeName] = newType
	return nil
}

// RegisterEnumerationType registers a new enumeration type
func (d *Document) RegisterEnumerationType(
	typeName,
	description string,
	values map[string]string,
) (errors ModelErrors) {
	// Copy the key-value pairs
	valuesCopy := make(EnumerationValues, len(values))
	for item, val := range values {
		valuesCopy[item] = val
	}

	return d.registerEnumerationType(nil, typeName, description, valuesCopy)
}
