package repositories

import (
	"context"
	"zatrano/configs/databaseconfig"
	"zatrano/models"
)

type IPaymentRepository interface {
	GetPaymentByID(id uint) (*models.Payment, error)
	CreatePayment(ctx context.Context, payment *models.Payment) error
	DeletePayment(ctx context.Context, id uint) error
	GetPaymentCount() (int64, error)
}

type PaymentRepository struct {
	base IBaseRepository[models.Payment]
}

func NewPaymentRepository() IPaymentRepository {
	base := NewBaseRepository[models.Payment](databaseconfig.GetDB())
	base.SetAllowedSortColumns([]string{"id", "amount", "installment", "status", "created_at"})
	return &PaymentRepository{base: base}
}

func (r *PaymentRepository) GetPaymentByID(id uint) (*models.Payment, error) {
	return r.base.GetByID(id)
}

func (r *PaymentRepository) CreatePayment(ctx context.Context, payment *models.Payment) error {
	return r.base.Create(ctx, payment)
}

func (r *PaymentRepository) DeletePayment(ctx context.Context, id uint) error {
	return r.base.Delete(ctx, id)
}

func (r *PaymentRepository) GetPaymentCount() (int64, error) {
	return r.base.GetCount()
}

var _ IPaymentRepository = (*PaymentRepository)(nil)
var _ IBaseRepository[models.Payment] = (*BaseRepository[models.Payment])(nil)
