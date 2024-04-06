package server

import (
	"errors"
	"os"
	"path/filepath"
)

type configureRequest struct {
	WriteFiles []writeFile `json:"write-files"`
}

func (c *configureRequest) exec() error {
	var errs error
	for _, wf := range c.WriteFiles {
		if err := wf.exec(); err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}

type writeFile struct {
	Filepath string `json:"filepath"`
	Content  string `json:"content"`
}

func (wf *writeFile) exec() error {
	dir := filepath.Dir(wf.Filepath)
	if err := os.MkdirAll(dir, 0644); err != nil {
		return err
	}
	return os.WriteFile(wf.Filepath, []byte(wf.Content), 0644)
}

type errorResponse struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}
