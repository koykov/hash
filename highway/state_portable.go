package highway

import "encoding/binary"

type lanes [4]uint64

var (
	init0 = lanes{0xdbe6d5d5fe4cce2f, 0xa4093822299f31d0, 0x13198a2e03707344, 0x243f6a8885a308d3}
	init1 = lanes{0x3bd39e10cb0ef593, 0xc0acf169b5f18a8c, 0xbe5466cf34e90c6c, 0x452821e638d01377}
)

func initPortable(s *[16]uint64, key Key) {
	copy(s[8:], init0[:])
	copy(s[12:], init1[:])
	s[0], s[1], s[2], s[3] = init0[0]^key[0], init0[1]^key[1], init0[2]^key[2], init0[3]^key[3]
	key[0], key[1], key[2], key[3] = key[0]>>32|key[0]<<32, key[1]>>32|key[1]<<32, key[2]>>32|key[2]<<32, key[3]>>32|key[3]<<32
	s[4], s[5], s[6], s[7] = init1[0]^key[0], init1[1]^key[1], init1[2]^key[2], init1[3]^key[3]
}

func updatePortable(s *[16]uint64, data []byte) {
	for len(data) >= 32 {
		m := data[:32]

		s[4] += binary.LittleEndian.Uint64(m) + s[8]
		s[8] ^= uint64(uint32(s[4])) * (s[0] >> 32)
		s[0] += s[12]
		s[12] ^= uint64(uint32(s[0])) * (s[4] >> 32)

		s[5] += binary.LittleEndian.Uint64(m[8:]) + s[9]
		s[9] ^= uint64(uint32(s[5])) * (s[1] >> 32)
		s[1] += s[13]
		s[13] ^= uint64(uint32(s[1])) * (s[5] >> 32)

		s[6] += binary.LittleEndian.Uint64(m[16:]) + s[10]
		s[10] ^= uint64(uint32(s[6])) * (s[2] >> 32)
		s[2] += s[14]
		s[14] ^= uint64(uint32(s[2])) * (s[6] >> 32)

		s[7] += binary.LittleEndian.Uint64(m[24:]) + s[11]
		s[11] ^= uint64(uint32(s[7])) * (s[3] >> 32)
		s[3] += s[15]
		s[15] ^= uint64(uint32(s[3])) * (s[7] >> 32)

		val0 := s[4]
		val1 := s[5]
		res := val0 & (0xff << (2 * 8))
		res2 := (val0 & (0xff << (7 * 8))) + (val1 & (0xff << (2 * 8)))
		res += (val1 & (0xff << (7 * 8))) >> 8
		res2 += (val0 & (0xff << (6 * 8))) >> 8
		res += ((val0 & (0xff << (5 * 8))) + (val1 & (0xff << (6 * 8)))) >> 16
		res2 += (val1 & (0xff << (5 * 8))) >> 16
		res += ((val0 & (0xff << (3 * 8))) + (val1 & (0xff << (4 * 8)))) >> 24
		res2 += ((val1 & (0xff << (3 * 8))) + (val0 & (0xff << (4 * 8)))) >> 24
		res += (val0 & (0xff << (1 * 8))) << 32
		res2 += (val1 & 0xff) << 48
		res += val0 << 56
		res2 += (val1 & (0xff << (1 * 8))) << 24

		s[0] += res
		s[1] += res2

		val0 = s[6]
		val1 = s[7]
		res = val0 & (0xff << (2 * 8))
		res2 = (val0 & (0xff << (7 * 8))) + (val1 & (0xff << (2 * 8)))
		res += (val1 & (0xff << (7 * 8))) >> 8
		res2 += (val0 & (0xff << (6 * 8))) >> 8
		res += ((val0 & (0xff << (5 * 8))) + (val1 & (0xff << (6 * 8)))) >> 16
		res2 += (val1 & (0xff << (5 * 8))) >> 16
		res += ((val0 & (0xff << (3 * 8))) + (val1 & (0xff << (4 * 8)))) >> 24
		res2 += ((val1 & (0xff << (3 * 8))) + (val0 & (0xff << (4 * 8)))) >> 24
		res += (val0 & (0xff << (1 * 8))) << 32
		res2 += (val1 & 0xff) << 48
		res += val0 << 56
		res2 += (val1 & (0xff << (1 * 8))) << 24

		s[2] += res
		s[3] += res2

		val0 = s[0]
		val1 = s[1]
		res = val0 & (0xff << (2 * 8))
		res2 = (val0 & (0xff << (7 * 8))) + (val1 & (0xff << (2 * 8)))
		res += (val1 & (0xff << (7 * 8))) >> 8
		res2 += (val0 & (0xff << (6 * 8))) >> 8
		res += ((val0 & (0xff << (5 * 8))) + (val1 & (0xff << (6 * 8)))) >> 16
		res2 += (val1 & (0xff << (5 * 8))) >> 16
		res += ((val0 & (0xff << (3 * 8))) + (val1 & (0xff << (4 * 8)))) >> 24
		res2 += ((val1 & (0xff << (3 * 8))) + (val0 & (0xff << (4 * 8)))) >> 24
		res += (val0 & (0xff << (1 * 8))) << 32
		res2 += (val1 & 0xff) << 48
		res += val0 << 56
		res2 += (val1 & (0xff << (1 * 8))) << 24

		s[4] += res
		s[5] += res2

		val0 = s[2]
		val1 = s[3]
		res = val0 & (0xff << (2 * 8))
		res2 = (val0 & (0xff << (7 * 8))) + (val1 & (0xff << (2 * 8)))
		res += (val1 & (0xff << (7 * 8))) >> 8
		res2 += (val0 & (0xff << (6 * 8))) >> 8
		res += ((val0 & (0xff << (5 * 8))) + (val1 & (0xff << (6 * 8)))) >> 16
		res2 += (val1 & (0xff << (5 * 8))) >> 16
		res += ((val0 & (0xff << (3 * 8))) + (val1 & (0xff << (4 * 8)))) >> 24
		res2 += ((val1 & (0xff << (3 * 8))) + (val0 & (0xff << (4 * 8)))) >> 24
		res += (val0 & (0xff << (1 * 8))) << 32
		res2 += (val1 & 0xff) << 48
		res += val0 << 56
		res2 += (val1 & (0xff << (1 * 8))) << 24

		s[6] += res
		s[7] += res2
		data = data[32:]
	}
}

