package generic

type Job = func()

type chunk[T any] struct {
	items []T
	next  *chunk[T]
}

type chunkedQueue[T any] struct {
	// invariant: head and tail are never allowed to be nil!
	head *chunk[T]
	tail *chunk[T]

	headIdx int
	tailIdx int

	pool *chunk[T]

	chunkSize int
}

func (chqu *chunkedQueue[T]) peek() (result T, found bool) {
	if chqu.head == chqu.tail && chqu.headIdx == chqu.tailIdx {
		return
	}
	return chqu.head.items[chqu.headIdx], true
}

// must only be called after a succesful peek
func (chqu *chunkedQueue[T]) consume() {
	// assume head exists
	chqu.headIdx++
	// reset head and tail indecies and add to pool
	if chqu.headIdx == len(chqu.head.items) {
		// move the head to the next chunk
		// unless this is the last chunk, just reset it
		if chqu.head.next == nil {
			Assert(chqu.head == chqu.tail, "head has no next but tail is different!")
			chqu.headIdx = 0
			chqu.tailIdx = 0
		} else {
			// move head node to pool
			n := chqu.head
			chqu.head = n.next
			chqu.headIdx = 0
			n.next = chqu.pool
			chqu.pool = n
		}
	}
}

func allocChunk[T any](size int) (n *chunk[T]) {
	n = new(chunk[T])
	n.items = make([]T, size)
	return n
}

func (chqu *chunkedQueue[T]) allocChunk() *chunk[T] {
	if chqu.pool != nil {
		n := chqu.pool
		chqu.pool = n.next
		n.next = nil
		return n
	} else {
		return allocChunk[T](chqu.chunkSize)
	}
}

func (chqu *chunkedQueue[T]) push(item T) {
	chqu.tail.items[chqu.tailIdx] = item
	chqu.tailIdx++
	if chqu.tailIdx == len(chqu.tail.items) {
		chqu.tail.next = chqu.allocChunk()
		chqu.tail = chqu.tail.next
		chqu.tailIdx = 0
	}
}

func createChunkedQueue[T any](chunkSize int) *chunkedQueue[T] {
	Assert(chunkSize > 0, "invalid chunk size")
	chqu := new(chunkedQueue[T])
	chqu.chunkSize = chunkSize
	chqu.head = allocChunk[T](chunkSize)
	chqu.tail = chqu.head
	return chqu
}

type JobQueue struct {
	submitCh  chan Job
	workersCh chan Job
}

func MakeJobQueue(workerCount int) *JobQueue {
	jq := &JobQueue{
		submitCh:  make(chan Job),
		workersCh: make(chan Job),
	}

	// launch workers goroutines
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range jq.workersCh {
				job()
			}
		}()
	}

	// launch dispatch coordinator goroutine
	go func() {
		chqu := createChunkedQueue[Job](4096)
		for {
			var pushCh chan Job
			peek, found := chqu.peek()
			if found {
				pushCh = jq.workersCh
			}
			select {
			case pushCh <- peek:
				chqu.consume()
			case job := <-jq.submitCh:
				chqu.push(job)
			}
		}
	}()
	return jq
}

func (jq *JobQueue) Submit(job Job) {
	jq.submitCh <- job
}
