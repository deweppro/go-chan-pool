package pool

import (
	"testing"
)

type test struct {
	I int64
}

//TestChanPool test pool
func TestChanPool(t *testing.T) {
	pool := ChanPool{
		Size: 1,
		New: func() interface{} {
			return &test{}
		},
	}

	a := pool.Get().(*test)

	a.I = 1234

	pool.Put(a)

	b := pool.Get().(*test)

	if b.I != a.I {
		t.Fail()
	}
}

func BenchmarkChanPool_1(b *testing.B) {
	pool := ChanPool{
		Size: 1,
		New: func() interface{} {
			return &test{}
		},
	}

	for n := 0; n < b.N; n++ {
		a := pool.Get().(*test)
		a.I = 1234
		pool.Put(a)
	}
}

func BenchmarkChanPool_100(b *testing.B) {
	pool := ChanPool{
		Size: 100,
		New: func() interface{} {
			return &test{}
		},
	}

	for n := 0; n < b.N; n++ {
		a := pool.Get().(*test)
		a.I = 1234
		pool.Put(a)
	}
}

func BenchmarkChanPool_10000(b *testing.B) {
	pool := ChanPool{
		Size: 10000,
		New: func() interface{} {
			return &test{}
		},
	}

	for n := 0; n < b.N; n++ {
		a := pool.Get().(*test)
		a.I = 1234
		pool.Put(a)
	}
}
