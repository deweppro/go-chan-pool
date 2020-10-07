# Pool on Chan

[![Release](https://img.shields.io/github/release/deweppro/go-chan-pool.svg?style=flat-square)](https://github.com/deweppro/go-chan-pool/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/deweppro/go-chan-pool)](https://goreportcard.com/report/github.com/deweppro/go-chan-pool)
[![Build Status](https://travis-ci.com/deweppro/go-chan-pool.svg?branch=master)](https://travis-ci.com/deweppro/go-chan-pool)

# How to use it

```go
import pool "github.com/deweppro/go-chan-pool"

type Object struct {
	I int64
}

func main() {
	cpool := &pool.ChanPool{
		Size: 1, // pool size
		New: func() interface{} { // handler for creating a new object in pool
			return &Object{}
		},
	}
	a := cpool.Get().(*Object) //get object from pool
	cpool.Put(a) //set object to pool
}
```