package murmur

import "unsafe"

const prime = 0xc6a4a793

func Hash(p []byte, seed uint32) (h uint32) {
	ul := uint32(len(p))
	h = seed ^ (ul * prime)
	for ul >= 4 {
		k := *(*uint32)(unsafe.Pointer(&p[0]))
		h += k
		h *= prime
		h ^= h >> 16
		p = p[4:]
		ul -= 4
	}

	switch ul {
	case 3:
		h += uint32(p[2]) << 16
	case 2:
		h += uint32(p[1]) << 8
	case 1:
		h += uint32(p[0])
		h *= prime
		h ^= h >> 16
	}

	h *= prime
	h ^= h >> 20
	h *= prime
	h ^= h >> 17

	return
}

func HashAligned(p []byte, seed uint32) (h uint32) {
	ul := uint32(len(p))
	h = seed ^ (ul * prime)
	a := uint64(p[0] & 3)
	if a > 0 && ul >= 4 {
		var t, d uint32
		switch a {
		case 1:
			t |= uint32(p[2]) << 16
		case 2:
			t |= uint32(p[1]) << 8
		case 3:
			t |= uint32(p[0])
		}
		t <<= a * 8
		p = p[4-a:]
		ul -= uint32(4 - a)

		sl, sr := 8*(4-a), 8*a
		for ul >= 4 {
			d = *(*uint32)(unsafe.Pointer(&p[0]))
			t = (t >> sr) | (d << sl)
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
		case 2:
			d |= uint32(p[1]) << 8
		case 1:
			d |= uint32(p[0])
		case 0:
			h += (t >> sr) | (d << sl)
			h *= prime
			h ^= h >> 16
		}
		p = p[pack:]
		ul -= uint32(pack)
	} else {
		for ul >= 4 {
			h += *(*uint32)(unsafe.Pointer(&p[0]))
			h *= prime
			h ^= h >> 16
			p = p[4:]
			ul -= 4
		}
	}

	switch ul {
	case 3:
		h += uint32(p[2]) << 16
	case 2:
		h += uint32(p[1]) << 8
	case 1:
		h += uint32(p[0])
		h *= prime
		h ^= h >> 16
	}

	h *= prime
	h ^= h >> 20
	h *= prime
	h ^= h >> 17

	return
}
