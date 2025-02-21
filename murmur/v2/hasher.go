package murmur

type Hasher[T byteseq] struct {
	Seed uint32
}

func (h Hasher[T]) Sum32(p T) uint32 {
	return Hash(p, h.Seed)
}

type HasherAligned[T byteseq] struct {
	Seed uint32
}

func (h HasherAligned[T]) Sum32(p T) uint32 {
	return HashAligned(p, h.Seed)
}

type HasherNeutral2[T byteseq] struct {
	Seed uint32
}

func (h HasherNeutral2[T]) Sum32(p T) uint32 {
	return HashNeutral2(p, h.Seed)
}

type Hasher2A[T byteseq] struct {
	Seed uint32
}

func (h Hasher2A[T]) Sum32(p T) uint32 {
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
