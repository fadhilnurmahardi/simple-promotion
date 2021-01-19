package endpoint

import (
	"context"

	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/service"
	"github.com/go-kit/kit/endpoint"
)

func Calculate(s service.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		reqData := request.([]model.Payload)
		return s.Calculate(ctx, reqData)
	}
}
