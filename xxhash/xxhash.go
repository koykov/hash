package xxhash

import (
	"math/bits"
)

type byteseq interface {
	~[]byte | ~string
}

const (
	prime0 uint64 = 11400714785074694791
	prime1 uint64 = 14029467366897019727
	prime2 uint64 = 1609587929392839161
	prime3 uint64 = 9650029242287828579
	prime4 uint64 = 2870177450012600261
)

var primes = [5]uint64{prime0, prime1, prime2, prime3, prime4}

func Hash64[T byteseq](p T) (h uint64) {
	ul := uint64(len(p))
	if ul >= 32 {
		v0, v1, v2, v3 := primes[0]+prime1, prime1, uint64(0), -primes[0]
		for len(p) >= 32 {
			v0, v1, v2, v3, p = rnd(v0, leu64(p[0:8])), rnd(v1, leu64(p[8:16])), rnd(v2, leu64(p[16:24])), rnd(v3, leu64(p[24:32])), p[32:]
		}
		h = rotl(v0, 1) + rotl(v1, 7) + rotl(v2, 12) + rotl(v3, 18)
		h = mrnd(h, v0)
		h = mrnd(h, v1)
		h = mrnd(h, v2)
		h = mrnd(h, v3)
	} else {
		h = prime4
	}
	h += ul
	for len(p) >= 8 {
		k := rnd(0, leu64(p[:8]))
		h ^= k
		h, p = rotl(h, 27)*prime0+prime3, p[8:]
	}
	if len(p) >= 4 {
		h ^= uint64(leu32(p[:4])) * prime0
		h, p = rotl(h, 23)*prime1+prime2, p[4:]
	}
	for i := 0; i < len(p); i++ {
		h ^= uint64(p[i]) * prime4
		h = rotl(h, 11) * prime0
	}

	h ^= h >> 33
	h *= prime1
	h ^= h >> 29
	h *= prime2
	h ^= h >> 32

	return
}

func rnd(a, in uint64) uint64 {
	a += in * prime1
	a = rotl(a, 31)
	a *= prime0
	return a
}

func mrnd(a, v uint64) uint64 {
	v = rnd(0, v)
	a ^= v
	a = a*prime0 + prime3
	return a
}

func rotl(v uint64, r int) uint64 {
	return bits.RotateLeft64(v, r)
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
