package main

import (
	"context"
	"fmt"
)

type TaskStorage struct {
	Storage
	counter int8
}

func (s *TaskStorage) Counter() int8 {
	return s.counter
}

func (s *TaskStorage) SetCounter(counter int8) {
	s.counter = counter
}

func NewStorage() *TaskStorage {
	return &TaskStorage{
		Storage: &DefaultStorage{},
	}
}

type MainInterface interface {
	TaskMethodInterface
}

func main() {
	ctx := context.Background()
	storage := NewStorage()
	runner := New[MainInterface]().
		Add(task1[MainInterface]()).
		Add(task2[MainInterface]()).
		Add(task3[MainInterface]()).
		Run

	err := runner(ctx, storage)
	if err != nil {
		return
	}
}

type TaskMethodInterface interface {
	Storage
	Counter() int8
	SetCounter(int8)
}

func task1[S TaskMethodInterface]() Task[S] {
	return func(ctx context.Context, store S) error {
		count := store.Counter()
		fmt.Println("counter:", count)
		store.SetCounter(count + 1)
		return nil
	}
}

func task2[S TaskMethodInterface]() Task[S] {
	return func(ctx context.Context, store S) error {
		count := store.Counter()
		fmt.Println("counter:", count)
		store.SetCounter(count + 1)
		return nil
	}
}

func task3[S TaskMethodInterface]() Task[S] {
	return func(ctx context.Context, store S) error {
		count := store.Counter()
		fmt.Println("counter:", count)
		store.SetCounter(count + 1)
		return nil
	}
}
