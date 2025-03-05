package murmur

import "testing"

var stages = []struct {
	data       string
	seed       uint32
	seed64     uint64
	hash       uint32
	hash128x86 [4]uint32
	hash128x64 [2]uint64
}{
	{
		data:       "foo",
		seed:       0x12345678,
		seed64:     0x12345678,
		hash:       4084948345,
		hash128x86: [4]uint32{1918736524, 2601350885, 2601350885, 2601350885},
		hash128x64: [2]uint64{15043494824502257628, 1408089765573727336},
	},
	{
		data:       "foobar",
		seed:       0x12345678,
		seed64:     0x12345678,
		hash:       3074407335,
		hash128x86: [4]uint32{2963197276, 938027862, 3579455989, 3579455989},
		hash128x64: [2]uint64{12751283261750357652, 7917742029045779757},
	},
	{
		data:       "The quick brown fox jumps over the lazy dog",
		seed:       0x9747b28c,
		seed64:     0x9747b28c,
		hash:       3647120414,
		hash128x86: [4]uint32{2329204062, 1287151985, 2393324188, 3451287870},
		hash128x64: [2]uint64{8325606756057297185, 17961889624427075301},
	},
	{
		data:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget risus vitae est sagittis euismod. Integer id nibh ut ligula aliquam sagittis.",
		seed:       0x12345678,
		seed64:     0x12345678,
		hash:       991832159,
		hash128x86: [4]uint32{368352206, 1111628908, 560953716, 1558498190},
		hash128x64: [2]uint64{1320662463624528565, 6505946595269923697},
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

func TestHash_x86_128(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := Hash128x86(stage.data, stage.seed)
			if hash != stage.hash128x86 {
				t.Errorf("Hash128x86(%s, %d) = %v, want %v", stage.data, stage.seed, hash, stage.hash128x86)
			}
		})
	}
}

func TestHash_x64_128(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := Hash128x64(stage.data, stage.seed64)
			if hash != stage.hash128x64 {
				t.Errorf("Hash128x64(%s, %d) = %v, want %v", stage.data, stage.seed, hash, stage.hash128x64)
			}
		})
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(stage.data)))
			for j := 0; j < b.N; j++ {
				Hash(stage.data, stage.seed)
			}
		})
	}
}

func BenchmarkHash128x86(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(stage.data)))
			for j := 0; j < b.N; j++ {
				Hash128x86(stage.data, stage.seed)
			}
		})
	}
}

func BenchmarkHash128x64(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(stage.data)))
			for j := 0; j < b.N; j++ {
				Hash128x64(stage.data, stage.seed64)
			}
		})
	}
}
