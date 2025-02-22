package service

import (
	"crave/hub/cmd/model"
	repository "crave/hub/cmd/target/cmd/api/infrastructure/repository"
	"fmt"
)

type Service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Init(origin, destination *model.Target) error {
	errChan := make(chan error, 2)

	go func() {
		_, err := s.repo.Create(origin)
		errChan <- err
	}()
	go func() {
		_, err := s.repo.Create(destination)
		errChan <- err
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return fmt.Errorf("failed to create target: %w", err)
		}
	}

	return nil
}
