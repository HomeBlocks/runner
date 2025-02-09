package main

import (
	"context"
	"runner/example/usecase"
)

func main() {
	storage := usecase.NewStorage()
	ctx := context.Background()

	work := usecase.Works[usecase.WorkStorage]()

	err := work(ctx, storage)
	if err != nil {
		return
	}
}
