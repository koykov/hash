package v3

type Hasher32[T byteseq] struct {
	Seed uint32
}

func (h Hasher32[T]) Sum32(p T) uint32 {
	return Hash(p, h.Seed)
}

type Hasher_x86_128[T byteseq] struct {
	Seed uint32
}

func (h Hasher_x86_128[T]) Sum128(p T) [4]uint32 {
	return Hash_x86_128(p, h.Seed)
}

type Hasher_x64_128[T byteseq] struct {
	Seed uint64
}

func (h Hasher_x64_128[T]) Sum128(p T) [2]uint64 {
	return Hash_x64_128(p, h.Seed)
}
