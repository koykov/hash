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
