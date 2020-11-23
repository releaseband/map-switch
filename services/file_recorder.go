package services

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type FileRecorder struct {
}

func NewRecorder() FileRecorder {
	return FileRecorder{}
}

func createFile(filename string) (*os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("create file failed: %w", err)
	}

	return f, nil
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func removeFile(filename string) error {
	return os.Remove(filename)
}

func getGenFilename(filename string) string {
	n := strings.Replace(filename, ".go", "", 1)
	return n + ".gen.go"
}

func (w FileRecorder) RecordToFile(filename string, t *template.Template, i interface{}) error {
	gFileName := getGenFilename(filename)

	if fileExist(gFileName) {
		if err := removeFile(gFileName); err != nil {
			return fmt.Errorf("remove file failed: %w", err)
		}
	}

	f, err := createFile(gFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	if err := t.Execute(f, i); err != nil {
		return fmt.Errorf("exucute template failed: %w", err)
	}

	return nil
}
