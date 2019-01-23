package workerPool

import (
	"log"
	"net"
	"testing"
	"time"
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

func TestWorkerPool_Listen(t *testing.T) {
	wp := &WorkerPool{
		WorkerFunc: func(conn net.Conn) error {
			time.Sleep(1e7)
			conn.Close()
			return nil
		},
		LogAllErrors:   false,
		MaxWorkerCount: 10,
		Logger:         log.Logger{},
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

func TestWorkerPool_Dial(t *testing.T) {
	dial := func() {
		conn, er := net.Dial("tcp", "127.0.0.1:6666")
		if er != nil {
			return
		}
		defer conn.Close()

		msg := make([]byte, 1024)
		for {
			n, er := conn.Read(msg)
			if er != nil {
				return
			}
			log.Printf("msg: %v\n", msg[:n])
		}
	}

	for i := 0; i < 2000; i++ {
		dial()
		//time.Sleep(1e6)
	}
	time.Sleep(10e9)
}
