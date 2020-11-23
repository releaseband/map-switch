package internal

type SpecType struct {
	Imp      string
	Type     string
	FullType string
}

func NewSpecType(_import, _type string) SpecType {
	fullType := _type
	if len(_import) > 0 {
		fullType = _import + "." + _type
	}

	return SpecType{
		Imp:      _import,
		Type:     _type,
		FullType: fullType,
	}
}

type MapKeyValue struct {
	Key SpecType
	Val SpecType
}

func NewMapKeyVal(key, val SpecType) *MapKeyValue {
	return &MapKeyValue{
		Key: key,
		Val: val,
	}
}
