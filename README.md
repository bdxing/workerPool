## <p align="center">workerPool</p>
This is a goroutine pool, which can avoid a large amount of performance consumption of creation and destruction under high concurrency, ensure the stable scheduling of modules, and automatically scale the size of the co-program pool to fit the current business scheduling.

## Installation

To install this package, you need to setup your Go workspace.  The simplest way to install the library is to run:

```
go get github.com/bdxing/workerPool
```

## Example 

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
	workerFunc := func(tmp interface{}) {
		ta := tmp.(*TestAdd)
		ta.a += ta.b
	}

	wp := &WorkerPool{
		MaxWorkerCount:        DefaultConcurrency, // max worker goroutine number, Hot add
		MaxIdleWorkerDuration: 5 * time.Second,    // worker goroutine max Idle Worker Duration
		WorkerFunc:            workerFunc,         // worker method
	}
	wp.Start()

	nowTime := time.Now()

	for i := 0; i < 100000000; i++ {
		if !wp.Serve(&TestAdd{
			a: i,
			b: i + 1,
		}) {
			log.Printf("wp.Serve(): timeout\n")
		}
	}

	log.Printf("consuming time: %v\n", time.Now().Sub(nowTime))

	// shutdown worker pool
	//wp.Stop()

}
```

## Benchmark

```text
goos: windows
goarch: amd64
pkg: github.com/bdxing/workerPool
cpu: Intel(R) Core(TM) i5-10400F CPU @ 2.90GHz
BenchmarkWorkerPool_Serve
    worker_pool_test.go:80: taskCount:          1, workerCount:          1
    worker_pool_test.go:80: taskCount:        100, workerCount:         16
    worker_pool_test.go:80: taskCount:      10000, workerCount:         45
    worker_pool_test.go:80: taskCount:    1000000, workerCount:        111
    worker_pool_test.go:80: taskCount:    2567829, workerCount:        206
BenchmarkWorkerPool_Serve-12    	 2567829	       467.9 ns/op	      16 B/op	       1 allocs/op
PASS
```