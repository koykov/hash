package hash

type wrapfn32[T byteseq] struct {
	fn func(data T) uint32
}

func (w wrapfn32[T]) Sum32(data T) uint32 {
	return w.fn(data)
}

// WrapFunc32 makes Hasher32 over given fn.
func WrapFunc32[T byteseq](fn func(data T) uint32) Hasher32[T] {
	return wrapfn32[T]{fn: fn}
}

// ---

type wrapfn64[T byteseq] struct {
	fn func(data T) uint64
}

func (w wrapfn64[T]) Sum64(data T) uint64 {
	return w.fn(data)
}

// WrapFunc64 makes Hasher64 over given fn.
func WrapFunc64[T byteseq](fn func(data T) uint64) Hasher64[T] {
	return wrapfn64[T]{fn: fn}
}

// ---

type wrapfn128[T byteseq] struct {
	fn func(data T) [2]uint64
}

func (w wrapfn128[T]) Sum128(data T) [2]uint64 {
	return w.fn(data)
}

// WrapFunc128 makes Hasher128 over given fn.
func WrapFunc128[T byteseq](fn func(data T) [2]uint64) Hasher128x64[T] {
	return wrapfn128[T]{fn: fn}
}

// ---

type wrapfn256[T byteseq] struct {
	fn func(data T) [4]uint64
}

func (w wrapfn256[T]) Sum256(data T) [4]uint64 {
	return w.fn(data)
}

// WrapFunc256 makes Hasher256 over given fn.
func WrapFunc256[T byteseq](fn func(data T) [4]uint64) Hasher256x64[T] {
	return wrapfn256[T]{fn: fn}
}
