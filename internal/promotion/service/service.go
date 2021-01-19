package service

import (
	"context"

	promotionError "github.com/fadhilnurmahardi/simple-promotion/internal/promotion/error"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy"
)

type Service struct {
	discountStrategy []strategy.IStrategy
}

func New(discountStrategy []strategy.IStrategy) *Service {
	return &Service{
		discountStrategy: discountStrategy,
	}
}

func (s *Service) Calculate(ctx context.Context, payload []model.Payload) (*model.Result, error) {
	realTotal := 0.0
	for i := 0; i < len(payload); i++ {
		realTotal += payload[0].Price * float64(payload[0].Qty)
	}
	totalDiscount := 0.0
	for i := 0; i < len(s.discountStrategy); i++ {
		discount, err := s.discountStrategy[0].CalculatePromo(ctx, payload)
		if err != nil && err != promotionError.NotEligible {
			return nil, err
		}
		if discount != nil {
			totalDiscount += *discount
		}
	}
	return &model.Result{
		Total:          realTotal,
		Discount:       totalDiscount,
		TotalAfterDisc: realTotal - totalDiscount,
	}, nil
}
