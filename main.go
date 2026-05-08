package main

import (
	"job-scheduler/scheduler"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	s := scheduler.New()

	jobs := scheduler.GenerateJobs()

	s.AddJob(jobs)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	s.Close()
}
