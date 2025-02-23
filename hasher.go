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

// Hasher128x86 describes uint128([4]uint32) hash generator.
type Hasher128x86[T byteseq] interface {
	Sum128(T) [4]uint32
}

// Hasher128x64 describes uint128([2]uint64) hash generator.
type Hasher128x64[T byteseq] interface {
	Sum128(T) [2]uint64
}

// Hasher256x64 describes uint256([4]uint64) hash generator.
type Hasher256x64[T byteseq] interface {
	Sum256(T) [4]uint64
}
