package murmur

type Hasher128x86[T byteseq] struct {
	Seed uint32
}

func (h Hasher128x86[T]) Sum128(p T) [4]uint32 {
	return Hash128x86(p, h.Seed)
}

type Hasher128x64[T byteseq] struct {
	Seed uint64
}

func (h Hasher128x64[T]) Sum128(p T) [2]uint64 {
	return Hash128x64(p, h.Seed)
}
