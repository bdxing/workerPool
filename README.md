### Worker Pool
net connection short handler goruntime pool.

> to view `fasthttp` code, found a good library.

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

// The worker handling method, if it is a long connection, is best here to be an interim process.

// For example, the connection authentication authorization and other operations,
// if the completion of authentication authorization, the best to the later processing module,
// or can not play the maximum performance of the reuse work pool
func handler(conn net.Conn) error {
	// For example: connection validation
	// time.Sleep(1e7)

	// For example: verification success, transfer to the subsequent module processing
	// logic <- conn

	return nil
}

```