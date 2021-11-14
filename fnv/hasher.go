package fnv

const (
	longStrThreshold = 256
)

type Hasher struct{}

func (h Hasher) Sum32(s string) uint32 {
	if len(s) > longStrThreshold {
		return Hash32StringLong(s)
	}
	return Hash32String(s)
}

func (h Hasher) Sum64(s string) uint64 {
	if len(s) > longStrThreshold {
		return Hash64StringLong(s)
	}
	return Hash64String(s)
}

type BHasher struct{}

func (h BHasher) Sum32(s []byte) uint32 {
	if len(s) > longStrThreshold {
		return Hash32Long(s)
	}
	return Hash32(s)
}

func (h BHasher) Sum64(s []byte) uint64 {
	if len(s) > longStrThreshold {
		return Hash64Long(s)
	}
	return Hash64(s)
}
