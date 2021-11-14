package hash

// Hasher describes uint64 hash generation from given string.
type Hasher Hasher64

// Hasher32 describes uint32 hash generation from given string.
type Hasher32 interface {
	Sum32(string) uint32
}

// Hasher64 describes uint64 hash generation from given string.
type Hasher64 interface {
	Sum64(string) uint64
}
