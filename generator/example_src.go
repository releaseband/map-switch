package generator

import "fmt"

const (
	A = 0
	B = 1
)

func print() {
	fmt.Println("111")
}

//map_gen:name=sList
var Param1 = map[uint32][]int{
	3: {1, 2, 3, 4},
}

//map_gen:name=lList
var Param2 = map[uint32][]int{
	A: {10, 20, 3, 4},
	B: {11, 22, 3, 4},
}
