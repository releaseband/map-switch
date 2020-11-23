package analysis

import (
	"fmt"
	"text/template"

	"github.com/releaseband/map-switch/templates"
)

type TemplateRecorder interface {
	RecordToFile(filename string, t *template.Template, i interface{}) error
}

func prepareTemplate(templatePath string) *template.Template {

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}

	return template.Must(template.New("Map to Switch case").Funcs(funcMap).Parse(templatePath))
}

func GenerateMapByString(w TemplateRecorder, data MapParams) error {
	result, err := analysisFileByMap(data)
	if err != nil {
		return err
	}

	t := prepareTemplate(templates.MapImpl)

	if err := w.RecordToFile(data.FilePath, t, result); err != nil {
		return fmt.Errorf("write to file failed: %w", err)
	}

	return nil
}
