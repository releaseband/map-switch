package analysis

import "strings"

type MapParams struct {
	FilePath string
}

func NewMapParams(filepath string) MapParams {
	return MapParams{
		FilePath: strings.Replace(filepath, ".go", "", 1) + ".go",
	}
}
