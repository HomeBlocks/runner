package usecase

import "runner"

type Storage struct {
	runner.DefaultStorage

	counter int8
}

func (s *Storage) Counter() int8 {
	return s.counter
}

func (s *Storage) SetCounter(counter int8) {
	s.counter = counter
}

func NewStorage() *Storage {
	return &Storage{}
}
