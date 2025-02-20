package murmur

type byteseq interface {
	~[]byte | ~string
}

const (
	prime   = 0x5bd1e995
	prime64 = 0xc6a4a7935bd1e995
)

func Hash[T byteseq](p T, seed uint32) (h uint32) {
	ul := uint32(len(p))
	h = seed ^ ul
	var k uint32
	for ul >= 4 {
		k = leu32(p[:4])
		k *= prime
		k ^= k >> 24
		k *= prime
		h *= prime
		h ^= k
		p = p[4:]
		ul -= 4
	}

	switch ul {
	case 3:
		h ^= uint32(p[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(p[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(p[0])
		h *= prime
	}

	h ^= h >> 13
	h *= prime
	h ^= h >> 15

	return
}

func Hash64A[T byteseq](p T, seed uint64) (h uint64) {
	const r = 47

	ul := uint64(len(p))
	h = seed ^ ul*prime64

	l8 := int(ul / 8)
	var k uint64
	for i := 0; i < l8; i++ {
		i8 := i * 8
		k = leu64(p[i8 : i8+8])
		k *= prime64
		k ^= k >> r
		k *= prime64
		h ^= k
		h *= prime64
	}

	p = p[l8*8:]
	switch ul & 7 {
	case 7:
		h ^= uint64(p[6]) << 48
		fallthrough
	case 6:
		h ^= uint64(p[5]) << 40
		fallthrough
	case 5:
		h ^= uint64(p[4]) << 32
		fallthrough
	case 4:
		h ^= uint64(p[3]) << 24
		fallthrough
	case 3:
		h ^= uint64(p[2]) << 16
		fallthrough
	case 2:
		h ^= uint64(p[1]) << 8
		fallthrough
	case 1:
		h ^= uint64(p[0])
		h *= prime64
	}

	h ^= h >> r
	h *= prime64
	h ^= h >> r

	return
}

func Hash64B[T byteseq](p T, seed uint64) (h uint64) {
	const r = 24

	ul := uint64(len(p))
	h1, h2 := uint32(seed)^uint32(ul), uint32(0)
	var k1, k2 uint32
	for ul >= 8 {
		k1 = leu32(p[:4])
		k1 *= prime
		k1 ^= k1 >> r
		k1 *= prime
		h1 *= prime
		h1 ^= k1
		p = p[4:]
		ul -= 4

		k2 = leu32(p[:4])
		k2 *= prime
		k2 ^= k2 >> r
		k2 *= prime
		h2 *= prime
		h2 ^= k2
		p = p[4:]
		ul -= 4
	}

	if ul >= 4 {
		k1 = leu32(p[:4])
		k1 *= prime
		k1 ^= k1 >> r
		k1 *= prime
		h1 *= prime
		h1 ^= k1
		p = p[4:]
		ul -= 4
	}

	switch ul {
	case 3:
		h2 ^= uint32(p[2]) << 16
		fallthrough
	case 2:
		h2 ^= uint32(p[1]) << 8
		fallthrough
	case 1:
		h2 ^= uint32(p[0])
		h2 *= prime
	}

	h1 ^= h2 >> 18
	h1 *= prime
	h2 ^= h1 >> 22
	h2 *= prime
	h1 ^= h2 >> 17
	h1 *= prime
	h2 ^= h1 >> 19
	h2 *= prime

	h = uint64(h1)
	h = (h << 32) | uint64(h2)

	return

}

func Hash2A[T byteseq](p T, seed uint32) (h uint32) {
	const r = 24

	ul := uint32(len(p))
	h = seed
	var k uint32
	for ul >= 4 {
		k = leu32(p[:4])

		k *= prime
		k ^= k >> r
		k *= prime
		h *= prime
		h ^= k

		p = p[4:]
		ul -= 4
	}

	var t uint32 = 0
	switch ul {
	case 3:
		t ^= uint32(p[2]) << 16
		fallthrough
	case 2:
		t ^= uint32(p[1]) << 8
		fallthrough
	case 1:
		t ^= uint32(p[0])
	}

	t *= prime
	t ^= t >> r
	t *= prime
	h *= prime
	h ^= t
	var ll uint32
	ll = uint32(ul) * prime
	ll ^= ll >> r
	ll *= prime
	h *= prime
	h ^= ll

	h ^= h >> 13
	h *= prime
	h ^= h >> 15

	return
}

func HashNeutral2[T byteseq](p T, seed uint32) (h uint32) {
	ul := uint32(len(p))
	h = seed ^ ul
	var k uint32
	for ul >= 4 {
		k = leu32(p[:4])
		k *= prime
		k ^= k >> 24
		k *= prime
		h *= prime
		h ^= k
		p = p[4:]
		ul -= 4
	}

	switch ul {
	case 3:
		h ^= uint32(p[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(p[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(p[0])
		h *= prime
	}

	h ^= h >> 13
	h *= prime
	h ^= h >> 15

	return
}

func HashAligned2[T byteseq](p T, seed uint32) (h uint32) {
	const r = 24

	ul := uint32(len(p))
	h = seed ^ ul
	align := int(uint64(p[0])) & 3
	var k uint32
	if align != 0 && ul >= 4 {
		var t, d uint32 = 0, 0

		switch align {
		case 1:
			t |= uint32(p[2]) << 16
			fallthrough
		case 2:
			t |= uint32(p[1]) << 8
			fallthrough
		case 3:
			t |= uint32(p[0])
		}

		t <<= 8 * uint32(align)
		p = p[4-align:]
		ul -= uint32(4 - align)

		sl, sr := 8*(4-align), 8*align
		for ul >= 4 {
			d = leu32(p[:4])
			t = (t >> uint32(sr)) | (d << uint32(sl))

			k = t
			k *= prime
			k ^= k >> r
			k *= prime
			h *= prime
			h ^= k

			t = d
			p = p[4:]
			ul -= 4
		}

		d = 0

		if ul >= uint32(align) {
			switch align {
			case 3:
				d |= uint32(p[2]) << 16
				fallthrough
			case 2:
				d |= uint32(p[1]) << 8
				fallthrough
			case 1:
				d |= uint32(p[0])
			}

			k = (t >> uint32(sr)) | (d << uint32(sl))
			k *= prime
			k ^= k >> r
			k *= prime
			h *= prime
			h ^= k

			p = p[align:]
			ul -= uint32(align)

			switch ul {
			case 3:
				h ^= uint32(p[2]) << 16
				fallthrough
			case 2:
				h ^= uint32(p[1]) << 8
				fallthrough
			case 1:
				h ^= uint32(p[0])
				h *= prime
			}
		} else {
			switch ul {
			case 3:
				d |= uint32(p[2]) << 16
				fallthrough
			case 2:
				d |= uint32(p[1]) << 8
				fallthrough
			case 1:
				h |= uint32(p[0])
				fallthrough
			case 0:
				h ^= (t >> uint32(sr)) | (d << uint32(sl))
				h *= prime
			}
		}

		h ^= h >> 13
		h *= prime
		h ^= h >> 15

		return h
	} else {
		for ul >= 4 {
			k = leu32(p[:4])

			k *= prime
			k ^= k >> r
			k *= prime
			h *= prime
			h ^= k

			p = p[4:]
			ul -= 4
		}

		switch ul {
		case 3:
			h ^= uint32(p[2]) << 16
			fallthrough
		case 2:
			h ^= uint32(p[1]) << 8
			fallthrough
		case 1:
			h ^= uint32(p[0])
			h *= prime
		}

		h ^= h >> 13
		h *= prime
		h ^= h >> 15

		return
	}
}

func leu32[T byteseq](b T) uint32 {
	_ = b[3]
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func leu64[T byteseq](b T) uint64 {
	_ = b[7]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}
