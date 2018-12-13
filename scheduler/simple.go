package scheduler

import "learngo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}
//会改变workerChan的内容，所以用指针类型
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}


