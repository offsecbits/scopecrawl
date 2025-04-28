package ratelimiter

import (
	"time"
)

// Job represents a task to run — usually a fetch URL operation.
type Job struct {
	URL      string
	Callback func(string)
}

// Limiter manages both concurrency and request rate.
type Limiter struct {
	jobQueue    chan Job
	stop        chan struct{}
	concurrency int
	rate        int
}

// NewLimiter creates a Limiter with concurrency and rate limits.
func NewLimiter(concurrency, rate int) *Limiter {
	return &Limiter{
		jobQueue:    make(chan Job),
		stop:        make(chan struct{}),
		concurrency: concurrency,
		rate:        rate,
	}
}

// Start launches the worker pool with rate-limiting applied.
func (l *Limiter) Start() {
	// Rate limiter: one tick per (1/rate) second
	ticker := time.NewTicker(time.Second / time.Duration(l.rate))

	for i := 0; i < l.concurrency; i++ {
		go func() {
			for {
				select {
				case <-l.stop:
					return

				case <-ticker.C:		// Wait for the ticker 1st
					job := <-l.jobQueue     // grab job
					job.Callback(job.URL) // Process job
				}
			}
		}()
	}
}

// Submit adds a job to the queue.
func (l *Limiter) Submit(url string, callback func(string)) {
	l.jobQueue <- Job{URL: url, Callback: callback}
}

// Stop stops all workers gracefully.
func (l *Limiter) Stop() {
	select {
	case <-l.stop:
		// Already closed, do nothing
	default:
		close(l.stop)
	}
}
