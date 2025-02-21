package murmur

type Hasher[T byteseq] struct {
	Seed uint32
}

func (h Hasher[T]) Sum32(p T) uint32 {
	return Hash(p, h.Seed)
}

func (h Hasher[T]) Sum64(p T) uint64 {
	return uint64(Hash(p, h.Seed))
}

type HasherAligned[T byteseq] struct {
	Seed uint32
}

func (h HasherAligned[T]) Sum32(p T) uint32 {
	return HashAligned(p, h.Seed)
}

func (h HasherAligned[T]) Sum64(p T) uint64 {
	return uint64(HashAligned(p, h.Seed))
}

type HasherNeutral2[T byteseq] struct {
	Seed uint32
}

func (h HasherNeutral2[T]) Sum32(p T) uint32 {
	return HashNeutral2(p, h.Seed)
}

func (h HasherNeutral2[T]) Sum64(p T) uint64 {
	return uint64(HashNeutral2(p, h.Seed))
}

type Hasher2A[T byteseq] struct {
	Seed uint32
}

func (h Hasher2A[T]) Sum32(p T) uint32 {
	return Hash2A(p, h.Seed)
}

func (h Hasher2A[T]) Sum64(p T) uint64 {
	return uint64(Hash2A(p, h.Seed))
}

type Hasher64A[T byteseq] struct {
	Seed uint64
}

func (h Hasher64A[T]) Sum32(p T) uint32 {
	return uint32(Hash64A(p, h.Seed))
}

func (h Hasher64A[T]) Sum64(p T) uint64 {
	return Hash64A(p, h.Seed)
}

type Hasher64B[T byteseq] struct {
	Seed uint64
}

func (h Hasher64B[T]) Sum32(p T) uint32 {
	return uint32(Hash64B(p, h.Seed))
}

func (h Hasher64B[T]) Sum64(p T) uint64 {
	return Hash64B(p, h.Seed)
}
