package murmur

type byteseq interface {
	~[]byte | ~string
}

// Hash implement default Murmur v3 (x86/32) hash algorithm.
func Hash[T byteseq](p T, seed uint32) uint32 {
	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
		c3 = 0xe6546b64
	)

	ul := uint32(len(p))
	var nb = ul / 4

	var h1 = seed
	var k1, i uint32
	for i = 0; i < nb; i++ {
		k1 = leu32(p[i*4 : i*4+4])

		k1 *= c1
		k1 = rotl(k1, 15)
		k1 *= c2

		h1 ^= k1
		h1 = rotl(h1, 13)
		h1 = h1*5 + c3
	}

	tail := p[nb*4:]
	switch ul & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1
		k1 = rotl(k1, 15)
		k1 *= c2
		h1 ^= k1
	}

	return Fmix32(h1 ^ ul)
}

func Hash128x86[T byteseq](p T, seed uint32) (h [4]uint32) {
	const (
		c1 = 0x239b961b
		c2 = 0xab0e9789
		c3 = 0x38b34ae5
		c4 = 0xa1e38b93
	)

	ul := uint64(len(p))
	var nb = ul / 16
	h[0], h[1], h[2], h[3] = seed, seed, seed, seed
	var k1, k2, k3, k4 uint32
	for i := uint64(0); i < nb; i++ {
		k1, k2, k3, k4 = leu32(p[i*16:i*16+4]), leu32(p[i*16+4:i*16+8]), leu32(p[i*16+8:i*16+12]), leu32(p[i*16+12:i*16+16])

		k1 *= c1
		k1 = rotl(k1, 15)
		k1 *= c2
		h[0] ^= k1
		h[0] = rotl(h[0], 19)
		h[0] += h[1]
		h[0] = h[0]*5 + 0x561ccd1b
		k2 *= c2
		k2 = rotl(k2, 16)
		k2 *= c3
		h[1] ^= k2
		h[1] = rotl(h[1], 17)
		h[1] += h[2]
		h[1] = h[1]*5 + 0x0bcaa747
		k3 *= c3
		k3 = rotl(k3, 17)
		k3 *= c4
		h[2] ^= k3
		h[2] = rotl(h[2], 15)
		h[2] += h[3]
		h[2] = h[2]*5 + 0x96cd1c35
		k4 *= c4
		k4 = rotl(k4, 18)
		k4 *= c1
		h[3] ^= k4
		h[3] = rotl(h[3], 13)
		h[3] += h[0]
		h[3] = h[3]*5 + 0x32ac3b17
	}
	tail := p[nb*16:]

	k1, k2, k3, k4 = 0, 0, 0, 0
	switch ul & 15 {
	case 15:
		k4 ^= uint32(tail[14]) << 16
		fallthrough
	case 14:
		k4 ^= uint32(tail[13]) << 8
		fallthrough
	case 13:
		k4 ^= uint32(tail[12]) << 0
		k4 *= c4
		k4 = rotl(k4, 18)
		k4 *= c1
		h[3] ^= k4
		fallthrough
	case 12:
		k3 ^= uint32(tail[11]) << 24
		fallthrough
	case 11:
		k3 ^= uint32(tail[10]) << 16
		fallthrough
	case 10:
		k3 ^= uint32(tail[9]) << 8
		fallthrough
	case 9:
		k3 ^= uint32(tail[8]) << 0
		k3 *= c3
		k3 = rotl(k3, 17)
		k3 *= c4
		h[2] ^= k3
		fallthrough
	case 8:
		k2 ^= uint32(tail[7]) << 24
		fallthrough
	case 7:
		k2 ^= uint32(tail[6]) << 16
		fallthrough
	case 6:
		k2 ^= uint32(tail[5]) << 8
		fallthrough
	case 5:
		k2 ^= uint32(tail[4]) << 0
		k2 *= c2
		k2 = rotl(k2, 16)
		k2 *= c3
		h[1] ^= k2
		fallthrough
	case 4:
		k1 ^= uint32(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0]) << 0
		k1 *= c1
		k1 = rotl(k1, 15)
		k1 *= c2
		h[0] ^= k1
	}

	h[0] ^= uint32(ul)
	h[1] ^= uint32(ul)
	h[2] ^= uint32(ul)
	h[3] ^= uint32(ul)

	h[0] += h[1]
	h[0] += h[2]
	h[0] += h[3]
	h[1] += h[0]
	h[2] += h[0]
	h[3] += h[0]

	h[0], h[1], h[2], h[3] = Fmix32(h[0]), Fmix32(h[1]), Fmix32(h[2]), Fmix32(h[3])

	h[0] += h[1]
	h[0] += h[2]
	h[0] += h[3]
	h[1] += h[0]
	h[2] += h[0]
	h[3] += h[0]

	return
}

func Hash128x64[T byteseq](p T, seed uint64) (h [2]uint64) {
	const (
		c1 = 0x87c37b91114253d5
		c2 = 0x4cf5ad432745937f
	)

	ul := uint64(len(p))
	var nb = ul / 16
	h[0], h[1] = seed, seed
	var k1, k2 uint64
	for i := uint64(0); i < nb; i++ {
		k1, k2 = leu64(p[i*16:i*16+8]), leu64(p[i*16+8:i*16+16])

		k1 *= c1
		k1 = rotl64(k1, 31)
		k1 *= c2
		h[0] ^= k1
		h[0] = rotl64(h[0], 27)
		h[0] += h[1]
		h[0] = h[0]*5 + 0x52dce729
		k2 *= c2
		k2 = rotl64(k2, 33)
		k2 *= c1
		h[1] ^= k2
		h[1] = rotl64(h[1], 31)
		h[1] += h[0]
		h[1] = h[1]*5 + 0x38495ab5
	}
	tail := p[nb*16:]

	k1, k2 = 0, 0
	switch ul & 15 {
	case 15:
		k2 ^= uint64(tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(tail[8]) << 0
		k2 *= c2
		k2 = rotl64(k2, 33)
		k2 *= c1
		h[1] ^= k2
		fallthrough
	case 8:
		k1 ^= uint64(tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(tail[0]) << 0
		k1 *= c1
		k1 = rotl64(k1, 31)
		k1 *= c2
		h[0] ^= k1
	}

	h[0] ^= ul
	h[1] ^= ul

	h[0] += h[1]
	h[1] += h[0]

	h[0], h[1] = Fmix64(h[0]), Fmix64(h[1])

	h[0] += h[1]
	h[1] += h[0]

	return
}

func rotl(x uint32, r int8) (z uint32) {
	s := uint8(r) & (32 - 1)
	return (x << s) | (x >> (32 - s))
}

func rotl64(x uint64, r int8) (z uint64) {
	s := uint8(r) & (64 - 1)
	return (x << s) | (x >> (64 - s))
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
