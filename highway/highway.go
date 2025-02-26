package highway

const blocksz = 32

type byteseq interface {
	~[]byte | ~string
}

type Key [4]uint64

func Hash64[T byteseq](data T, key Key) uint64 {
	var state [16]uint64
	globalInit(&state, key)
	if n := len(data) & (^(blocksz - 1)); n > 0 {
		globalUpdate(&state, []byte(data[:n]))
		data = data[n:]
	}
	if len(data) > 0 {
		var block [blocksz]byte
		offset := copy(block[:], data)
		globalUpdateRemainder(&state, &block, offset)
	}
	return globalFinalize8(&state)
}

func Hash128[T byteseq](data T, key Key) [2]uint64 {
	var state [16]uint64
	globalInit(&state, key)
	if n := len(data) & (^(blocksz - 1)); n > 0 {
		globalUpdate(&state, []byte(data[:n]))
		data = data[n:]
	}
	if len(data) > 0 {
		var block [blocksz]byte
		offset := copy(block[:], data)
		globalUpdateRemainder(&state, &block, offset)
	}
	return globalFinalize16(&state)
}

func Hash256[T byteseq](data T, key Key) [4]uint64 {
	var state [16]uint64
	globalInit(&state, key)
	if n := len(data) & (^(blocksz - 1)); n > 0 {
		globalUpdate(&state, []byte(data[:n]))
		data = data[n:]
	}
	if len(data) > 0 {
		var block [blocksz]byte
		offset := copy(block[:], data)
		globalUpdateRemainder(&state, &block, offset)
	}
	return globalFinalize32(&state)
}
