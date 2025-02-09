package usecase

import (
	"context"

	"runner"
	"runner/_example/jobs"
)

type WorkStorage interface {
	runner.Storage
	jobs.IncCounterStorage
}

func Works[S WorkStorage]() func(context.Context, S) error {
	return runner.New[S]().
		Add(jobs.IncCounter[S]()).
		Add(jobs.IncCounter[S]()).
		Add(jobs.IncCounter[S]()).
		Add(func(_ context.Context, storage S) error {
			storage.Close()

			return nil
		}).
		Add(jobs.IncCounter[S]()).
		Run
}
