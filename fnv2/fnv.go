package fnv

type byteseq interface {
	~[]byte | ~string
}

const (
	offset32 = uint32(2166136261)
	offset64 = uint64(14695981039346656037)
	prime32  = uint32(16777619)
	prime64  = uint64(1099511628211)
)

// Hash32 calculates FNV-1 32 hash fast.
func Hash32[T byteseq](p T) uint32 {
	n := len(p)
	if n == 0 {
		return 0
	}
	h := offset32
	_ = p[n-1]
	for i := 0; i < n; i++ {
		h *= prime32
		h ^= uint32(p[i])
	}
	return h
}

// Hash32a calculates FNV-1a 32 hash fast.
func Hash32a[T byteseq](p T) uint32 {
	n := len(p)
	if n == 0 {
		return 0
	}
	h := offset32
	_ = p[n-1]
	for i := 0; i < n; i++ {
		h = (h ^ uint32(p[i])) * prime32
	}
	return h
}

// Hash64 calculates FNV-1 64 hash fast.
func Hash64[T byteseq](p T) uint64 {
	n := len(p)
	if n == 0 {
		return 0
	}
	h := offset64
	_ = p[n-1]
	for i := 0; i < n; i++ {
		h *= prime64
		h ^= uint64(p[i])
	}
	return h
}

// Hash64a calculates FNV-1a 64 hash fast.
func Hash64a[T byteseq](p T) uint64 {
	n := len(p)
	if n == 0 {
		return 0
	}
	h := offset64
	_ = p[n-1]
	for i := 0; i < n; i++ {
		h = (h ^ uint64(p[i])) * prime64
	}
	return h
}
