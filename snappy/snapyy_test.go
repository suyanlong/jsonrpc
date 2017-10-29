package snappy

import (
	"testing"
	"fmt"
)

func TestCitaCompress(t *testing.T) {
	var d []byte
	d = append(d, 0xde, 0xad, 0xd0, 0x0d)
	//t.Log(CitaCompress(d))
	fmt.Println(CitaCompress(d))
}

func TestCitaDecompress(t *testing.T) {
	var d []byte
	d = append(d, 0xde, 0xad, 0xd0, 0x0d)
	//t.Log(CitaCompress(d))
	fmt.Println(d)
	fmt.Println(CitaDecompress(d))
}
