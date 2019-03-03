package workerPool

import (
	"testing"
)

func TestWorkerPoolStartStop(t *testing.T) {
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

	for i := 0; i < 3000000; i++ {
		if !wp.Serve(&TestAdd{
			a: i,
			b: i + 1,
		}) {
			t.Logf("wp.Serve(): timeout\n")
		}
	}
	t.Logf("workerCount: %v\n", wp.workersCount)
}

func BenchmarkWorkerPool_Serve(b *testing.B) {
	b.ReportAllocs()

	s := make(chan struct{}, b.N)
	wp := &WorkerPool{
		WorkerFunc: func(i interface{}) {
			ta := i.(*TestAdd)
			ta.a += ta.b
			s <- struct{}{}
		},
		MaxWorkerCount: DefaultConcurrency,
	}
	wp.Start()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !wp.Serve(&TestAdd{
			a: i,
			b: i + 1,
		}) {
			b.Logf("wp.Serve(): timeout\n")
		}
	}

	sNum := 0
	for {
		<-s
		sNum++
		if sNum == b.N {
			break
		}
	}

	b.Logf("taskCount: %10d, workerCount: %10d\n", b.N, wp.workersCount)
}
