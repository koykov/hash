package v3

import "testing"

var (
	stages = []struct {
		data         string
		seed         uint32
		seed64       uint64
		hash         uint32
		hash_x86_128 [4]uint32
		hash_x64_128 [2]uint64
	}{
		{
			data:         "foo",
			seed:         0x12345678,
			seed64:       0x12345678,
			hash:         4084948345,
			hash_x86_128: [4]uint32{1918736524, 2601350885, 2601350885, 2601350885},
			hash_x64_128: [2]uint64{15043494824502257628, 1408089765573727336},
		},
		{
			data:         "foobar",
			seed:         0x12345678,
			seed64:       0x12345678,
			hash:         3074407335,
			hash_x86_128: [4]uint32{2963197276, 938027862, 3579455989, 3579455989},
			hash_x64_128: [2]uint64{12751283261750357652, 7917742029045779757},
		},
		{
			data:         "The quick brown fox jumps over the lazy dog",
			seed:         0x9747b28c,
			seed64:       0x9747b28c,
			hash:         3647120414,
			hash_x86_128: [4]uint32{2329204062, 1287151985, 2393324188, 3451287870},
			hash_x64_128: [2]uint64{8325606756057297185, 17961889624427075301},
		},
		{
			data:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget risus vitae est sagittis euismod. Integer id nibh ut ligula aliquam sagittis.",
			seed:         0x12345678,
			seed64:       0x12345678,
			hash:         991832159,
			hash_x86_128: [4]uint32{368352206, 1111628908, 560953716, 1558498190},
			hash_x64_128: [2]uint64{1320662463624528565, 6505946595269923697},
		},
	}
	total int64
)

func init() {
	for i := 0; i < len(stages); i++ {
		total += int64(len(stages[i].data))
	}
}

func equal_x86_128(a, b [4]uint32) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3]
}

func equal_x64_128(a, b [2]uint64) bool {
	return a[0] == b[0] && a[1] == b[1]
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
			hash := Hash_x86_128(stage.data, stage.seed)
			if !equal_x86_128(hash, stage.hash_x86_128) {
				t.Errorf("Hash_x86_128(%s, %d) = %v, want %v", stage.data, stage.seed, hash, stage.hash_x86_128)
			}
		})
	}
}

func TestHash_x64_128(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			hash := Hash_x64_128(stage.data, stage.seed64)
			if !equal_x64_128(hash, stage.hash_x64_128) {
				t.Errorf("Hash_x64_128(%s, %d) = %v, want %v", stage.data, stage.seed, hash, stage.hash_x64_128)
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

func BenchmarkHash_x86_128(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(total)
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		Hash_x86_128(stage.data, stage.seed)
	}
}

func BenchmarkHash_x64_128(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(total)
	for i := 0; i < b.N; i++ {
		stage := &stages[b.N%len(stages)]
		Hash_x64_128(stage.data, stage.seed64)
	}
}
