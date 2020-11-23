package templates

const MapImpl = `//CODE GENERATED AUTOMATICALLY. DO NOT EDIT.
package {{.Package}}

import ({{range $k, $val := $.Imports}}
{{$val}}{{end}}	
)
{{range $i, $result := $.R}}
func {{$result.FuncName}} (s {{$result.KeyType}}, count {{$result.CountType}}) {{$result.ValType}} {
	switch s {
		{{range $key, $val := $result.Map.Data }}
			case {{$key}}:
				switch count { {{range $j, $v := $val}}
					case {{inc $j}}: return {{$v}}{{end}}
				}{{end}}}
	return 0
}

{{end}}
`
