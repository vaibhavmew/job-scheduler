package scheduler

import "job-scheduler/utils"

func GenerateJobs() []Job {
	jobs := []Job{}
	for i := 0; i < 10; i++ {
		jobs = append(jobs, Job{
			ID:         utils.RandomString(),
			RetryCount: 0,
		})
	}
	return jobs
}
