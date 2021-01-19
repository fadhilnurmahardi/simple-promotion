package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fadhilnurmahardi/simple-promotion/internal/global/helper"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/endpoint"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/model"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/service"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeSend ...
func MakeSend(svc service.IService, logger log.Logger, opts []kithttp.ServerOption) http.Handler {
	end := endpoint.Calculate(svc)

	end = helper.LogRequest(logger)(end)

	return kithttp.NewServer(
		end,
		func(ctx context.Context, r *http.Request) (request interface{}, err error) {
			var body *model.RawPayload
			err = json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				return nil, helper.NewHttpError(412, err)
			}
			return body.Cart, nil
		},
		func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			if e, ok := response.(error); ok && e.Error() != "" {
				return e
			}

			if response == nil {
				w.WriteHeader(http.StatusNoContent)
				json.NewEncoder(w).Encode("")
				return nil
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(response)
			return nil
		},
		opts...,
	)
}
