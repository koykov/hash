package xxhash

type Hasher64[T byteseq] struct{}

func (Hasher64[T]) Sum64(p T) uint64 {
	return Hash64(p)
}
