package fnv

type Hasher32[T byteseq] struct{}

func (h Hasher32[T]) Sum32(s T) uint32 {
	return Hash32(s)
}

type Hasher64[T byteseq] struct{}

func (h Hasher64[T]) Sum64(s T) uint64 {
	return Hash64(s)
}

type Hasher32a[T byteseq] struct{}

func (h Hasher32a[T]) Sum32(s T) uint32 {
	return Hash32a(s)
}

type Hasher64a[T byteseq] struct{}

func (h Hasher64a[T]) Sum64(s T) uint64 {
	return Hash64a(s)
}
