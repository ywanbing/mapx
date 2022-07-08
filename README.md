# mapx
如果你经常使用map，但是里面的内容又非常小，那么推荐你使用`mapx` ，性能和占用内存将得到提升。

## 特性
- 这个包只支持固定的场景，因为在其他场景他的性能可能不是那么好。
- 数据通过保存到map中，来作为内存缓存的情况下，并且每一个map的对象不是很多，但是map又有很多。
- 对于这个场景下，不采用map存放，而改用切片的方式。因为存放的数据只会拿来读，不会修改。
- 只针对读操作进行优化，写和删除的性能不在考虑范围内。

## 使用
和正常的使用方式一样，提供：
- Set(K,V)
- Get(K)
- GetOk(K)
- Del(K)
- Range(func(K,V))  

这些方法

## 压测结果：
```shell
goos: windows
goarch: amd64
pkg: github.com/ywanbing/mapx
cpu: Intel(R) Core(TM) i5-10600KF CPU @ 4.10GHz

Benchmark_Get8
Benchmark_Get8/Mapx
Benchmark_Get8/Mapx-12          353295718                3.379 ns/op
Benchmark_Get8/map
Benchmark_Get8/map-12           304061962                3.938 ns/op

Benchmark_Get16
Benchmark_Get16/Mapx
Benchmark_Get16/Mapx-12         287812369                4.199 ns/op
Benchmark_Get16/map
Benchmark_Get16/map-12          208549993                5.803 ns/op

Benchmark_Get20
Benchmark_Get20/Mapx
Benchmark_Get20/Mapx-12         214145983                5.357 ns/op
Benchmark_Get20/map
Benchmark_Get20/map-12          196552050                6.122 ns/op

Benchmark_GetStruct8
Benchmark_GetStruct8/Mapx
Benchmark_GetStruct8/Mapx-12    356028678                3.374 ns/op
Benchmark_GetStruct8/map
Benchmark_GetStruct8/map-12     304303459                3.932 ns/op

Benchmark_GetStruct16
Benchmark_GetStruct16/Mapx
Benchmark_GetStruct16/Mapx-12   278751666                4.263 ns/op
Benchmark_GetStruct16/map
Benchmark_GetStruct16/map-12    203867362                5.844 ns/op

Benchmark_GetStruct20
Benchmark_GetStruct20/Mapx
Benchmark_GetStruct20/Mapx-12   217390791                5.546 ns/op
Benchmark_GetStruct20/map
Benchmark_GetStruct20/map-12    197844157                6.040 ns/op

Benchmark_GetStructPtr8
Benchmark_GetStructPtr8/Mapx
Benchmark_GetStructPtr8/Mapx-12                 375415772                3.197 ns/op
Benchmark_GetStructPtr8/map
Benchmark_GetStructPtr8/map-12                  305616154                3.934 ns/op

Benchmark_GetStructPtr16
Benchmark_GetStructPtr16/Mapx
Benchmark_GetStructPtr16/Mapx-12                220163488                5.517 ns/op
Benchmark_GetStructPtr16/map
Benchmark_GetStructPtr16/map-12                 200725656                5.956 ns/op

Benchmark_GetStructPtr20
Benchmark_GetStructPtr20/Mapx
Benchmark_GetStructPtr20/Mapx-12                232454480                5.164 ns/op
Benchmark_GetStructPtr20/map
Benchmark_GetStructPtr20/map-12                 195612604                6.132 ns/op
```