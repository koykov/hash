package v3

type Hasher[T byteseq] struct {
	Seed uint32
}

func (h Hasher[T]) Sum32(p T) uint32 {
	return Hash(p, h.Seed)
}

func (h Hasher[T]) Sum64(p T) uint64 {
	return uint64(Hash(p, h.Seed))
}

type Hasher_x86_128[T byteseq] struct {
	Seed uint32
}

func (h Hasher_x86_128[T]) Sum32_x86_128(p T) [4]uint32 {
	return Hash_x86_128(p, h.Seed)
}

func (h Hasher_x86_128[T]) Sum64_x64_128(p T) [4]uint64 {
	d := Hash_x86_128(p, h.Seed)
	return [4]uint64{uint64(d[0]), uint64(d[1]), uint64(d[2]), uint64(d[3])}
}

type Hasher_x64_128[T byteseq] struct {
	Seed uint64
}

func (h Hasher_x64_128[T]) Sum32_x64_128(p T) [2]uint32 {
	d := Hash_x64_128(p, h.Seed)
	return [2]uint32{uint32(d[0]), uint32(d[1])}
}

func (h Hasher_x64_128[T]) Sum64_x64_128(p T) [2]uint64 {
	return Hash_x64_128(p, h.Seed)
}
