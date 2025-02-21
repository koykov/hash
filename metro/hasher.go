package metro

type Hasher64[T byteseq] struct {
	Seed uint64
}

func (h Hasher64[T]) Sum64(p T) uint64 {
	return Hash64(p, h.Seed)
}

type Hasher128 struct {
	Seed uint64
}

func (h Hasher128) Sum128(p []byte) [2]uint64 {
	return Hash128(p, h.Seed)
}
