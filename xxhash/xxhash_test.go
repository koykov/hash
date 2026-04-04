package xxhash

import (
	"math"
	"strconv"
	"testing"
)

type stage struct {
	data   string
	hash64 uint64
}

var stages = []stage{
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
var bstages []stage

func init() {
	sizes := []int{4, 16, 100, 4000, 10000000}
	for _, size := range sizes {
		b := make([]byte, size)
		for i := 0; i < size; i++ {
			b[i] = byte(i % math.MaxUint8)
		}
		bstages = append(bstages, stage{data: string(b)})
	}
}

func TestHash64(t *testing.T) {
	for _, stg := range stages {
		hash := Hash64(stg.data)
		if hash != stg.hash64 {
			t.Errorf("Hash64(%s) = %d, want %d", stg.data, hash, stg.hash64)
		}
	}
}

func BenchmarkHash64(b *testing.B) {
	for i := 0; i < len(bstages); i++ {
		stg := &bstages[i]
		b.Run(strconv.Itoa(len(stg.data)), func(b *testing.B) {
			b.SetBytes(int64(len(stg.data)))
			b.ReportAllocs()
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				Hash64(stg.data)
			}
		})
	}
}