func updateRemainderPortable(s *[16]uint64, buf *[32]byte, off int) {
	var block [blocksz]byte
	mod32 := (uint64(off) << 32) + uint64(off)
	for i := range s[:4] {
		s[i] += mod32
	}
	for i := range s[4:8] {
		t0 := uint32(s[i+4])
		t0 = (t0 << uint(off)) | (t0 >> uint(32-off))

		t1 := uint32(s[i+4] >> 32)
		t1 = (t1 << uint(off)) | (t1 >> uint(32-off))

		s[i+4] = (uint64(t1) << 32) | uint64(t0)
	}

	mod4 := off & 3
	remain := off - mod4

	copy(block[:], buf[:remain])
	if off >= 16 {
		copy(block[28:], buf[off-4:])
	} else if mod4 != 0 {
		last := uint32(buf[remain])
		last += uint32(buf[remain+mod4>>1]) << 8
		last += uint32(buf[off-1]) << 16
		binary.LittleEndian.PutUint32(block[16:], last)
	}
	updatePortable(s, block[:])
}

func finalize8Portable(s *[16]uint64) uint64 {
	var (
		perm [4]uint64
		buf  [32]byte
	)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)

	return s[0] + s[4] + s[8] + s[12]
}

func finalize16Portable(s *[16]uint64) (h [2]uint64) {
	var (
		perm [4]uint64
		buf  [32]byte
	)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)

	h[0], h[1] = s[0]+s[6]+s[8]+s[14], s[1]+s[7]+s[9]+s[15]
	return
}

func finalize32Portable(s *[16]uint64) (h [4]uint64) {
	var (
		perm [4]uint64
		buf  [32]byte
	)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)
	finalizeRunPortable(s, &perm, &buf)

	h[0], h[1] = reduce(s[0]+s[8], s[1]+s[9], s[4]+s[12], s[5]+s[13])
	h[2], h[3] = reduce(s[2]+s[10], s[3]+s[11], s[6]+s[14], s[7]+s[15])
	return
}

func finalizeRunPortable(s *[16]uint64, perm *[4]uint64, buf *[32]byte) {
	perm[0], perm[1], perm[2], perm[3] = s[2]>>32|s[2]<<32, s[3]>>32|s[3]<<32, s[0]>>32|s[0]<<32, s[1]>>32|s[1]<<32
	binary.LittleEndian.PutUint64(buf[0:], perm[0])
	binary.LittleEndian.PutUint64(buf[8:], perm[1])
	binary.LittleEndian.PutUint64(buf[16:], perm[2])
	binary.LittleEndian.PutUint64(buf[24:], perm[3])
	updatePortable(s, buf[:])
}

func reduce(v0, v1, v2, v3 uint64) (r0, r1 uint64) {
	v3 &= 0x3FFFFFFFFFFFFFFF

	r0, r1 = v2, v3

	v3 = (v3 << 1) | (v2 >> (64 - 1))
	v2 <<= 1
	r1 = (r1 << 2) | (r0 >> (64 - 2))
	r0 <<= 2

	r0 ^= v0 ^ v2
	r1 ^= v1 ^ v3
	return
}
