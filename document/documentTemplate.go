package document

type ScalarType struct {
	Description string `yaml:"description"`
}

type EnumerationType struct {
	Description string            `yaml:"description"`
	Values      map[string]string `yaml:"values"`
}

type TypeField struct {
	Type        DataType `yaml:"type"`
	Description string   `yaml:"description"`
	Nullable    bool     `yaml:"nullable"`
}

// Metadata maps the field names to a metadata field
type Metadata map[string]TypeField

// EntityRelations maps a related relation name to the relation metadata
type EntityRelations map[string]EntityRelation

type EntityRelation struct {
	Description string            `yaml:"description"`
	Metadata    Metadata          `yaml:"meta"`
	Type        string            `yaml:"type"`
	Direction   RelationDirection `yaml:"direction"`
	RelatedType string            `yaml:"related type"`
}

type CompositeType struct {
	Description string   `yaml:"description"`
	Metadata    Metadata `yaml:"meta"`
}

type EntityType struct {
	Description string          `yaml:"description"`
	Metadata    Metadata        `yaml:"meta"`
	Relations   EntityRelations `yaml:"relations"`
}

type Document struct {
	Title            string                     `yaml:"title"`
	Author           string                     `yaml:"author"`
	Version          string                     `yaml:"version"`
	Description      string                     `yaml:"description"`
	ScalarTypes      map[string]ScalarType      `yaml:"scalar types"`
	EnumerationTypes map[string]EnumerationType `yaml:"enumeration types"`
	CompositeTypes   map[string]CompositeType   `yaml:"composite types"`
	EntityTypes      map[string]EntityType      `yaml:"entity types"`
}
