package generic

import "sync/atomic"

type Job = func()

type JobQueue struct {
	waiting     atomic.Int32
	workerCount int
	ch          chan Job
}

func MakeJobQueue(workerCount int) *JobQueue {
	jq := &JobQueue{
		workerCount: workerCount, ch: make(chan Job),
	}
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range jq.ch {
				job()
				jq.waiting.Add(-1)
			}
		}()
	}
	return jq
}

func (jq *JobQueue) Submit(job Job) {
	jq.waiting.Add(1)
	jq.ch <- job
}

func (jq *JobQueue) Close() {
	close(jq.ch)
}

func (jq *JobQueue) WaitingCount() int {
	return int(jq.waiting.Load())
}
