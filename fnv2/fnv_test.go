package fnv

import "testing"

var stages = []struct {
	data    string
	hash32  uint32
	hash32a uint32
	hash64  uint64
	hash64a uint64
}{
	{
		data:    "foo",
		hash32:  1083137555,
		hash32a: 2851307223,
		hash64:  15621798640163566899,
		hash64a: 15902901984413996407,
	},
	{
		data:    "foobar",
		hash32:  837857890,
		hash32a: 3214735720,
		hash64:  3750802935296928194,
		hash64a: 9625390261332436968,
	},
	{
		data:    "The quick brown fox jumps over the lazy dog",
		hash32:  3922226286,
		hash32a: 76545936,
		hash64:  12156045600678443726,
		hash64a: 17580284887202820368,
	},
	{
		data:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget risus vitae est sagittis euismod. Integer id nibh ut ligula aliquam sagittis.",
		hash32:  3292857305,
		hash32a: 1188270929,
		hash64:  15615971020568113913,
		hash64a: 8713077696443677873,
	},
}

func TestHash32(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			result := Hash32(stage.data)
			if result != stage.hash32 {
				t.Errorf("Hash32(%s) = %d, want %d", stage.data, result, stage.hash32)
			}
		})
	}
}

func TestHash32a(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			result := Hash32a(stage.data)
			if result != stage.hash32a {
				t.Errorf("Hash32a(%s) = %d, want %d", stage.data, result, stage.hash32a)
			}
		})
	}
}

func TestHash64(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			result := Hash64(stage.data)
			if result != stage.hash64 {
				t.Errorf("Hash64(%s) = %d, want %d", stage.data, result, stage.hash64)
			}
		})
	}
}

func TestHash64a(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			result := Hash64a(stage.data)
			if result != stage.hash64a {
				t.Errorf("Hash64a(%s) = %d, want %d", stage.data, result, stage.hash64a)
			}
		})
	}
}

func BenchmarkHash32(b *testing.B) {
	b.ReportAllocs()
	var h Hasher32[string]
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		h.Sum32(stage.data)
	}
}

func BenchmarkHash32a(b *testing.B) {
	b.ReportAllocs()
	var h Hasher32a[string]
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		h.Sum32(stage.data)
	}
}

func BenchmarkHash64(b *testing.B) {
	b.ReportAllocs()
	var h Hasher64[string]
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		h.Sum64(stage.data)
	}
}

func BenchmarkHash64a(b *testing.B) {
	b.ReportAllocs()
	var h Hasher64a[string]
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		h.Sum64(stage.data)
	}
}
