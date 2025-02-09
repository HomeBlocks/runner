package runner

import (
	"context"
	"github.com/pkg/errors"
)

type Runner[S Storage] struct {
	jobs       []jobInternal[S]
	jobHandler func(Job[S], context.Context, S) error
}

func (r Runner[B]) internalRun(job jobInternal[B], ctx context.Context, buffer B) error {
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
		if storage.IsCancelled() {
			break
		}

		if err := r.internalRun(job, ctx, storage); err != nil {
			return err
		}
	}

	return nil
}

func (r Runner[B]) Add(job Job[B], opts ...JobOption[B]) Runner[B] {
	ji := jobInternal[B]{
		job: job,
	}

	for _, opt := range opts {
		opt(&ji)
	}

	return Runner[B]{
		jobs:       append(r.jobs, ji),
		jobHandler: r.jobHandler,
	}
}

type Option[S Storage] func(*Runner[S])

func New[S Storage](opts ...Option[S]) Runner[S] {
	p := Runner[S]{
		jobs:       make([]jobInternal[S], 0, 2),
		jobHandler: defaultJobHandler[S],
	}
	for _, opt := range opts {
		opt(&p)
	}
	return p
}
