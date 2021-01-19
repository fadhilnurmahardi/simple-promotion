package strategy

import (
	"context"

	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
)

type IStrategy interface {
	CalculatePromo(ctx context.Context, items []model.Payload) (*float64, error)
}
