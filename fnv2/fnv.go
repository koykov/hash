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

// Hash32Long is a loop rolled version of FNV-1 32 for very long input.
func Hash32Long[T byteseq](p T) uint32 {
	n := len(p)
	if n == 0 {
		return 0
	}
	h := offset32
	_ = p[n-1]
	for len(p) >= 8 {
		h *= prime32
		h ^= uint32(p[0])
		h *= prime32
		h ^= uint32(p[1])
		h *= prime32
		h ^= uint32(p[2])
		h *= prime32
		h ^= uint32(p[3])
		h *= prime32
		h ^= uint32(p[4])
		h *= prime32
		h ^= uint32(p[5])
		h *= prime32
		h ^= uint32(p[6])
		h *= prime32
		h ^= uint32(p[7])
		p = p[8:]
	}

	if len(p) >= 4 {
		h *= prime32
		h ^= uint32(p[0])
		h *= prime32
		h ^= uint32(p[1])
		h *= prime32
		h ^= uint32(p[2])
		h *= prime32
		h ^= uint32(p[3])
		p = p[4:]
	}

	if len(p) >= 2 {
		h *= prime32
		h ^= uint32(p[0])
		h *= prime32
		h ^= uint32(p[1])
		p = p[2:]
	}

	if len(p) > 0 {
		h *= prime32
		h ^= uint32(p[0])
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

// Hash32aLong is a loop rolled version of FNV-1a 32 for very long input.
func Hash32aLong[T byteseq](p T) uint32 {
	n := len(p)
	if n == 0 {
		return 0
	}
	h := offset32

	for len(p) >= 8 {
		h = (h ^ uint32(p[0])) * prime32
		h = (h ^ uint32(p[1])) * prime32
		h = (h ^ uint32(p[2])) * prime32
		h = (h ^ uint32(p[3])) * prime32
		h = (h ^ uint32(p[4])) * prime32
		h = (h ^ uint32(p[5])) * prime32
		h = (h ^ uint32(p[6])) * prime32
		h = (h ^ uint32(p[7])) * prime32
		p = p[8:]
	}

	if len(p) >= 4 {
		h = (h ^ uint32(p[0])) * prime32
		h = (h ^ uint32(p[1])) * prime32
		h = (h ^ uint32(p[2])) * prime32
		h = (h ^ uint32(p[3])) * prime32
		p = p[4:]
	}

	if len(p) >= 2 {
		h = (h ^ uint32(p[0])) * prime32
		h = (h ^ uint32(p[1])) * prime32
		p = p[2:]
	}

	if len(p) > 0 {
		h = (h ^ uint32(p[0])) * prime32
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

// Hash64Long is a loop rolled version of FNV-1 64 for very long input.
func Hash64Long[T byteseq](p T) uint64 {
	n := len(p)
	if n == 0 {
		return 0
	}
	h := offset64
	_ = p[n-1]
	for len(p) >= 8 {
		h *= prime64
		h ^= uint64(p[0])
		h *= prime64
		h ^= uint64(p[1])
		h *= prime64
		h ^= uint64(p[2])
		h *= prime64
		h ^= uint64(p[3])
		h *= prime64
		h ^= uint64(p[4])
		h *= prime64
		h ^= uint64(p[5])
		h *= prime64
		h ^= uint64(p[6])
		h *= prime64
		h ^= uint64(p[7])
		p = p[8:]
	}

	if len(p) >= 4 {
		h *= prime64
		h ^= uint64(p[0])
		h *= prime64
		h ^= uint64(p[1])
		h *= prime64
		h ^= uint64(p[2])
		h *= prime64
		h ^= uint64(p[3])
		p = p[4:]
	}

	if len(p) >= 2 {
		h *= prime64
		h ^= uint64(p[0])
		h *= prime64
		h ^= uint64(p[1])
		p = p[2:]
	}

	if len(p) > 0 {
		h *= prime64
		h ^= uint64(p[0])
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

// Hash64aLong is a loop rolled version of FNV-1a 64 for very long input.
func Hash64aLong[T byteseq](p T) uint64 {
	n := len(p)
	if n == 0 {
		return 0
	}
	h := offset64
	_ = p[n-1]
	for len(p) >= 8 {
		h = (h ^ uint64(p[0])) * prime64
		h = (h ^ uint64(p[1])) * prime64
		h = (h ^ uint64(p[2])) * prime64
		h = (h ^ uint64(p[3])) * prime64
		h = (h ^ uint64(p[4])) * prime64
		h = (h ^ uint64(p[5])) * prime64
		h = (h ^ uint64(p[6])) * prime64
		h = (h ^ uint64(p[7])) * prime64
		p = p[8:]
	}

	if len(p) >= 4 {
		h = (h ^ uint64(p[0])) * prime64
		h = (h ^ uint64(p[1])) * prime64
		h = (h ^ uint64(p[2])) * prime64
		h = (h ^ uint64(p[3])) * prime64
		p = p[4:]
	}

	if len(p) >= 2 {
		h = (h ^ uint64(p[0])) * prime64
		h = (h ^ uint64(p[1])) * prime64
		p = p[2:]
	}

	if len(p) > 0 {
		h = (h ^ uint64(p[0])) * prime64
	}

	return h
}
