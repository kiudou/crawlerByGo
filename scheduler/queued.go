package scheduler

import "learngo/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	wokerChan chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.wokerChan <- w
}


func (s *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}


func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.wokerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activityRequest engine.Request
			var activityWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activityRequest = requestQ[0]
				activityWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.wokerChan:
				workerQ = append(workerQ, w)
			case activityWorker <- activityRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
