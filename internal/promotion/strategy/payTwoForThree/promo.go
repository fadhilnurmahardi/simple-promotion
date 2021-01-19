package payTwoForThree

import (
	"context"
	"math"

	promotionError "github.com/fadhilnurmahardi/simple-promotion/internal/promotion/error"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
)

type Promo struct {
	eligibleSKU []string
	minimalBuy  int
}

func New(eligibleSKU []string, minimalBuy int, percentage int) *Promo {
	return &Promo{
		eligibleSKU: eligibleSKU,
		minimalBuy:  minimalBuy,
	}
}

func (p *Promo) CalculatePromo(ctx context.Context, items []model.Payload) (*float64, error) {
	for i := 0; i < len(p.eligibleSKU); i++ {
		for q := 0; q < len(items); q++ {
			if p.eligibleSKU[i] == items[q].SKU {
				if items[q].Qty >= p.minimalBuy {
					qty := items[q].Qty
					price := items[q].Price
					afterDisc := math.Floor(float64(qty/p.minimalBuy)*(price*2)) + (float64(qty%p.minimalBuy) * price)
					realTotal := float64(qty) * price
					discount := realTotal - afterDisc
					return &discount, nil
				}
			}
		}
	}

	return nil, promotionError.NotEligible
}
