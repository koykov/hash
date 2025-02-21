# ~~Fast FNV-1 hash~~

It's a collection of fast and alloc-free FNV-1 hash calculation.

### Caution!

> Package is abandoned, use [fnv2](/fnv2) instead.
> This package kept for backward compatibility in other packages.

FNV-1 32 benchmark:
```
BenchmarkFnv32-8                  300000000    5.79 ns/op    0 B/op    0 allocs/op
BenchmarkFnv32Native-8            100000000    21.6 ns/op    4 B/op    1 allocs/op
```

FNV-1a 32 benchmark:
```
BenchmarkFnv32a-8                 300000000    5.81 ns/op    0 B/op    0 allocs/op
BenchmarkFnv32aNative-8           100000000    22.0 ns/op    4 B/op    1 allocs/op
```

FNV-1 64 benchmark:
```
BenchmarkFnv64-8                  300000000    5.94 ns/op    0 B/op    0 allocs/op
BenchmarkFnv64Native-8            50000000     24.5 ns/op    8 B/op    1 allocs/op
```

FNV-1a 64 benchmark:
```
BenchmarkFnv64a-8                 200000000    6.51 ns/op    0 B/op    0 allocs/op
BenchmarkFnv64aNative-8           50000000     24.7 ns/op    8 B/op    1 allocs/op
```

Tech note: function with suffix "Long" (eg. Fnv32Long) uses loop-rolling to speed-up hash calculation of long strings. On short string it has
no effect or maybe a bit slowly than simple loop.

