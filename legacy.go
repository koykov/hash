package hash

// BHasher is a shorthand alias of BHasher64.
// DEPRECATED: use Hasher instead.
type BHasher BHasher64

// BHasher32 describes uint32 hash generation from given bytes.
// DEPRECATED: use Hasher32 instead.
type BHasher32 interface {
	Sum32([]byte) uint32
}

// BHasher64 describes uint64 hash generation from given bytes.
// DEPRECATED: use Hasher64 instead.
type BHasher64 interface {
	Sum64([]byte) uint64
}
