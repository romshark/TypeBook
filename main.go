package main

import (
	"bytes"
	"docbuilder/doc"
	"docbuilder/rend"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var inputFilePath = flag.String(
	"i",
	"./example.yml",
	"YAML Input file path",
)
var outputFilePath = flag.String(
	"o",
	"./compiled.html",
	"HTML Output file path",
)

func main() {
	flag.Parse()

	startProcess := time.Now()

	// Initialize renderer
	renderer, rendererInitStats, err := rend.New()
	if err != nil {
		log.Fatalf("Couldn't initialize renderer: %s", err)
	}

	document, docParsingStats, err := doc.NewFromFile(*inputFilePath)
	if err != nil {
		log.Fatalf("Couldn't read document: %s", err)
	}

	// Create document model
	documentModel, errorMessages, _, err := rend.NewModel(
		document,
		&rend.ModelInitOptions{
			CheckReferences: true,
		},
	)
	if err != nil {
		log.Fatalf("Couldn't initialize document model: %s", err)
	}

	if len(errorMessages) > 0 {
		fmt.Printf("%d errors:\n", len(errorMessages))
		for i, errMsg := range errorMessages {
			fmt.Printf(" %d: %s\n", i, errMsg)
		}
		os.Exit(1)
	}

	// Render the document
	var buf bytes.Buffer
	renderingStats, err := renderer.Render(documentModel, &buf)
	if err != nil {
		log.Fatalf("Couldn't render document: %s", err)
	}

	// Write to file
	startFinalizing := time.Now()
	if err := ioutil.WriteFile(
		*outputFilePath,
		buf.Bytes(),
		0644,
	); err != nil {
		log.Fatalf("Couldn't write rendered document to file: %s", err)
	}
	finalizingDur := time.Since(startFinalizing)

	totalProcessDur := time.Since(startProcess)

	fmt.Printf("Rendered to:         %s\n", *outputFilePath)
	fmt.Printf("Compiling Template:  %s\n", rendererInitStats.CompileTemplateDur)
	fmt.Printf("Parsing:             %s\n", docParsingStats.ParsingInputFileDur)
	fmt.Printf("Rendering:           %s\n", renderingStats.RenderingDur)
	fmt.Printf("Finalizing:          %s\n", finalizingDur)
	fmt.Printf("Total:               %s\n", totalProcessDur)
}
