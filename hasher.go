package hash

type Hasher interface {
	Sum64(string) uint64
}
