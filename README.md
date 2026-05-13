Job Scheduler executes a batch of jobs. 
The failed jobs are pushed into a queue and executed again after a configured time interval.
This process repeats itself until maxRetryCount limit is reached.
The program gracefully shuts down by waiting for pending tasks to be completely executed first.
