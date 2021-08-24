package fnv

const (
	longStrThreshold = 256
)

type Hasher struct{}

func (h Hasher) Sum64(s string) uint64 {
	if len(s) > longStrThreshold {
		return Hash64StringLong(s)
	}
	return Hash64String(s)
}
