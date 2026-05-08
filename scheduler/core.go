package scheduler

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	workerCount = 4
	channelSize = 10 //size can be increased based on load
	interval    = 5
)

type Scheduler struct {
	jobs  map[int64][]Job
	ch    chan []Job
	sleep chan int64

	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

type Job struct {
	ID         string
	RetryCount int
	Timestamp  int64
}

func New() *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())

	s := &Scheduler{
		jobs:   map[int64][]Job{},
		ch:     make(chan []Job, channelSize),
		sleep:  make(chan int64, channelSize),
		ctx:    ctx,
		cancel: cancel,
	}

	for i := 0; i < workerCount; i++ {
		go s.Worker(i)
	}

	go s.Sleep()

	return s
}

func (s *Scheduler) AddJob(jobs []Job) {
	timestamp := time.Now().Add(interval * time.Second).Unix()
	jobs[0].Timestamp = timestamp
	s.jobs[timestamp] = jobs
	s.sleep <- timestamp
}

func (s *Scheduler) Worker(i int) {
	defer s.Done()
	s.Add()
	fmt.Println("initializing worker no", i)
	for {
		select {
		case jobs := <-s.ch:
			s.Process(i, jobs)
		case <-s.ctx.Done():
			fmt.Println("exiting worker no", i)
			return
		}
	}
}

func (s *Scheduler) Sleep() {
	defer s.Done()
	s.Add()

	for {
		select {
		case timestamp := <-s.sleep:
			difference := timestamp - time.Now().Unix()
			<-time.After(time.Second * time.Duration(difference))
			s.ch <- s.jobs[timestamp]
		case <-s.ctx.Done():
			fmt.Println("exiting the sleep function")
			return
		}
	}
}

func (s *Scheduler) Process(i int, jobs []Job) {
	retry := []Job{}
	for _, job := range jobs {
		if job.RetryCount == 3 {
			fmt.Println("retry count limit reached. dropping job no ", job.ID)
			continue
		}

		fmt.Println("job no", job.ID, "processed by worker no", i)

		if !Success() {
			retry = append(retry, Job{
				ID:         job.ID,
				RetryCount: job.RetryCount + 1,
			})
		}
	}

	//failed jobs get pushed back
	if len(retry) > 0 {
		s.AddJob(retry)
	}

	//remove the entry from the job map
	s.Delete(jobs[0].Timestamp)
}

// used for mocking status
func Success() bool {
	return rand.IntN(2) == 1
}

func (s *Scheduler) Close() {
	s.cancel()
	s.wg.Wait()
}

func (s *Scheduler) Add() {
	s.wg.Add(1)
}

func (s *Scheduler) Done() {
	s.wg.Done()
}

func (s *Scheduler) Delete(timestamp int64) {
	delete(s.jobs, timestamp)
}

func (s *Scheduler) Store() {

}

func (s *Scheduler) Load() {

}
