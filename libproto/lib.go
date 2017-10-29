package libproto

func CmdId(submodule, topic uint32) uint32 {
	return (submodule << 16) + topic
}

func DeCmdId(cmdId uint32) (uint32, uint32) {
	submodule := cmdId >> 16
	var sub = submodule
	var tmp = submodule << 16
	topic := cmdId - tmp
	return sub, topic
}

