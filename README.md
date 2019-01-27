# Worker Pool

> 我在阅读 `fasthttp` 源码的时候发现，稍作修改。

# Usage 

这是一个协程池。复用协程，避免大量的创建和销毁的性能消耗，合理的自动伸缩，性能很棒。


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