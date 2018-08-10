package rend

type ModelInitOptions struct {
	CheckReferences bool
}

func DefaultModelInitOptions() *ModelInitOptions {
	return &ModelInitOptions{
		CheckReferences: true,
	}
}
