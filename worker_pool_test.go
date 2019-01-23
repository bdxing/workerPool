package workerPool

import (
	"log"
	"net"
	"testing"
)

func TestWorkerPoolStartStopSerial(t *testing.T) {
	testWorkerPoolStartStop(t)
}

func testWorkerPoolStartStop(t *testing.T) {
	wp := &WorkerPool{
		WorkerFunc: func(conn net.Conn) error {

			return nil
		},
		MaxWorkerCount: 10,
		Logger:         log.Logger{},
	}
	for i := 0; i < 10; i++ {
		wp.Start()
		wp.Stop()
	}
}

func TestWorkerPool_Serve(t *testing.T) {

}
