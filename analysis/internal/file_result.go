package internal

type FileResult struct {
	Package string
	Imports []string
	R       []Result
}

func getUniqueImports(r []Result) []string {
	var imports []string
	if len(r) > 0 {
		imports = r[0].Imports
	}

	for i := 1; i < len(r); i++ {
		for _, imp := range r[i].Imports {
			ok := true
			for _, reqImp := range imports {
				if imp == reqImp {
					ok = false
					break
				}
			}

			if ok {
				imports = append(imports, imp)
			}
		}
	}

	return imports
}

func NewFileResult(_package string, results []Result) *FileResult {
	return &FileResult{
		Package: _package,
		Imports: getUniqueImports(results),
		R:       results,
	}
}
