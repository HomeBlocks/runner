package runner

import "context"

type Job[S Storage] func(context.Context, S) error
type JobOption[S Storage] func(*jobInternal[S])

type jobInternal[S Storage] struct {
	job          Job[S]
	errorHandler func(error) error
}

func defaultJobHandler[S Storage](job Job[S], ctx context.Context, storage S) error {
	return job(ctx, storage)
}
