package murmur

type Hasher64x86[T byteseq] struct {
	Seed uint32
}

func (h Hasher64x86[T]) Sum64(p T) uint64 {
	r := Hash128x86(p, h.Seed)
	return uint64(r[0]) + uint64(r[1]) + uint64(r[2]) + uint64(r[3])
}

type Hasher64x64[T byteseq] struct {
	Seed uint64
}

func (h Hasher64x64[T]) Sum64(p T) uint64 {
	r := Hash128x64(p, h.Seed)
	return r[0] + r[1]
}
