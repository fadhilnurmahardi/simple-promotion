package service

import (
	"context"

	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
)

type IService interface {
	Calculate(ctx context.Context, payload []model.Payload) (*model.Result, error)
}
