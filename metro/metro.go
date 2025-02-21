package metro

import (
	"math/bits"
)

type byteseq interface {
	~[]byte | ~string
}

func Hash64[T byteseq](p T, seed uint64) (h uint64) {
	const (
		k0 = 0xD6D018F5
		k1 = 0xA2AA033B
		k2 = 0x62992FC1
		k3 = 0x30BC5B29
	)

	h = (seed + k2) * k0
	if len(p) >= 32 {
		v0, v1, v2, v3 := h, h, h, h

		for len(p) >= 32 {
			v0 += leu64(p[:8]) * k0
			v0 = bits.RotateLeft64(v0, -29) + v2
			v1 += leu64(p[8:16]) * k1
			v1 = bits.RotateLeft64(v1, -29) + v3
			v2 += leu64(p[16:24]) * k2
			v2 = bits.RotateLeft64(v2, -29) + v0
			v3 += leu64(p[24:32]) * k3
			v3 = bits.RotateLeft64(v3, -29) + v1
			p = p[32:]
		}

		v2 ^= bits.RotateLeft64(((v0+v3)*k0)+v1, -37) * k1
		v3 ^= bits.RotateLeft64(((v1+v2)*k1)+v0, -37) * k0
		v0 ^= bits.RotateLeft64(((v0+v2)*k0)+v3, -37) * k1
		v1 ^= bits.RotateLeft64(((v1+v3)*k1)+v2, -37) * k0
		h += v0 ^ v1
	}
	if len(p) >= 16 {
		v0 := h + (leu64(p[:8]) * k2)
		v0 = bits.RotateLeft64(v0, -29) * k3
		v1 := h + (leu64(p[8:16]) * k2)
		v1 = bits.RotateLeft64(v1, -29) * k3
		v0 ^= bits.RotateLeft64(v0*k0, -21) + v1
		v1 ^= bits.RotateLeft64(v1*k3, -21) + v0
		h += v1
		p = p[16:]
	}
	if len(p) >= 8 {
		h += leu64(p[:8]) * k3
		p = p[8:]
		h ^= bits.RotateLeft64(h, -55) * k1
	}
	if len(p) >= 4 {
		h += uint64(leu32(p[:4])) * k3
		h ^= bits.RotateLeft64(h, -26) * k1
		p = p[4:]
	}
	if len(p) >= 2 {
		h += uint64(leu16(p[:2])) * k3
		p = p[2:]
		h ^= bits.RotateLeft64(h, -48) * k1
	}
	if len(p) >= 1 {
		h += uint64(p[0]) * k3
		h ^= bits.RotateLeft64(h, -37) * k1
	}

	h ^= bits.RotateLeft64(h, -28)
	h *= k0
	h ^= bits.RotateLeft64(h, -29)

	return
}

func Hash128[T byteseq](p T, seed uint64) (h [2]uint64) {
	const (
		k0 = 0xC83A91E1
		k1 = 0x8648DBDB
		k2 = 0x7BDEC03B
		k3 = 0x2F5870A5
	)

	h[0], h[1] = (seed-k0)*k3, (seed+k1)*k2
	var v [2]uint64
	if len(p) >= 32 {
		v[0], v[1] = (seed+k0)*k2, (seed-k1)*k3

		for len(p) >= 32 {
			h[0] += leu64(p) * k0
			p = p[8:]
			h[0] = rotr(h[0], 29) + v[0]
			h[1] += leu64(p) * k1
			p = p[8:]
			h[1] = rotr(h[1], 29) + v[1]
			v[0] += leu64(p) * k2
			p = p[8:]
			v[0] = rotr(v[0], 29) + h[0]
			v[1] += leu64(p) * k3
			p = p[8:]
			v[1] = rotr(v[1], 29) + h[1]
		}

		v[0] ^= rotr(((h[0]+v[1])*k0)+h[1], 21) * k1
		v[1] ^= rotr(((h[1]+v[0])*k1)+h[0], 21) * k0
		h[0] ^= rotr(((h[0]+v[0])*k0)+v[1], 21) * k1
		h[1] ^= rotr(((h[1]+v[1])*k1)+v[0], 21) * k0
	}
	if len(p) >= 16 {
		h[0] += leu64(p) * k2
		p = p[8:]
		h[0] = rotr(h[0], 33) * k3
		h[1] += leu64(p) * k2
		p = p[8:]
		h[1] = rotr(h[1], 33) * k3
		h[0] ^= rotr((h[0]*k2)+h[1], 45) * k1
		h[1] ^= rotr((h[1]*k3)+h[0], 45) * k0
	}
	if len(p) >= 8 {
		h[0] += leu64(p) * k2
		p = p[8:]
		h[0] = rotr(h[0], 33) * k3
		h[0] ^= rotr((h[0]*k2)+h[1], 27) * k1
	}
	if len(p) >= 4 {
		h[1] += uint64(leu32(p)) * k2
		p = p[4:]
		h[1] = rotr(h[1], 33) * k3
		h[1] ^= rotr((h[1]*k3)+h[0], 46) * k0
	}
	if len(p) >= 2 {
		h[0] += uint64(leu16(p)) * k2
		p = p[2:]
		h[0] = rotr(h[0], 33) * k3
		h[0] ^= rotr((h[0]*k2)+h[1], 22) * k1
	}
	if len(p) >= 1 {
		h[1] += uint64(p[0]) * k2
		h[1] = rotr(h[1], 33) * k3
		h[1] ^= rotr((h[1]*k3)+h[0], 58) * k0
	}
	h[0] += rotr((h[0]*k0)+h[1], 13)
	h[1] += rotr((h[1]*k1)+h[0], 37)
	h[0] += rotr((h[0]*k2)+h[1], 13)
	h[1] += rotr((h[1]*k3)+h[0], 37)

	return
}

func rotr(v uint64, k uint) uint64 {
	return (v >> k) | (v << (64 - k))
}

func leu16[T byteseq](b T) uint64 {
	_ = b[1]
	return uint64(b[0]) | uint64(b[1])<<8
}

func leu32[T byteseq](b T) uint64 {
	_ = b[3]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24
}

func leu64[T byteseq](b T) uint64 {
	_ = b[7]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}
