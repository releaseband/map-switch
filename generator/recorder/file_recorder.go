package recorder

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/releaseband/map-switch/generator/parser/models"
	"github.com/releaseband/map-switch/generator/recorder/templates"
)

const (
	postfix = ".gen.go"
)

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
	return strings.Replace(filename, ".go", "", 1) + postfix
}

func recordToFile(filename string, t *template.Template, i interface{}) error {
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

func makeTemplate(tempData string) *template.Template {
	return template.Must(template.New("map => switch").
		Funcs(template.FuncMap{
			"increment": func(i int) int {
				return i + 1
			},
		}).Parse(tempData))
}

func RecordMap(fd *models.FileDeclaration) error {
	if err := recordToFile(fd.PackageName, makeTemplate(templates.Template), fd); err != nil {
		return fmt.Errorf("recordToFile: %w", err)
	}

	return nil
}
