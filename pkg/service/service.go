package service

import "ecurrency/pkg/repository"

type Currency interface {
	GetRate() (float32, error)
}

type Email interface {
}

type Service struct {
	Currency
	Email
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
