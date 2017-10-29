package snappy

import (
	"github.com/golang/snappy"
)

const CITACOMPRESSSIZE uint64 = 40 * 1024

func CitaDecompress(src []byte) []byte {
	var dst []byte
	if ret, err := snappy.Decode(dst, src); err == nil {
		return ret
	} else {
		return src
	}
}

func CitaCompress(src []byte) []byte {
	if uint64(len(src)) > CITACOMPRESSSIZE {
		var dst []byte
		return snappy.Encode(dst, src)
	} else {
		return src
	}
}
