package mapx

import (
	"sync"
)

type Mapx[K comparable, V any] struct {
	keys []K
	vals []V

	// 支持最大的不使用 map 的大小
	maxSize int

	// 当前KV的大小
	len int

	m    map[K]V
	once *sync.Once
}

func NewMapx[K comparable, V any](size int) *Mapx[K, V] {
	return &Mapx[K, V]{
		keys:    make([]K, 0, size),
		vals:    make([]V, 0, size),
		maxSize: size,
		once:    &sync.Once{},
	}
}

func (m *Mapx[K, V]) Set(k K, v V) {
	if m.len > m.maxSize {
		m.insertmap(k, v)
	} else if idx := m.index(k); idx != -1 {
		m.update(idx, v)
	} else {
		m.insert(k, v)
	}
}

func (m *Mapx[K, V]) Get(k K) (v V) {
	if m.len > m.maxSize {
		return m.m[k]
	}

	if idx := m.index(k); idx != -1 {
		return m.vals[idx]
	}
	return v

}

func (m *Mapx[K, V]) GetOk(k K) (v V, ok bool) {
	if m.len > m.maxSize {
		v, ok = m.m[k]
		return
	}

	if idx := m.index(k); idx != -1 {
		return m.vals[idx], true
	}

	return v, false
}

func (m *Mapx[K, V]) Del(k K) {
	if m.len > m.maxSize {
		delete(m.m, k)
		m.len = len(m.m)

		if m.len <= m.maxSize {
			m.narrow()
		}
		return
	}

	if idx := m.index(k); idx != -1 {
		m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
		m.vals = append(m.vals[:idx], m.vals[idx+1:]...)
		m.len--
		m.keys = m.keys[:m.len]
		m.vals = m.vals[:m.len]
	}
}

func (m *Mapx[K, V]) update(idx int, v V) {
	m.vals[idx] = v
}

func (m *Mapx[K, V]) insert(k K, v V) {
	if m.len < m.maxSize {
		m.keys = append(m.keys, k)
		m.vals = append(m.vals, v)
		m.len++

		return
	}

	for i := 0; i < m.len; i++ {
		m.migrateMap(m.keys[i], m.vals[i])
	}
	m.insertmap(k, v)
}

func (m *Mapx[K, V]) insertmap(k K, v V) {
	m.m[k] = v
	m.len = len(m.m)
}

func (m *Mapx[K, V]) migrateMap(k K, v V) {
	m.once.Do(func() {
		m.m = make(map[K]V, m.maxSize+1)
	})

	m.m[k] = v
}

func (m *Mapx[K, V]) narrow() {
	i := 0
	for k2, v2 := range m.m {
		m.keys[i] = k2
		m.vals[i] = v2
		delete(m.m, k2)
		i++
	}
	m.keys = m.keys[:i]
	m.vals = m.vals[:i]
}

func (m *Mapx[K, V]) IsExistsKey(k K) bool {
	return m.index(k) != -1
}

func (m *Mapx[K, V]) index(k K) int {
	if m.len == 0 {
		return -1
	}

	if m.len > m.maxSize {
		panic("The data is in the map, this method should not be used")
	}

	for i, key := range m.keys {
		if key == k {
			return i
		}
	}
	return -1
}

func (m *Mapx[K, V]) Len() int {
	return m.len
}

// Range 遍历map元素，最好不要通过它来循环删除元素
func (m *Mapx[K, V]) Range(f func(k K, v V)) {
	if m.len > m.maxSize {
		for k, v := range m.m {
			f(k, v)
		}
		return
	}

	for i := 0; i < m.len; i++ {
		f(m.keys[i], m.vals[i])
	}
}
