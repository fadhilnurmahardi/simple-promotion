package percentage

import (
	"context"

	promotionError "github.com/fadhilnurmahardi/simple-promotion/internal/promotion/error"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
)

type Promo struct {
	eligibleSKU []string
	minimalBuy  int
	percentage  float64
}

func New(eligibleSKU []string, minimalBuy int, percentage int) *Promo {
	percent := float64(percentage) / 100
	return &Promo{
		eligibleSKU: eligibleSKU,
		minimalBuy:  minimalBuy,
		percentage:  percent,
	}
}

func (p *Promo) CalculatePromo(ctx context.Context, items []model.Payload) (*float64, error) {
	for i := 0; i < len(p.eligibleSKU); i++ {
		for q := 0; q < len(items); q++ {
			if p.eligibleSKU[i] == items[q].SKU {
				if items[q].Qty >= p.minimalBuy {
					discount := (items[q].Price * float64(items[q].Qty) * p.percentage)
					return &discount, nil
				}
			}
		}
	}

	return nil, promotionError.NotEligible
}
