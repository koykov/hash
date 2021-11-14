package hash

// Hasher is a shorthand alias of Hasher64.
type Hasher Hasher64

// Hasher32 describes uint32 hash generation from given string.
type Hasher32 interface {
	Sum32(string) uint32
}

// Hasher64 describes uint64 hash generation from given string.
type Hasher64 interface {
	Sum64(string) uint64
}

// BHasher is a shorthand alias of BHasher64.
type BHasher BHasher64

// BHasher32 describes uint32 hash generation from given bytes.
type BHasher32 interface {
	Sum32([]byte) uint32
}

// BHasher64 describes uint64 hash generation from given bytes.
type BHasher64 interface {
	Sum64([]byte) uint64
}
