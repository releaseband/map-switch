package internal

type Result struct {
	Imports   []string
	FuncName  string
	KeyType   string
	ValType   string
	Map       MapData
	CountType string
	Condition string
}

func NewResult(countType string) *Result {
	condition := ""
	switch countType {
	case "int8", "int16", "int32", "int64":
		condition = "&& count >= 0"
	}

	return &Result{
		Condition: condition,
		CountType: countType,
	}
}

func (r *Result) SetMapData(md MapData) {
	r.KeyType = md.KV.Key.FullType
	r.ValType = md.KV.Val.FullType

	r.Map = md
}

func (r *Result) SetImports(imps []string) {
	r.Imports = imps
}
