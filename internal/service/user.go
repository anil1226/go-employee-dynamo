package service

import (
	"context"

	"github.com/anil1226/go-employee-dynamo/internal/models"
)

// type UserService struct {
// 	db.UserStore
// }

func (s *Service) GetUser(ctx context.Context, usr models.User) (models.User, error) {
	emp, err := s.Store.GetUser(ctx, usr)
	if err != nil {
		return models.User{}, err
	}
	return emp, nil
}

func (s *Service) CreateUser(ctx context.Context, emp models.User) error {
	err := s.Store.CreateUser(ctx, emp)
	if err != nil {
		return err
	}
	return nil
}
