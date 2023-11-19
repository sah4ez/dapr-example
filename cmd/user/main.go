package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"

	"github.com/sah4ez/dapr-example/interfaces"
	"github.com/sah4ez/dapr-example/internal/config"
	"github.com/sah4ez/dapr-example/internal/repository"
	"github.com/sah4ez/dapr-example/internal/repository/postgres"
	"github.com/sah4ez/dapr-example/internal/services/balance"
	"github.com/sah4ez/dapr-example/internal/services/user"
	"github.com/sah4ez/dapr-example/internal/transport"
	"github.com/sah4ez/dapr-example/pkg/utils/header"
)

const (
	serviceName = "userBalance"
)

func main() {

	log.Logger = config.Service().Logger()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	log.Info().Str("service", serviceName).Msg("start service")
	defer log.Info().Msg("shutdown server")

	var storeUser repository.User
	storeUser = &postgres.Mock{}

	var ptrUser *user.Service
	var svcUser interfaces.User
	{
		ptrUser = user.New(storeUser)
		svcUser = ptrUser
	}

	var ptrBalance *balance.Service
	var svcBalance interfaces.Balance
	{
		ptrBalance = balance.New(ptrUser)
		svcBalance = ptrBalance
	}

	services := []transport.Option{
		transport.Use(cors.New(cors.Config{AllowHeaders: config.Service().AllowHeaders})),
		transport.WithRequestID(header.RequestHeader),

		transport.User(transport.NewUser(svcUser)),
		transport.Balance(transport.NewBalance(svcBalance)),
	}

	srv := transport.New(log.Logger, services...).WithMetrics().WithLog()

	transport.ServeMetrics(log.Logger, "/metrics", config.Service().BindMetrics)
	srv.Fiber().Get("/healthcheck", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusOK)
		return ctx.JSON(map[string]string{"status": "OK"})
	})

	go func() {
		log.Info().Str("bind", config.Service().Bind).Msg("listen on")
		if err := srv.Fiber().Listen(config.Service().Bind); err != nil {
			log.Panic().Err(err).Stack().Msg("server error")
		}
	}()

	if config.Service().EnabledPPROF {
		go func() {
			err := http.ListenAndServe(config.Service().BindPPROF, nil)
			if err != nil {
				log.Err(err).Str("bind", config.Service().BindPPROF).Msg("start pprof")
			}
		}()
	}

	<-shutdown
}
