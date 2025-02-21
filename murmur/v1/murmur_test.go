package murmur

import "testing"

var (
	stages = []struct {
		data  string
		seed  uint32
		hash  uint32
		hashA uint32
	}{
		{
			data:  "foo",
			seed:  0x12345678,
			hash:  2077645772,
			hashA: 2077645772,
		},
		{
			data:  "foobar",
			seed:  0x12345678,
			hash:  1884093689,
			hashA: 1884093689,
		},
		{
			data:  "The quick brown fox jumps over the lazy dog",
			seed:  0x9747b28c,
			hash:  3950314481,
			hashA: 2484540071,
		},
		{
			data:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget risus vitae est sagittis euismod. Integer id nibh ut ligula aliquam sagittis.",
			seed:  0x12345678,
			hash:  2016556361,
			hashA: 291077287,
		},
	}
	total int64
)

func init() {
	for i := 0; i < len(stages); i++ {
		total += int64(len(stages[i].data))
	}
}

func TestHash(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			result := Hash(stage.data, stage.seed)
			if result != stage.hash {
				t.Errorf("Hash(%s, %d) = %d, want %d", stage.data, stage.seed, result, stage.hash)
			}
		})
	}
}

func TestHashAligned(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			result := HashAligned(stage.data, stage.seed)
			if result != stage.hashA {
				t.Errorf("HashAligned(%s, %d) = %d, want %d", stage.data, stage.seed, result, stage.hash)
			}
		})
	}
}

func BenchmarkHash(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(total)
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		Hash(stage.data, stage.seed)
	}
}

func BenchmarkHashAligned(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(total)
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		HashAligned(stage.data, stage.seed)
	}
}
