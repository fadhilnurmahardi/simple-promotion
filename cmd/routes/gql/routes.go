package gql

import (
	"net/http"

	"github.com/fadhilnurmahardi/simple-promotion/cmd/containerService"
	promotionGQL "github.com/fadhilnurmahardi/simple-promotion/internal/promotion/transport/gql"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
)

func MakeHandler(container *containerService.Container, logger log.Logger) http.Handler {
	router := chi.NewRouter()

	router.HandleFunc("/promotion", promotionGQL.MakeHandler(container.PromotionService, logger))

	return router
}
