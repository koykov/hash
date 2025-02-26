package highway

import (
	"testing"
)

var stages = []struct {
	data    string
	key     Key
	hash64  uint64
	hash128 [2]uint64
	hash256 [4]uint64
}{
	{
		data:    "foo",
		key:     Key{1, 2, 3, 4},
		hash64:  9859378340892451915,
		hash128: [2]uint64{11706677712814371206, 4410658883535296641},
		hash256: [4]uint64{11976519617788884014, 7931315462684698195, 16103404768441509400, 7280876034644783291},
	},
	{
		data:    "foobar",
		key:     Key{1, 2, 3, 4},
		hash64:  9918315496747124426,
		hash128: [2]uint64{11075638380637778854, 4473070949374928523},
		hash256: [4]uint64{5824518272286942966, 16861439246469016720, 13762183892873225750, 17763125831714132369},
	},
	{
		data:    "The quick brown fox jumps over the lazy dog",
		key:     Key{1, 2, 3, 4},
		hash64:  2834632039609883303,
		hash128: [2]uint64{1392881838727189294, 342905740284669911},
		hash256: [4]uint64{15237698151047879401, 3671790918387159650, 8435811979156640686, 15111839126228456129},
	},
	{
		data:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget risus vitae est sagittis euismod. Integer id nibh ut ligula aliquam sagittis.",
		key:     Key{1, 2, 3, 4},
		hash64:  8213030592213311714,
		hash128: [2]uint64{15455821470471741311, 4834011644375232670},
		hash256: [4]uint64{10122432682956746660, 17119628479380443130, 17398149750842642152, 1551113674346516528},
	},
}

func TestHash64(t *testing.T) {
	for _, stage := range stages {
		hash := Hash64([]byte(stage.data), stage.key)
		if hash != stage.hash64 {
			t.Errorf("Hash64(%s) = %d, want %d", stage.data, hash, stage.hash64)
		}
	}
}

func TestHash128(t *testing.T) {
	for _, stage := range stages {
		hash := Hash128([]byte(stage.data), stage.key)
		if hash != stage.hash128 {
			t.Errorf("Hash128(%s) = %v, want %v", stage.data, hash, stage.hash128)
		}
	}
}

func TestHash256(t *testing.T) {
	for _, stage := range stages {
		hash := Hash256([]byte(stage.data), stage.key)
		if hash != stage.hash256 {
			t.Errorf("Hash256(%s) = %v, want %v", stage.data, hash, stage.hash256)
		}
	}
}

func BenchmarkHash64(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(stage.data)))
			for j := 0; j < b.N; j++ {
				Hash64(stage.data, stage.key)
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
				Hash128(stage.data, stage.key)
			}
		})
	}
}

func BenchmarkHash256(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(stage.data)))
			for j := 0; j < b.N; j++ {
				Hash256(stage.data, stage.key)
			}
		})
	}
}
