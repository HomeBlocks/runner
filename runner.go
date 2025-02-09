package runner

import (
	"context"
	"github.com/pkg/errors"
)

type Runner[S Storage] struct {
	jobs       []jobInternal[S]
	jobHandler func(Job[S], context.Context, S) error
}

func (r Runner[B]) internalRun(ctx context.Context, job jobInternal[B], buffer B) error {
	var err error

	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		err = r.jobHandler(job.job, ctx, buffer)
	}

	if job.errorHandler != nil {
		err = job.errorHandler(err)
	}

	return errors.WithMessagef(err, "run job fail")
}

func (r Runner[S]) Run(ctx context.Context, storage S) error {
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

func (r Runner[B]) Add(job Job[B]) Runner[B] {
	ji := jobInternal[B]{
		job: job,
	}

	return Runner[B]{
		jobs:       append(r.jobs, ji),
		jobHandler: r.jobHandler,
	}
}

func New[S Storage]() Runner[S] {
	r := Runner[S]{
		jobs:       make([]jobInternal[S], 0, 5),
		jobHandler: defaultJobHandler[S],
	}

	return r
}
