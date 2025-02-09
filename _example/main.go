package main

import (
	"context"

	"runner/_example/usecase"
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
