package murmur

import "testing"

var stages = []struct {
	data    string
	seed    uint32
	seed64  uint64
	hash    uint32
	hash64A uint64
	hash64B uint64
	hash2A  uint32
	hashN2  uint32
	hashA   uint32
}{
	{
		data:    "foo",
		seed:    0x12345678,
		hash:    1532637697,
		hash64A: 14834356025302342401,
		hash64B: 5484325262697493828,
		hash2A:  2221687037,
		hashN2:  1532637697,
		hashA:   1532637697,
	},
	{
		data:    "foobar",
		seed:    0x12345678,
		hash:    151733797,
		hash64A: 15321041522486911382,
		hash64B: 4480287663481255131,
		hash2A:  362873721,
		hashN2:  151733797,
		hashA:   151733797,
	},
	{
		data:    "The quick brown fox jumps over the lazy dog",
		seed:    0x9747b28c,
		hash:    495243318,
		hash64A: 6163679885495272987,
		hash64B: 8470663747738974033,
		hash2A:  3814933764,
		hashN2:  495243318,
		hashA:   495243318,
	},
	{
		data:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget risus vitae est sagittis euismod. Integer id nibh ut ligula aliquam sagittis.",
		seed:    0x12345678,
		hash:    2279665421,
		hash64A: 928382629942818458,
		hash64B: 1345853299853536290,
		hash2A:  507073689,
		hashN2:  2279665421,
		hashA:   2279665421,
	},
}

func TestHash(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := Hash(stage.data, stage.seed)
			if hash != stage.hash {
				t.Errorf("Hash(%s, %d) = %d, want %d", stage.data, stage.seed, hash, stage.hash)
			}
		})
	}
}

func TestHash64A(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := Hash64A(stage.data, stage.seed64)
			if hash != stage.hash64A {
				t.Errorf("Hash64A(%s, %d) = %d, want %d", stage.data, stage.seed, hash, stage.hash64A)
			}
		})
	}
}

func TestHash64B(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := Hash64B(stage.data, stage.seed64)
			if hash != stage.hash64B {
				t.Errorf("Hash64B(%s, %d) = %d, want %d", stage.data, stage.seed, hash, stage.hash64B)
			}
		})
	}
}

func TestHash2A(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := Hash2A(stage.data, stage.seed)
			if hash != stage.hash2A {
				t.Errorf("Hash2A(%s, %d) = %d, want %d", stage.data, stage.seed, hash, stage.hash2A)
			}
		})
	}
}

func TestHashNeutral2(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := HashNeutral2(stage.data, stage.seed)
			if hash != stage.hashN2 {
				t.Errorf("HashNeutral2(%s, %d) = %d, want %d", stage.data, stage.seed, hash, stage.hashN2)
			}
		})
	}
}

func TestHashAligned(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := HashAligned(stage.data, stage.seed)
			if hash != stage.hashA {
				t.Errorf("HashAligned(%s, %d) = %d, want %d", stage.data, stage.seed, hash, stage.hashA)
			}
		})
	}
}

func BenchmarkHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		Hash(stage.data, stage.seed)
	}
}

func BenchmarkHash64A(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		Hash64A(stage.data, stage.seed64)
	}
}

func BenchmarkHash64B(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		Hash64B(stage.data, stage.seed64)
	}
}

func BenchmarkHash2A(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		Hash2A(stage.data, stage.seed)
	}
}

func BenchmarkHashNeutral2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		HashNeutral2(stage.data, stage.seed)
	}
}

func BenchmarkHashAligned(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		HashAligned(stage.data, stage.seed)
	}
}
