package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	registerChan    chan string
	unregisterChan  chan string
	kickHandlerChan chan string
}

func NewWorker() *Worker {
	work := &Worker{
		registerChan:    make(chan string),
		unregisterChan:  make(chan string),
		kickHandlerChan: make(chan string),
	}

	go func() {
		for {
			work.registerChan <- "123"
			work.unregisterChan <- "456"
			work.kickHandlerChan <- "789"
		}
	}()

	return work
}

func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			slog.Info("Received signal:", "start done1", "done")
			return
		default:
		}

		slog.Info("Received signal:", "start done???", "done")

		select {
		case <-ctx.Done():
			slog.Info("Received signal:", "start done2", "done")
			return
		case client := <-w.registerChan:
			w.handleRegistration(ctx, client)
		case client := <-w.unregisterChan:
			w.handleUnregistration(ctx, client)
		case onlineInfo := <-w.kickHandlerChan:
			w.handleKick(ctx, onlineInfo)
		}
	}
}

func (w *Worker) handleRegistration(ctx context.Context, client string) {
	// Your registration logic here
	select {
	case <-ctx.Done():
		slog.Info("Received signal:", "handleRegistration done", "done")
	}
}

func (w *Worker) handleUnregistration(ctx context.Context, client string) {
	// Your unregistration logic here
	select {
	case <-ctx.Done():
		slog.Info("Received signal:", "handleUnregistration done", "done")
		return
	default:
		// Your unregistration logic here
		time.Sleep(2 * time.Second) // Simulate some work
	}
}

func (w *Worker) handleKick(ctx context.Context, onlineInfo string) {
	// Your kick logic here
	select {
	case <-ctx.Done():
		slog.Info("Received signal:", "handleKick done", "done")
		return
	default:
		// Your kick logic here
		time.Sleep(2 * time.Second) // Simulate some work
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	worker := NewWorker()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		worker.Start(ctx)
	}()

	// Wait for a signal to gracefully shut down the application
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		// Received signal, shutting down...
		slog.Info("Received signal:", "signal", sig)
		cancel()
	}

	// Wait for worker to finish
	wg.Wait()
	slog.Info("Application gracefully shut down.")
}
