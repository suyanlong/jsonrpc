package libproto_test

import (
	"testing"
	"jsonrpc/libproto"
)

func TestCmdId(t *testing.T) {
	println(libproto.CmdId(12, 12))
}
