package main

import (
	"context"
)

type DefaultStorage struct {
	cancelled bool
}

func (d *DefaultStorage) Cancel() {
	d.cancelled = true
}

func (d *DefaultStorage) IsCancelled() bool {
	return d.cancelled
}

type Storage interface {
	Cancel()
	IsCancelled() bool
}

type Task[S Storage] func(context.Context, S) error

type taskStruct[S Storage] struct {
	errorHandler func(error) error
	taskFn       Task[S]
}

type Runner[S Storage] struct {
	tasks      []taskStruct[S]
	taskRunner func(Task[S], context.Context, S) error
}

func taskRunner[S Storage](task Task[S], ctx context.Context, storage S) error {
	return task(ctx, storage)
}

type Option[S Storage] func(runner *Runner[S])

func New[S Storage](opts ...Option[S]) Runner[S] {
	r := Runner[S]{
		tasks:      make([]taskStruct[S], 0, 2),
		taskRunner: taskRunner[S],
	}

	for _, opt := range opts {
		opt(&r)
	}

	return r
}

func (r Runner[S]) Run(ctx context.Context, storage S) error {
	return r.linerRun(ctx, storage)
}

func (r Runner[S]) Add(task Task[S]) Runner[S] {
	tasks := taskStruct[S]{
		taskFn: task,
	}

	return Runner[S]{
		tasks:      append(r.tasks, tasks),
		taskRunner: r.taskRunner,
	}
}

func (r Runner[S]) linerRun(ctx context.Context, storage S) error {
	for _, entity := range r.tasks {
		if storage.IsCancelled() {
			break
		}

		var err error

		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
			err = r.taskRunner(entity.taskFn, ctx, storage)
		}

		if entity.errorHandler != nil {
			err = entity.errorHandler(err)
		}

		return err
	}

	return nil
}
