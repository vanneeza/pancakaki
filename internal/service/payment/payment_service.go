package paymentservice

import (
	"pancakaki/internal/domain/entity"
	paymentrepository "pancakaki/internal/repository/payment"
)

type PaymentService interface {
	InsertPayment(newPayment *entity.Payment) (*entity.Payment, error)
}

type paymentService struct {
	paymentRepo paymentrepository.PaymentRepository
}

// DeletePayment implements PaymentService
func (s *paymentService) InsertPayment(newPayment *entity.Payment) (*entity.Payment, error) {
	return s.paymentRepo.InsertPayment(newPayment)
}
