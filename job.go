package runner

import "context"

// Job represents a function that performs work using the provided context and storage.
type Job[S Storage] func(context.Context, S) error

// JobOption is a function that modifies the internal job configuration.
type JobOption[S Storage] func(*jobInternal[S])

// jobInternal represents the internal job structure, including an optional error handler.
type jobInternal[S Storage] struct {
	job          Job[S]
	errorHandler func(error) error
}

// defaultJobHandler executes the job and returns any errors encountered.
func defaultJobHandler[S Storage](job Job[S], ctx context.Context, storage S) error {
	return job(ctx, storage)
}
