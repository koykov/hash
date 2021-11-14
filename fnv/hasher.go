package fnv

const (
	longStrThreshold = 256
)

type Hasher Hasher64

type Hasher32 struct{}

func (h Hasher32) Sum32(s string) uint32 {
	if len(s) > longStrThreshold {
		return Hash32StringLong(s)
	}
	return Hash32String(s)
}

type Hasher64 struct{}

func (h Hasher64) Sum64(s string) uint64 {
	if len(s) > longStrThreshold {
		return Hash64StringLong(s)
	}
	return Hash64String(s)
}
