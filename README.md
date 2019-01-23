### WorkerPool
net connection short handler goruntime pool.

#### Example
```go
package main

import (
	"github.com/bdxing/workerPool"
	"log"
	"net"
)

func main() {
	wp := &workerPool.WorkerPool{
		WorkerFunc:     handler,
		LogAllErrors:   false,
		MaxWorkerCount: 10,
	}
	wp.Start()

	l, er := net.Listen("tcp", "0.0.0.0:6666")
	if er != nil {
		panic(er)
	}
	for {
		conn, er := l.Accept()
		if er != nil {
			continue
		}
		if !wp.Serve(conn) {
			log.Printf("wp.Serve timeout\n")
		}
	}
}

// 工人处理方法，如果是长连接，这里最好是一个过度过程。
// 比如：连接的验证授权等操作，如果完成验证授权，最好转到后面处理模块，不然无法发挥复用工作池的最大性能
func handler(conn net.Conn) error {
	// 比如：连接验证
	// time.Sleep(1e7)

	// 比如：验证成功，移交到后续模块处理
	// logic <- conn

	return nil
}

```