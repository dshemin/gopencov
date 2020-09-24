package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dshemin/gopencov/internal/config"
	"github.com/dshemin/gopencov/internal/logger"
	"github.com/dshemin/gopencov/internal/server"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Got error while execute application: %s", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Collect()
	if err != nil {
		return fmt.Errorf("cannot collect configuration: %w", err)
	}

	log, err := logger.New(cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("cannot initialize logger: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Got panic %q: %s", r, string(debug.Stack()))
		}
	}()
	log.Infof("Start application with settings %s", spew.Sdump(cfg))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer log.Debugf("Stop API server on %q", cfg.Address)
		defer wg.Done()
		log.Infof("Start API server on %q", cfg.Address)
		if err := server.Run(ctx, log, cfg.Address); err != nil {
			log.Errorf("API server error: %s", err)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	log.Infof("Receive signal %q. Stopping ...", <-c)
	cancel()

	if !waitWithTimeout(&wg, cfg.StopTimeout) {
		return fmt.Errorf("cannot gracefully stop within %s duration", cfg.StopTimeout)
	}
	return nil
}

func waitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	stop := make(chan struct{})
	go func() {
		wg.Wait()
		close(stop)
	}()

	select {
	case <-stop:
		return true
	case <-time.After(timeout):
		return false
	}
}
