package pool

import (
	"sync"
)

//ChanPool pool object
type ChanPool struct {
	c     chan interface{}
	Size  int                //Size pool size
	empty interface{}        //empty object
	New   func() interface{} //New create object for pool
	sync.Mutex
}

func (p *ChanPool) new() {
	p.Lock()
	if p.c == nil {
		if p.Size <= 0 {
			p.Size = 1
		}
		p.c = make(chan interface{}, p.Size)
	}
	p.Unlock()
}

//Get get object from pool
func (p *ChanPool) Get() interface{} {
	p.new()

	select {
	case x := <-p.c:
		return x

	default:
		if p.New != nil {
			return p.New()
		}
		return p.empty
	}

}

//Put put object to pool
func (p *ChanPool) Put(x interface{}) {
	if x != nil {
		select {
		case p.c <- x:
		default:

		}
	}
}
