package substrate

import (
	"github.com/m0ssc0de/scale.go/source"
	"github.com/m0ssc0de/scale.go/types"
)

func RegCustomTypes(sourceCode []byte) {
	types.RegCustomTypes(source.LoadTypeRegistry(sourceCode))
}
