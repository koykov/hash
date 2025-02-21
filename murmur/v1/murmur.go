package murmur

type byteseq interface {
	~[]byte | ~string
}

const prime = 0xc6a4a793

func Hash[T byteseq](p T, seed uint32) (h uint32) {
	ul := uint32(len(p))
	h = seed ^ ul*prime
	for ul >= 4 {
		k := leu32(p[:4])
		h += k
		h *= prime
		h ^= h >> 16
		p = p[4:]
		ul -= 4
	}

	switch ul {
	case 3:
		h += uint32(p[2]) << 16
		fallthrough
	case 2:
		h += uint32(p[1]) << 8
		fallthrough
	case 1:
		h += uint32(p[0])
		h *= prime
		h ^= h >> 16
	}

	h *= prime
	h ^= h >> 10
	h *= prime
	h ^= h >> 17

	return
}

func HashAligned[T byteseq](p T, seed uint32) (h uint32) {
	ul := uint32(len(p))
	h = seed ^ (ul * prime)
	a := uint64(p[0] & 3)
	if a > 0 && ul >= 4 {
		var t, d uint32
		switch a {
		case 1:
			t |= uint32(p[2]) << 16
			fallthrough
		case 2:
			t |= uint32(p[1]) << 8
			fallthrough
		case 3:
			t |= uint32(p[0])
		}
		t <<= a * 8
		p = p[4-a:]
		ul -= uint32(4 - a)

		sl, sr := 8*(4-a), 8*a
		for ul >= 4 {
			d = leu32(p[:4])
			t = (t >> uint32(sr)) | (d << uint32(sl))
			h += t
			h *= prime
			h ^= h >> 16
			t = d

			p = p[4:]
			ul -= 4
		}
		pack := a
		if uint64(ul) < a {
			pack = uint64(ul)
		}
		d = 0
		switch pack {
		case 3:
			d |= uint32(p[2]) << 16
			fallthrough
		case 2:
			d |= uint32(p[1]) << 8
			fallthrough
		case 1:
			d |= uint32(p[0])
			fallthrough
		case 0:
			h += (t >> sr) | (d << sl)
			h *= prime
			h ^= h >> 16
		}
		p = p[pack:]
		ul -= uint32(pack)
	} else {
		for ul >= 4 {
			h += leu32(p[:4])
			p = p[4:]
			ul -= 4
		}
	}

	switch ul {
	case 3:
		h += uint32(p[2]) << 16
		fallthrough
	case 2:
		h += uint32(p[1]) << 8
		fallthrough
	case 1:
		h += uint32(p[0])
		h *= prime
		h ^= h >> 16
	}

	h *= prime
	h ^= h >> 10
	h *= prime
	h ^= h >> 17

	return
}

func leu32[T byteseq](b T) uint32 {
	_ = b[3]
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
