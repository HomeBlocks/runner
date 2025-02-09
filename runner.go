package runner

import (
	"context"

	"github.com/pkg/errors"
)

const defaultJobCount = 5

type Runner[S Storage] struct {
	jobs []jobInternal[S]
}

// internalRun executes a job and handles its error using the job's custom error handler.
func (r *Runner[S]) internalRun(ctx context.Context, job jobInternal[S], storage S) error {
	select {
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "context canceled")
	default:
		err := job.job(ctx, storage)
		if job.errorHandler != nil {
			err = job.errorHandler(err)
		}

		return errors.WithMessagef(err, "run job fail")
	}
}

// Run executes all jobs, respecting cancellation via context and storage cancellation.
func (r *Runner[S]) Run(ctx context.Context, storage S) error {
	for _, job := range r.jobs {
		if storage.IsClosed() {
			break
		}

		if err := r.internalRun(ctx, job, storage); err != nil {
			return err
		}
	}

	return nil
}

// Add adds a job to the runner with optional job-specific configurations.
func (r *Runner[S]) Add(job Job[S]) *Runner[S] {
	ji := jobInternal[S]{job: job}
	r.jobs = append(r.jobs, ji)

	return r
}

// New creates a new Runner with the default handler and an empty list of jobs.
func New[S Storage]() *Runner[S] {
	return &Runner[S]{
		jobs: make([]jobInternal[S], 0, defaultJobCount),
	}
}
