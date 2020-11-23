package internal

type MapData struct {
	KV   MapKeyValue
	Data map[string][]string
}

func NewMapData(keyVal MapKeyValue, data map[string][]string) MapData {
	return MapData{
		KV:   keyVal,
		Data: data,
	}
}
