package document

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-yaml/yaml"
)

// New reads a document from a file located at the given path
func NewFromFile(inputFilePath string) (doc *Document, stats *Stats, err error) {
	// Read file
	fileContents, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read file: %s", err)
	}
	return New(fileContents)
}

// New reads a document from a buffer
func New(buf []byte) (doc *Document, stats *Stats, err error) {
	doc = &Document{}

	// Unmarshal YAML
	startParsingInputFile := time.Now()
	if err := yaml.Unmarshal(buf, doc); err != nil {
		return nil, nil, fmt.Errorf("couldn't parse file: %s", err)
	}
	parsingInputFileDur := time.Since(startParsingInputFile)

	return doc, &Stats{
		ParsingInputFileDur: parsingInputFileDur,
	}, nil
}
