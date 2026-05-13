# Job Scheduler
- The job scheduler executes a batch of jobs. 
- The failed jobs are pushed into a queue and executed again after a configured time interval.
- One of the goroutine workers will pick the batch to be executed.
- This process repeats itself until maxRetryCount limit is reached.
- The program gracefully shuts down by waiting for pending tasks to be completely executed first.
