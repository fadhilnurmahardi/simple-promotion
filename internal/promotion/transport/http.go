package transport

import (
	"log"
	"net/http"

	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/endpoint"
	"github.com/fadhilnurmahardi/simple-promotion/internal/promotion/service"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeSend ...
func MakeSend(svc service.IService, logger log.Logger, opts []kithttp.ServerOption) http.Handler {
	end := endpoint.Calculate(svc)

}
