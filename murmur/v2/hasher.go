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

type Hasher32Neutral2[T byteseq] struct {
	Seed uint32
}

func (h Hasher32Neutral2[T]) Sum32(p T) uint32 {
	return HashNeutral2(p, h.Seed)
}

type Hasher32_2A[T byteseq] struct {
	Seed uint32
}

func (h Hasher32_2A[T]) Sum32(p T) uint32 {
	return Hash2A(p, h.Seed)
}

type Hasher64A[T byteseq] struct {
	Seed uint64
}

func (h Hasher64A[T]) Sum64(p T) uint64 {
	return Hash64A(p, h.Seed)
}

type Hasher64B[T byteseq] struct {
	Seed uint64
}

func (h Hasher64B[T]) Sum64(p T) uint64 {
	return Hash64B(p, h.Seed)
}
