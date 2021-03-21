package generator

import (
	"github.com/releaseband/map-switch/generator/parser"
	"testing"
)

func TestParse(t *testing.T) {
	const path = "./example_src.go"
	fd, err := parser.ParseFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if fd.PackageName != "generator" {
		t.Fatal("package name invalid")
	}

	if len(fd.Vars) != 2 {
		t.Fatal("maps count invalid")
	}

	sList := fd.Vars[0]
	lList := fd.Vars[1]

	if sList.Name != "sList" || lList.Name != "lList"{
		t.Fatal("map name invalid")
	}

	if len(sList.MapData) != 1 {
		t.Fatal("sList map values invalid")
	}

	if len(lList.MapData) != 2 {
		t.Fatal("lList map value invalid")
	}
}

func TestGenerate(t *testing.T) {
	const path = "./example_src.go"

	err := Run(path)
	if err != nil {
		t.Fatal(err)
	}
}

