package mapx

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

type MapTestS struct {
	id   int
	id2  int32
	str  string
	str2 string
}

func TestNewMapx(t *testing.T) {
	_ = NewMapx[int32, int32](16)
	_ = NewMapx[int, int32](16)
	_ = NewMapx[string, string](16)
	_ = NewMapx[int32, string](16)
	_ = NewMapx[int, string](16)
	_ = NewMapx[string, []byte](16)
	_ = NewMapx[*int, string](16)
	_ = NewMapx[*int32, *int32](16)
	_ = NewMapx[int, *MapTestS](16)
	_ = NewMapx[int, MapTestS](16)
}

func TestMapx_Get(t *testing.T) {
	m := NewMapx[int, int32](16)
	for i := 0; i < 16; i++ {
		m.Set(int(i), int32(i))
	}

	for i := 0; i < 50; i++ {
		n := rand.Int31n(16)
		v := m.Get(int(n))
		fmt.Println(v)
	}
}

func TestMapx_GetStruct(t *testing.T) {
	m := NewMapx[int, MapTestS](16)
	for i := 0; i < 16; i++ {
		m.Set(int(i), MapTestS{
			id:   int(rand.Int31n(16)),
			id2:  rand.Int31n(16),
			str:  strconv.Itoa(int(rand.Int31n(16))),
			str2: strconv.Itoa(int(rand.Int31n(16))),
		})
	}

	for i := 0; i < 50; i++ {
		n := rand.Int31n(16)
		v := m.Get(int(n))
		fmt.Println(v)
	}
}

func TestMapx_Get_Map(t *testing.T) {
	m := NewMapx[int, int32](8)
	for i := 0; i < 16; i++ {
		m.Set(int(i), int32(i))
	}

	for i := 0; i < 50; i++ {
		n := rand.Int31n(16)
		v := m.Get(int(n))
		fmt.Println(v)
	}

	for i := 0; i < 10; i++ {
		m.Del(int(i))
	}

	for i := 0; i < 20; i++ {
		n := rand.Int31n(16)
		v := m.Get(int(n))
		fmt.Println(v)
	}

}

func Benchmark_Get8(b *testing.B) {
	mx := NewMapx[int, int32](8)
	m := make(map[int]int32, 8)
	for i := 0; i < 8; i++ {
		mx.Set(int(i), int32(i))
		m[i] = int32(i)
	}

	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 8)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%8]
			_ = i2
		}
	})
}

func Benchmark_Get16(b *testing.B) {
	mx := NewMapx[int, int32](16)
	m := make(map[int]int32, 16)
	for i := 0; i < 16; i++ {
		mx.Set(int(i), int32(i))
		m[i] = int32(i)
	}

	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 16)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%16]
			_ = i2
		}
	})
}

func Benchmark_Get20(b *testing.B) {
	mx := NewMapx[int, int32](20)
	m := make(map[int]int32, 20)
	for i := 0; i < 20; i++ {
		mx.Set(int(i), int32(i))
		m[i] = int32(i)
	}

	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 20)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%20]
			_ = i2
		}
	})
}

func Benchmark_GetStruct8(b *testing.B) {
	mx := NewMapx[int, MapTestS](8)
	m := make(map[int]MapTestS, 8)
	for i := 0; i < 8; i++ {
		v := MapTestS{
			id:   int(rand.Int31n(100)),
			id2:  rand.Int31n(100),
			str:  strconv.Itoa(int(rand.Int31n(100))),
			str2: strconv.Itoa(int(rand.Int31n(100))),
		}
		mx.Set(int(i), v)
		m[i] = v
	}

	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 8)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%8]
			_ = i2
		}
	})
}

func Benchmark_GetStruct16(b *testing.B) {
	mx := NewMapx[int, MapTestS](16)
	m := make(map[int]MapTestS, 16)
	for i := 0; i < 16; i++ {
		v := MapTestS{
			id:   int(rand.Int31n(100)),
			id2:  rand.Int31n(100),
			str:  strconv.Itoa(int(rand.Int31n(100))),
			str2: strconv.Itoa(int(rand.Int31n(100))),
		}
		mx.Set(int(i), v)
		m[i] = v
	}
	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 16)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%16]
			_ = i2
		}
	})
}

func Benchmark_GetStruct20(b *testing.B) {
	mx := NewMapx[int, MapTestS](20)
	m := make(map[int]MapTestS, 20)
	for i := 0; i < 20; i++ {
		v := MapTestS{
			id:   int(rand.Int31n(100)),
			id2:  rand.Int31n(100),
			str:  strconv.Itoa(int(rand.Int31n(100))),
			str2: strconv.Itoa(int(rand.Int31n(100))),
		}
		mx.Set(int(i), v)
		m[i] = v
	}

	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 20)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%20]
			_ = i2
		}
	})
}

func Benchmark_GetStructPtr8(b *testing.B) {
	mx := NewMapx[int, *MapTestS](8)
	m := make(map[int]*MapTestS, 8)
	for i := 0; i < 8; i++ {
		v := MapTestS{
			id:   int(rand.Int31n(100)),
			id2:  rand.Int31n(100),
			str:  strconv.Itoa(int(rand.Int31n(100))),
			str2: strconv.Itoa(int(rand.Int31n(100))),
		}
		mx.Set(int(i), &v)
		m[i] = &v
	}

	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 8)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%8]
			_ = i2
		}
	})
}

func Benchmark_GetStructPtr16(b *testing.B) {
	mx := NewMapx[int, *MapTestS](16)
	m := make(map[int]*MapTestS, 16)
	for i := 0; i < 16; i++ {
		v := MapTestS{
			id:   int(rand.Int31n(100)),
			id2:  rand.Int31n(100),
			str:  strconv.Itoa(int(rand.Int31n(100))),
			str2: strconv.Itoa(int(rand.Int31n(100))),
		}
		mx.Set(int(i), &v)
		m[i] = &v
	}
	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 16)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%16]
			_ = i2
		}
	})
}

func Benchmark_GetStructPtr20(b *testing.B) {
	mx := NewMapx[int, *MapTestS](20)
	m := make(map[int]*MapTestS, 20)
	for i := 0; i < 20; i++ {
		v := MapTestS{
			id:   int(rand.Int31n(100)),
			id2:  rand.Int31n(100),
			str:  strconv.Itoa(int(rand.Int31n(100))),
			str2: strconv.Itoa(int(rand.Int31n(100))),
		}
		mx.Set(int(i), &v)
		m[i] = &v
	}

	b.Run("Mapx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mx.Get(i % 20)
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			i2 := m[i%20]
			_ = i2
		}
	})
}
