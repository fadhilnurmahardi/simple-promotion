package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fadhilnurmahardi/simple-promotion/internal/global/model"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

//EncodeError ...
func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	code := http.StatusInternalServerError
	if sc, ok := err.(*model.TransportError); ok {
		code = sc.Code
	}

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func NewHttpError(code int, err error) *model.TransportError {
	return &model.TransportError{
		Code: code,
		Err:  err,
	}
}

func LogRequest(logger log.Logger) endpoint.Middleware {
	return func(f endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				kv := make([]interface{}, 0)

				jsonString, _ := json.Marshal(request)
				jsonResp, _ := json.Marshal(response)
				kv = append(kv,
					"params", string(jsonString),
					"result", string(jsonResp),
					"took", time.Since(begin).String(),
				)

				logger.Log(kv...)
			}(time.Now())
			return f(ctx, request)
		}
	}
}

func TwoDigit(amount float64) float64 {
	amount, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", amount), 64)
	return amount
}
