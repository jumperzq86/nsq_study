package util

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

//exercise: 函数式编程风格，新建协程，调用入参函数
func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
