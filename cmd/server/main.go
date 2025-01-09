package main

import (
	"context"
	"fmt"
	"groot/internal/zlog"
	"groot/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer zlog.Sync()
	server, err := router.NewHTTPServer()
	if err != nil {
		fmt.Printf("server init err: %s\n", err)
		os.Exit(1)
	}
	go server.Run()
	GracefullyShutdown(server.Server, func() {})
}

func GracefullyShutdown(server *http.Server, callback func()) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sg := <-ch
	fmt.Printf("receive signal %s\n", sg)
	fmt.Println("shutting down server")
	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Printf("http shutdown err: %s\n", err)
	}
	callback()
	fmt.Println("server exiting")
}
