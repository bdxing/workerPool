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
		WorkerFunc: func(conn interface{}) {

		},
		MaxWorkerCount: 10,
	}
	for i := 0; i < 10; i++ {
		wp.Start()
		wp.Stop()
	}
}

func TestWorkerPool_Listen(t *testing.T) {
	wp := &WorkerPool{
		WorkerFunc: func(c interface{}) {
			conn := c.(net.Conn)
			time.Sleep(1e7)
			conn.Close()
		},
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
