package main

import pool "github.com/deweppro/go-chan-pool"

type Object struct {
	I int64
}

func main() {
	cpool := &pool.ChanPool{
		Size: 1,
		New: func() interface{} {
			return &Object{}
		},
	}
	a := cpool.Get().(*Object)
	cpool.Put(a)
}
