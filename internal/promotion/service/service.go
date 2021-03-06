package service

import (
	"context"

	"github.com/fadhilnurmahardi/simple-promotion/internal/global/helper"
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
		realTotal += (payload[i].Price * float64(payload[i].Qty))
	}
	totalDiscount := 0.0
	for i := 0; i < len(s.discountStrategy); i++ {
		discount, err := s.discountStrategy[i].CalculatePromo(ctx, payload)
		if err != nil && err != promotionError.NotEligible {
			return nil, err
		}
		if discount != nil {
			totalDiscount += *discount
		}
	}
	return &model.Result{
		Total:          helper.TwoDigit(realTotal),
		Discount:       helper.TwoDigit(totalDiscount),
		TotalAfterDisc: helper.TwoDigit(realTotal - totalDiscount),
	}, nil
}
