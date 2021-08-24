package hash

// Hasher describes uint64 hash generation from given string.
type Hasher interface {
	Sum64(string) uint64
}
