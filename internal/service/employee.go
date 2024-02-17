package service

import (
	"context"

	"github.com/anil1226/go-employee-dynamo/internal/models"
)

// type EmpService struct {
// 	db.EmpStore
// }

func (s *Service) GetEmployee(ctx context.Context, id string) (models.Employee, error) {
	emp, err := s.Store.GetEmployee(ctx, id)
	if err != nil {
		return models.Employee{}, err
	}
	return emp, nil
}

func (s *Service) CreateEmployee(ctx context.Context, emp models.Employee) error {
	err := s.Store.CreateEmployee(ctx, emp)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateEmployee(ctx context.Context, emp models.Employee) error {
	err := s.Store.UpdateEmployee(ctx, emp)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteEmployee(ctx context.Context, id string) error {
	err := s.Store.DeleteEmployee(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
