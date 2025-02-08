package runner

import "context"

type Storage interface {
	Cancel()
	IsCancelled() bool
}

type Task[S Storage] func(context.Context, S) error

type task[S Storage] struct {
	name         string
	errorHandler func(error) error
	taskFn       Task[S]
}

type Runner[S Storage] struct {
	tasks []task[S]
}

func New[S Storage](tasks ...Task[S]) Runner[S] {
	r := Runner[S]{
		tasks: make([]task[S], 0, len(tasks)),
	}

	return r
}
