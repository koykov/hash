package metro

import "testing"

var stages = []struct {
	data    string
	seed    uint64
	hash64  uint64
	hash128 [2]uint64
}{
	{
		data:    "foo",
		seed:    0x12345678,
		hash64:  2380772528196146020,
		hash128: [2]uint64{17949631740792705275, 8991474027653484836},
	},
	{
		data:    "foobar",
		seed:    0x12345678,
		hash64:  5426739015579503264,
		hash128: [2]uint64{3972895474774536955, 18246469663299570930},
	},
	{
		data:    "The quick brown fox jumps over the lazy dog",
		seed:    0x12345678,
		hash64:  3210435738747958464,
		hash128: [2]uint64{10519808379090762444, 395168348050836046},
	},
	{
		data:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget risus vitae est sagittis euismod. Integer id nibh ut ligula aliquam sagittis.",
		seed:    0x12345678,
		hash64:  13537174898478009123,
		hash128: [2]uint64{6001879636519351909, 15368227032245481045},
	},
}

func equal128(a, b [2]uint64) bool {
	return a[0] == b[0] && a[1] == b[1]
}

func TestHash64(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			result := Hash64(stage.data, stage.seed)
			if result != stage.hash64 {
				t.Errorf("Hash64(%s) = %d, want %d", stage.data, result, stage.hash64)
			}
		})
	}
}

func TestHash128(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run("", func(t *testing.T) {
			result := Hash128(stage.data, stage.seed)
			if !equal128(result, stage.hash128) {
				t.Errorf("Hash128(%s) = %v, want %v", stage.data, result, stage.hash128)
			}
		})
	}
}

func BenchmarkHash64(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(stage.data)))
			for j := 0; j < b.N; j++ {
				Hash64(stage.data, stage.seed)
			}
		})
	}
}

func BenchmarkHash128(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(stage.data)))
			for j := 0; j < b.N; j++ {
				Hash128(stage.data, stage.seed)
			}
		})
	}
}
