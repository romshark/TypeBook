package rend

import (
	"fmt"
	"io"
	"time"
)

// Render renders the given document model to a buffer
func (r *Renderer) Render(
	model *Document,
	outBuffer io.Writer,
) (*RenderingStats, error) {
	// Render
	startRendering := time.Now()

	if err := r.template.Execute(outBuffer, model); err != nil {
		return nil, fmt.Errorf("couldn't render to template: %s", err)
	}

	renderingDur := time.Since(startRendering)

	return &RenderingStats{
		RenderingDur: renderingDur,
	}, nil
}
