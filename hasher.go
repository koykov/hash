package hash

type byteseq interface {
	~[]byte | ~string
}

// Hasher is a shorthand alias of Hasher64.
type Hasher[T byteseq] Hasher64[T]

// Hasher32 describes uint32 hash generator.
type Hasher32[T byteseq] interface {
	Sum32(T) uint32
}

// Hasher64 describes uint64 hash generator.
type Hasher64[T byteseq] interface {
	Sum64(T) uint64
}

// Hasher_x86_128 describes uint128([4]uint32) hash generator.
type Hasher_x86_128[T byteseq] interface {
	Sum128(T) [4]uint32
}

type Hasher_x64_128[T byteseq] interface {
	Sum128(T) [2]uint64
}

type Hasher_x64_256[T byteseq] interface {
	Sum256(T) [4]uint64
}
