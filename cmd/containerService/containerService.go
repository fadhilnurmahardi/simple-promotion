package containerService

import (
	"strconv"
	"strings"

	"github.com/fadhilnurmahardi/simple-promotion/config"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/service"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy/bonusItem"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy/payTwoForThree"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy/percentage"
)

type Container struct {
	PromotionService service.IService
}

func New() *Container {
	configData := config.GetConfig()
	bonusItemMinBuy, err := strconv.Atoi(configData.BonusItemMinBuy)
	if err != nil {
		panic(err)
	}

	payTwoForThreeMinBuy, err := strconv.Atoi(configData.PayTwoForThreeMinBuy)
	if err != nil {
		panic(err)
	}

	percentageMinBuy, err := strconv.Atoi(configData.PercentageMinBuy)
	if err != nil {
		panic(err)
	}

	percentageValue, err := strconv.Atoi(configData.Percentage)
	if err != nil {
		panic(err)
	}

	var discountStrategies = make([]strategy.IStrategy, 0)
	bonusItemStrategy := bonusItem.New(configData.BonusItemEligible, configData.BonusItemSKU, bonusItemMinBuy)
	discountStrategies = append(discountStrategies, bonusItemStrategy)
	payTwoForThreeStrategy := payTwoForThree.New(strings.Split(configData.PayTwoForThreeEligible, ","), payTwoForThreeMinBuy)
	discountStrategies = append(discountStrategies, payTwoForThreeStrategy)
	percentageStrategy := percentage.New(strings.Split(configData.PercentageEligible, ","), percentageMinBuy, percentageValue)
	discountStrategies = append(discountStrategies, percentageStrategy)

	return &Container{
		PromotionService: service.New(discountStrategies),
	}
}
