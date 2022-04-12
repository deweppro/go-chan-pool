package pool_test

import (
	"testing"

	pool "github.com/deweppro/go-chan-pool"
)

type test struct {
	I int64
}

//TestChanPool test pool
func TestChanPool(t *testing.T) {
	v := &pool.ChanPool{
		Size: 1,
		New: func() interface{} {
			return &test{}
		},
	}
	a, ok := v.Get().(*test)
	if !ok {
		t.Fail()
	}
	a.I = 1234
	v.Put(a)
	b, ok := v.Get().(*test)
	if !ok {
		t.Fail()
	}
	if b.I != a.I {
		t.Fail()
	}
}

func BenchmarkChanPool_1(b *testing.B) {
	v := &pool.ChanPool{
		Size: 1,
		New: func() interface{} {
			return &test{}
		},
	}
	for n := 0; n < b.N; n++ {
		a, ok := v.Get().(*test)
		if !ok {
			b.Fail()
		}
		a.I = 1234
		v.Put(a)
	}
}

func BenchmarkChanPool_100(b *testing.B) {
	v := &pool.ChanPool{
		Size: 100,
		New: func() interface{} {
			return &test{}
		},
	}
	for n := 0; n < b.N; n++ {
		a, ok := v.Get().(*test)
		if !ok {
			b.Fail()
		}
		a.I = 1234
		v.Put(a)
	}
}

func BenchmarkChanPool_10000(b *testing.B) {
	v := &pool.ChanPool{
		Size: 10000,
		New: func() interface{} {
			return &test{}
		},
	}
	for n := 0; n < b.N; n++ {
		a, ok := v.Get().(*test)
		if !ok {
			b.Fail()
		}
		a.I = 1234
		v.Put(a)
	}
}
