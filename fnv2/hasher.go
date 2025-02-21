package fnv

const (
	longStrThreshold = 256
)

type Hasher32[T byteseq] struct{}

func (h Hasher32[T]) Sum32(s T) uint32 {
	if len(s) > longStrThreshold {
		return Hash32Long(s)
	}
	return Hash32(s)
}

type Hasher64[T byteseq] struct{}

func (h Hasher64[T]) Sum64(s T) uint64 {
	if len(s) > longStrThreshold {
		return Hash64Long(s)
	}
	return Hash64(s)
}

type Hasher32a[T byteseq] struct{}

func (h Hasher32a[T]) Sum32(s T) uint32 {
	if len(s) > longStrThreshold {
		return Hash32aLong(s)
	}
	return Hash32a(s)
}

type Hasher64a[T byteseq] struct{}

func (h Hasher64a[T]) Sum64(s T) uint64 {
	if len(s) > longStrThreshold {
		return Hash64aLong(s)
	}
	return Hash64a(s)
}
