package highway

type Hasher64[T byteseq] struct {
	Key Key
}

func (h Hasher64[T]) Sum64(p T) uint64 {
	return Hash64(p, h.Key)
}

type Hasher128 struct {
	Key Key
}

func (h Hasher128) Sum128(p []byte) [2]uint64 {
	return Hash128(p, h.Key)
}

type Hasher256 struct {
	Key Key
}

func (h Hasher256) Sum256(p []byte) [4]uint64 {
	return Hash256(p, h.Key)
}
