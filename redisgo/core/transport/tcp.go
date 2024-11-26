package transport

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Config struct {
	Address string `yaml:"address"`
}

var ClientCounter int32

func ListenAndServeWithSignal(cfg *Config, tcpHandler Handler) error {
	closeChan := make(chan struct{})
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	listenAndServe(listener, tcpHandler, closeChan)
	return nil
}

func listenAndServe(listener net.Listener, tcpHandler Handler, closeChan <-chan struct{}) {
	// listen signal
	errCh := make(chan error, 1)
	defer close(errCh)
	go func() {
		select {
		case <-closeChan:
			log.Printf("get exit signal")
		case er := <-errCh:
			log.Printf("accept error: %s", er.Error())
		}
		_ = listener.Close() // listener.Accept() will return err immediately
		_ = tcpHandler.Close()
	}()

	ctx := context.Background()
	var waitDone sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			// learn from net/http/serve.go#Serve()
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				log.Printf("accept occurs temporary error: %v, retry in 5ms", err)
				time.Sleep(5 * time.Millisecond)
				continue
			}
			errCh <- err
			break //fixme-hl：一次accept失败,不应该退出服务?
		}
		// handle
		log.Printf("accept link")
		ClientCounter++
		waitDone.Add(1)
		go func() {
			defer func() {
				waitDone.Done()
				atomic.AddInt32(&ClientCounter, -1)
			}()
			// 核心处理逻辑入口(必须保证并发安全)
			tcpHandler.Handle(ctx, conn)
		}()
	}
	waitDone.Wait()
}
