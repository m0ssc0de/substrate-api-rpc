package substrate

import (
	"fmt"
	scale "github.com/m0ssc0de/scale.go"
	"github.com/m0ssc0de/scale.go/types"
	"github.com/m0ssc0de/substrate-api-rpc/metadata"
	"github.com/m0ssc0de/substrate-api-rpc/util"
)

// Event decode
func DecodeEvent(rawList string, metadata *metadata.Instant, spec int) (r interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovering from panic in DecodeEvent error is: %v \n", r)
		}
	}()
	m := types.MetadataStruct(*metadata)
	e := scale.EventsDecoder{}
	option := types.ScaleDecoderOption{Metadata: &m, Spec: spec}
	e.Init(types.ScaleBytes{Data: util.HexToBytes(rawList)}, &option)
	e.Process()
	return e.Value, nil
}
