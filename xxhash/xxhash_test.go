package xxhash

import "testing"

var stages = []struct {
	data   string
	hash64 uint64
}{
	{
		data:   "foo",
		hash64: 3728699739546630719,
	},
	{
		data:   "foobar",
		hash64: 11721187498075204345,
	},
	{
		data:   "The quick brown fox jumps over the lazy dog",
		hash64: 802816344064684476,
	},
	{
		data:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget risus vitae est sagittis euismod. Integer id nibh ut ligula aliquam sagittis.",
		hash64: 7536621141780310046,
	},
}

func TestHash64(t *testing.T) {
	for _, stage := range stages {
		hash := Hash64(stage.data)
		if hash != stage.hash64 {
			t.Errorf("Hash64(%s) = %d, want %d", stage.data, hash, stage.hash64)
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
				Hash64(stage.data)
			}
		})
	}
}
