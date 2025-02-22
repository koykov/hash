package highway

func globalInit(state *[16]uint64, key Key) {
	initPortable(state, key)
}

func globalUpdate(state *[16]uint64, data []byte) {
	updatePortable(state, data)
}

func globalUpdateRemainder(state *[16]uint64, buf *[32]byte, off int) {
	updateRemainderPortable(state, buf, off)
}

func globalFinalize8(state *[16]uint64) uint64 {
	return finalize8Portable(state)
}

func globalFinalize16(state *[16]uint64) [2]uint64 {
	return finalize16Portable(state)
}

func globalFinalize32(state *[16]uint64) [4]uint64 {
	return finalize32Portable(state)
}
