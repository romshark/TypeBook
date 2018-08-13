package document

import (
	"fmt"
	"regexp"
)

var plainDataTypePattern = regexp.MustCompile("^[A-Z][a-zA-Z_]*$")
var listDataTypePattern = regexp.MustCompile("^List\\s*<([A-Z][a-zA-Z_]*)>$")

type DataType struct {
	Name   string
	IsList bool
}

func (d DataType) String() string {
	if d.IsList {
		return fmt.Sprintf("List<%s>", d.Name)
	}
	return d.Name
}

func (d *DataType) FromBytes(buf []byte) error {
	str := string(buf)
	if listDataTypePattern.Match(buf) {
		d.Name = listDataTypePattern.FindStringSubmatch(str)[1]
		d.IsList = true
		return nil
	} else if plainDataTypePattern.Match(buf) {
		d.Name = str
		d.IsList = false
		return nil
	}
	return fmt.Errorf("invalid data type: '%s'", string(buf))
}

func (d *DataType) FromString(str string) error {
	return d.FromBytes([]byte(str))
}

// UnmarshalJSON implements the Go JSON interface
func (d *DataType) UnmarshalJSON(buf []byte) error {
	return d.FromBytes(buf)
}

// UnmarshalYAML implements the go-YAML interface
func (d *DataType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}
	return d.FromString(val)
}
