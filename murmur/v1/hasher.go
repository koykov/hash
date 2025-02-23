package murmur

type Hasher32[T byteseq] struct {
	Seed uint32
}

func (h Hasher32[T]) Sum32(p T) uint32 {
	return Hash(p, h.Seed)
}

type Hasher32Aligned[T byteseq] struct {
	Seed uint32
}

func (h Hasher32Aligned[T]) Sum32(p T) uint32 {
	return HashAligned(p, h.Seed)
}
