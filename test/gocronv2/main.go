package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

func main() {
	// create a scheduler
	s, err := gocron.NewScheduler(gocron.WithStopTimeout(time.Second * 30))
	if err != nil {
		// handle error
	}

	var i *int = new(int)
	*i = 1
	// add a job to the scheduler
	j, err := s.NewJob(
		gocron.DurationJob(
			1*time.Second,
		),
		gocron.NewTask(
			func() {
				*i++
				data := *i
				fmt.Println("Job exec start", data)
				time.Sleep(time.Second * 10)
				fmt.Println("Job exec end", data)
				panic(errors.New("panic test"))
			},
		),
		gocron.WithEventListeners(gocron.BeforeJobRuns(
			func(jobID uuid.UUID, jobName string) {
				if err := recover(); err != nil {
					fmt.Println("panic!!", err)
				}
			},
		)),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.Start()

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, os.Kill)
	<-shutdownSignal

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
	}
}
