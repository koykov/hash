package fnv

const (
	longStrThreshold = 256
)

type Hasher[T byteseq] struct{}

func (h Hasher[T]) Sum32(s T) uint32 {
	if len(s) > longStrThreshold {
		return Hash32Long(s)
	}
	return Hash32(s)
}

func (h Hasher[T]) Sum32a(s T) uint32 {
	if len(s) > longStrThreshold {
		return Hash32aLong(s)
	}
	return Hash32a(s)
}

func (h Hasher[T]) Sum64(s T) uint64 {
	if len(s) > longStrThreshold {
		return Hash64Long(s)
	}
	return Hash64(s)
}

func (h Hasher[T]) Sum64a(s T) uint64 {
	if len(s) > longStrThreshold {
		return Hash64aLong(s)
	}
	return Hash64a(s)
}
