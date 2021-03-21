package models

type Variant struct {
	Name    string
	MapData map[string][]string
}

func NewVariant(name string, mapData map[string][]string) *Variant {
	return &Variant{Name: name, MapData: mapData}
}

type FileDeclaration struct {
	PackageName string
	Vars        []Variant
}

func NewFileDeclaration(packageName string) *FileDeclaration {
	return &FileDeclaration{
		PackageName: packageName,
	}
}

func (d *FileDeclaration) AddVariants(variants []Variant) {
	if len(d.Vars) == 0 {
		d.Vars = variants
	} else {
		d.Vars = append(d.Vars, variants...)
	}
}
