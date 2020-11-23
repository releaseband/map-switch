package helpers

import (
	"strings"

	"github.com/releaseband/map-switch/analysis/internal"
)

func getFullPath(need string, allImports []string) (string, bool) {
	for _, imp := range allImports {
		if strings.Contains(imp, need) {
			return imp, true
		}
	}

	return "", false
}

func GetNeedImports(kv internal.MapKeyValue, allImports []string) []string {
	imps := make([]string, 0, 2)

	if len(kv.Key.Imp) > 0 {
		keyImpPath, ok := getFullPath(kv.Key.Imp, allImports)
		if ok {
			imps = append(imps, keyImpPath)
		}
	}
	if len(kv.Val.Imp) > 0 {
		valImpPath, ok := getFullPath(kv.Val.Imp, allImports)
		if ok {
			imps = append(imps, valImpPath)
		}
	}

	return imps
}
