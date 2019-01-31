# worker pool

这是一个复用协程池，避免在高并发下大量的创建和销毁的性能消耗，保证模块的稳定调度，合理的自动伸缩协程池的大小来适用当前业务调度。

我在阅读 `fasthttp` 项目源码的时候发现，稍作修改。

### Usage 

举例一个工作过程：

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

// 工作回调方法
// 必须要注意的是，如果你想发挥最大性能，你这里不能使用带阻塞的业务代码，如果阻塞时间过长，可能会得不到你想要的性能。

// 正常的使用方式:
// 长阻塞：可以采用编写连接验证授权的代码，验证授权完成后，把连接交给后续模块继续执行即可。
// 短阻塞：可直接写编写业务代码。
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