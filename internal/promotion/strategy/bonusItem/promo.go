package bonusItem

import (
	"context"
	"fmt"

	promotionError "github.com/fadhilnurmahardi/simple-promotion/internal/promotion/error"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
)

type Promo struct {
	eligibleSKU string
	bonusSKU    string
	minimalBuy  int
}

func New(eligibleSKU string, bonusSKU string, minimalBuy int) *Promo {
	return &Promo{
		eligibleSKU: eligibleSKU,
		minimalBuy:  minimalBuy,
		bonusSKU:    bonusSKU,
	}
}

func (p *Promo) CalculatePromo(ctx context.Context, items []model.Payload) (*float64, error) {
	foundSKU := 0
	for q := 0; q < len(items); q++ {
		if p.eligibleSKU == items[q].SKU {
			if items[q].Qty >= p.minimalBuy {
				foundSKU = items[q].Qty
				break
			}
		}
	}
	if foundSKU > 0 {
		for q := 0; q < len(items); q++ {
			fmt.Println(p.bonusSKU, items[q].SKU)
			if p.bonusSKU == items[q].SKU {
				if items[q].Qty >= foundSKU {
					discount := float64(foundSKU) * items[q].Price
					return &discount, nil
				}
			}
		}
	}

	return nil, promotionError.NotEligible
}
