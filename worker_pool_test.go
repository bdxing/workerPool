package workerPool

import (
	"log"
	"testing"
)

func TestWorkerPoolStartStopSerial(t *testing.T) {
	testWorkerPoolStartStop(t)
}

func testWorkerPoolStartStop(t *testing.T) {
	wp := &WorkerPool{
		WorkerFunc:     func(conn interface{}) {},
		MaxWorkerCount: 10,
	}
	for i := 0; i < 10; i++ {
		wp.Start()
		wp.Stop()
	}
}

type TestAdd struct {
	a int
	b int
}

func TestWorkerPool_Serve(t *testing.T) {
	wp := &WorkerPool{
		WorkerFunc: func(i interface{}) {
			ta := i.(*TestAdd)
			ta.a += ta.b
		},
		MaxWorkerCount: DefaultConcurrency,
	}
	wp.Start()

	for i := 0; i < 100000000; i++ {
		if !wp.Serve(&TestAdd{
			a: i,
			b: i + 1,
		}) {
			log.Printf("wp.Serve(): timeout\n")
		}
	}
	t.Logf("workerCount: %v\n", wp.workersCount)
}
