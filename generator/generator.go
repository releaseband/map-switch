package generator

import (
	"github.com/releaseband/map-switch/generator/recorder"

	"github.com/releaseband/map-switch/generator/parser"
)

func Run(path string) error {
	fileDecl, err := parser.ParseFile(path)
	if err != nil {
		return err
	}

	if err := recorder.RecordMap(fileDecl); err != nil {
		return err
	}

	return nil
}
