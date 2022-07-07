package mapx

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewMapx(t *testing.T) {
	_ = NewMapx[int32, int32](16)
	_ = NewMapx[int, int32](16)
	_ = NewMapx[string, string](16)
	_ = NewMapx[int32, string](16)
	_ = NewMapx[int, string](16)
	_ = NewMapx[string, []byte](16)
	_ = NewMapx[*int, string](16)
	_ = NewMapx[*int32, *int32](16)
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
