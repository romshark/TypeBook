package rend

import (
	"time"
)

type RenderingStats struct {
	RenderingDur time.Duration
}

type InitStats struct {
	CompileTemplateDur time.Duration
}

type ModelInitStats struct {
	ProcessingDur   time.Duration
	VerificationDur time.Duration
}
