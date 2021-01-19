package main

import (
	"net/http"
	"os"

	"github.com/fadhilnurmahardi/simple-promotion/cmd/containerService"
	"github.com/fadhilnurmahardi/simple-promotion/cmd/routes"
	"github.com/fadhilnurmahardi/simple-promotion/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
)

func main() {
	configData := config.GetConfig()
	container := containerService.New()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var g group.Group
	{
		httpLogger := log.With(logger, "component", "http")
		router := chi.NewRouter()
		corsHandler := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Client-ID", "Client-Secret"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		})
		router.Use(corsHandler.Handler)
		router.Mount("/v1", routes.MakeHandler(container, httpLogger))

		g.Add(func() error {
			_ = logger.Log("transport", "debug/HTTP", "addr", configData.HTTPAddress)
			return http.ListenAndServe(configData.HTTPAddress, router)
		}, func(err error) {
			if nil != err {
				_ = logger.Log("transport", "debug/HTTP", "addr", configData.HTTPAddress, "error", err)
				panic(err)
			}
		})
	}
	logger.Log("exit", g.Run())
}
