package document

type ScalarType struct {
	Description string `yaml:"descr"`
}

type EnumType struct {
	Description string            `yaml:"descr"`
	Values      map[string]string `yaml:"values"`
}

type TypeField struct {
	Type        DataType `yaml:"type"`
	Description string   `yaml:"descr"`
	Nullable    bool     `yaml:"nullable"`
}

type CompositeTypeMetadata map[string]TypeField

type CompositeType struct {
	Description string                `yaml:"descr"`
	Metadata    CompositeTypeMetadata `yaml:"meta"`

	TotalMetadataFields uint32
}

type Document struct {
	Title            string                   `yaml:"title"`
	Author           string                   `yaml:"author"`
	Version          string                   `yaml:"version"`
	Description      string                   `yaml:"description"`
	ScalarTypes      map[string]ScalarType    `yaml:"scalar types"`
	EnumerationTypes map[string]EnumType      `yaml:"enumeration types"`
	CompositeTypes   map[string]CompositeType `yaml:"composite types"`
}
