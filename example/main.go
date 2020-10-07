/*
 * Copyright (c) 2018-2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

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
