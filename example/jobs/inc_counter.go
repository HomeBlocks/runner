package jobs

import (
	"context"
	"fmt"
)

type IncCounterStorage interface {
	Counter() int8
	SetCounter(int8)
}

func IncCounter[S IncCounterStorage]() func(context.Context, S) error {
	return func(ctx context.Context, storage S) error {
		storageCounter := storage.Counter()
		storage.SetCounter(storageCounter + 1)

		fmt.Println(storageCounter + 1)
		return nil
	}
}
