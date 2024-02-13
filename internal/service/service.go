package service

import "github.com/anil1226/go-employee-dynamo/internal/db"

type Service struct {
	db.Store
}

func NewService(store db.Store) *Service {
	return &Service{
		Store: store,
	}
}
