package services

import (
	"context"
	"zatrano/models"
	"zatrano/repositories"
)

type IPaymentService interface {
	GetPaymentByID(id uint) (*models.Payment, error)
	CreatePayment(ctx context.Context, payment *models.Payment) error
	DeletePayment(ctx context.Context, id uint) error
	GetPaymentCount() (int64, error)
}

type PaymentService struct {
	repo repositories.IPaymentRepository
}

func NewPaymentService() IPaymentService {
	return &PaymentService{repo: repositories.NewPaymentRepository()}
}

func (s *PaymentService) GetPaymentByID(id uint) (*models.Payment, error) {
	return s.repo.GetPaymentByID(id)
}

func (s *PaymentService) CreatePayment(ctx context.Context, payment *models.Payment) error {
	return s.repo.CreatePayment(ctx, payment)
}

func (s *PaymentService) DeletePayment(ctx context.Context, id uint) error {
	return s.repo.DeletePayment(ctx, id)
}

func (s *PaymentService) GetPaymentCount() (int64, error) {
	return s.repo.GetPaymentCount()
}
