//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/subrahamanyam341/protobuf/protobuf  --gogoslick_out=. dct.proto
package dct

import (
	"math/big"

	dct "github.com/subrahamanyam341/andes-core-go/data/dct/proto"
)

// New returns a new batch from given buffers
func New() *dct.ESDigitalToken {
	return &dct.ESDigitalToken{
		Value: big.NewInt(0),
	}
}
