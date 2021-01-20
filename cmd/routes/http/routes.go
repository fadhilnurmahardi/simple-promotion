package routes

import (
	"net/http"

	"github.com/fadhilnurmahardi/simple-promotion/cmd/containerService"
	"github.com/fadhilnurmahardi/simple-promotion/internal/global/helper"
	httpTransport "github.com/fadhilnurmahardi/simple-promotion/internal/promotion/transport/http"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(container *containerService.Container, logger log.Logger) http.Handler {
	router := chi.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(helper.EncodeError),
	}

	router.Post("/promotion/check", httpTransport.MakeSend(container.PromotionService, logger, opts).ServeHTTP)

	return router
}
