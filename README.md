# worker pool

This is a goroutine pool, which can avoid a large amount of performance consumption of creation and destruction under high concurrency, ensure the stable scheduling of modules, and automatically scale the size of the co-program pool to fit the current business scheduling.

### Usage 

```go
package main

import (
	. "github.com/bdxing/workerPool"
	"log"
	"time"
)

type TestAdd struct {
	a int
	b int
}

func main() {
	wp := &WorkerPool{
		WorkerFunc: handler,
		MaxWorkerCount: DefaultConcurrency,
	}
	nowTime := time.Now()
	wp.Start()

	for i := 0; i < 100000000; i++ {
		if !wp.Serve(&TestAdd{
			a: i,
			b: i + 1,
		}) {
			log.Printf("wp.Serve(): timeout\n")
		}
	}
	log.Printf("consuming time: %v\n", time.Now().Sub(nowTime))
}

func handler(i interface{}) {
	// For example: connection validation
	ta := i.(*TestAdd)
	ta.a += ta.b

	// For example: verification success, transfer to the subsequent module processing
}
```

### Benchmark

CPU: Core(TM) i7-7700HQ

```text
goarch: amd64
pkg: workerPool
BenchmarkWorkerPool_Serve-8   	 3000000	       515 ns/op	      16 B/op	       1 allocs/op
--- BENCH: BenchmarkWorkerPool_Serve-8
    worker_pool_test.go:77: taskCount: 1, workerCount: 1
    worker_pool_test.go:77: taskCount: 100, workerCount: 16
    worker_pool_test.go:77: taskCount: 10000, workerCount: 62
    worker_pool_test.go:77: taskCount: 1000000, workerCount: 297
    worker_pool_test.go:77: taskCount: 3000000, workerCount: 660
PASS
```