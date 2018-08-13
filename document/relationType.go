package document

import "fmt"

// RelationDirection represents a relation direction
type RelationDirection bool

const (
	// InboundRelation represents inbound relations
	InboundRelation RelationDirection = true

	// OutboundRelation represents outbound relations
	OutboundRelation RelationDirection = false
)

// String stringifies the value
func (rd RelationDirection) String() string {
	if rd == InboundRelation {
		return "inbound"
	}
	return "outbound"
}

// FromString initializes the value from bytes
func (rd *RelationDirection) FromBytes(buf []byte) error {
	return rd.FromString(string(buf))
}

// FromString initializes the value from a string
func (rd *RelationDirection) FromString(str string) error {
	if str == "inbound" {
		*rd = InboundRelation
		return nil
	} else if str == "outbound" {
		*rd = OutboundRelation
		return nil
	}
	return fmt.Errorf("invalid relation direction: '%s'", str)
}

// UnmarshalJSON implements the Go JSON unmarshaller interface
func (rd *RelationDirection) UnmarshalJSON(buf []byte) error {
	return rd.FromBytes(buf)
}

// UnmarshalYAML implements the go-YAML unmarshaller interface
func (rd *RelationDirection) UnmarshalYAML(
	unmarshal func(interface{}) error,
) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}
	return rd.FromString(val)
}
