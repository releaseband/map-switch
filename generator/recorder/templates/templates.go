package templates

const Template = `//CODE GENERATED AUTOMATICALLY. DO NOT EDIT.
package {{.PackageName}}

{{range $i, $result := $.Vars}}
func {{$result.Name}} (s uint32, count uint8) uint16 {
	switch s {
		{{range $key, $val := $result.MapData }}
			case {{$key}}:
				switch count { {{range $j, $v := $val}}
					case {{increment $j}}: return {{$v}}{{end}}
				}{{end}}}
	return 0
}

{{end}}
`
