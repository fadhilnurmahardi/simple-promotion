package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/fadhilnurmahardi/simple-promotion/internal/global/helper"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/service"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy/bonusItem"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy/payTwoForThree"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/strategy/percentage"
)

var svc service.IService

func init() {
	var discountStrategies = make([]strategy.IStrategy, 0)
	bonusItemStrategy := bonusItem.New("43N23P", "234234", 1)
	discountStrategies = append(discountStrategies, bonusItemStrategy)
	payTwoForThreeStrategy := payTwoForThree.New([]string{"120P90"}, 3)
	discountStrategies = append(discountStrategies, payTwoForThreeStrategy)
	percentageStrategy := percentage.New([]string{"A304SD"}, 3, 10)
	discountStrategies = append(discountStrategies, percentageStrategy)
	svc = service.New(discountStrategies)
}

func TestService_Calculate(t *testing.T) {
	type TestCase struct {
		Name        string
		Request     []model.Payload
		ExpectError bool
		Result      *model.Result
	}

	testCases := []TestCase{
		{
			Request: []model.Payload{
				{
					SKU:   "A304SD",
					Price: 500.00,
					Qty:   3,
				},
			},
			Result: &model.Result{
				Total:          helper.TwoDigit(500.00 * 3),
				TotalAfterDisc: helper.TwoDigit((500.00 * 3) - (500.00 * 3 * 0.1)),
				Discount:       helper.TwoDigit(500.00 * 3 * 0.1),
			},
			Name:        "Percent Discount",
			ExpectError: true,
		},
		{
			Request: []model.Payload{
				{
					SKU:   "43N23P",
					Price: 500.00,
					Qty:   1,
				},
				{
					SKU:   "1696969",
					Price: 500.00,
					Qty:   1,
				},
				{
					SKU:   "234234",
					Price: 400.00,
					Qty:   1,
				},
			},
			Result: &model.Result{
				Total:          helper.TwoDigit(500.00 + 400.00 + 500.00),
				TotalAfterDisc: 1000.00,
				Discount:       400.00,
			},
			Name:        "Buy 1 item get 1 other item",
			ExpectError: true,
		},
		{
			Request: []model.Payload{
				{
					SKU:   "120P90",
					Price: 500.00,
					Qty:   3,
				},
			},
			Result: &model.Result{
				Total:          helper.TwoDigit(500.00 * 3),
				TotalAfterDisc: helper.TwoDigit(500.00 * 2),
				Discount:       helper.TwoDigit(500.00 * 1),
			},
			Name:        "Buy 3 item pay 2",
			ExpectError: true,
		},
		{
			Request: []model.Payload{
				{
					SKU:   "120P90",
					Price: 500.00,
					Qty:   2,
				},
			},
			Result: &model.Result{
				Total:          helper.TwoDigit(500.00 * 2),
				TotalAfterDisc: helper.TwoDigit(500.00 * 2),
				Discount:       helper.TwoDigit(500.00 * 0),
			},
			Name:        "no discount",
			ExpectError: true,
		},
		{
			Request: []model.Payload{
				{
					SKU:   "120P90",
					Price: 500.00,
					Qty:   3,
				},
				{
					SKU:   "A304SD",
					Price: 500.00,
					Qty:   3,
				},
				{
					SKU:   "234234",
					Price: 400.00,
					Qty:   1,
				},
				{
					SKU:   "43N23P",
					Price: 500.00,
					Qty:   1,
				},
			},
			Result: &model.Result{
				Total:          helper.TwoDigit((500.00 * 7) + 400.00),
				TotalAfterDisc: helper.TwoDigit(((500.00 * 7) + 400.00) - ((500.00 * 1) + (500.00 * 3 * 0.1) + 400.00)),
				Discount:       helper.TwoDigit((500.00 * 1) + (500.00 * 3 * 0.1) + 400.00),
			},
			Name:        "mix all discount",
			ExpectError: true,
		},
	}

	for _, val := range testCases {
		t.Log(fmt.Sprintf("TestService_Calculate %s", val.Name))
		result, err := svc.Calculate(context.Background(), val.Request)
		if !val.ExpectError && err != nil {
			t.Fatal(err)
		}
		if result.Discount != val.Result.Discount {
			fmt.Println(result.Discount, val.Result.Discount)
			t.Fatal("missmatch result discount")
		}
		if result.Total != val.Result.Total {
			t.Fatal("missmatch result Total")
		}
		if result.TotalAfterDisc != val.Result.TotalAfterDisc {
			t.Fatal("missmatch result TotalAfterDisc")
		}
	}
}
