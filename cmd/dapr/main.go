package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/sah4ez/dapr-example/internal/config"
	"github.com/sah4ez/dapr-example/internal/repository"
	"github.com/sah4ez/dapr-example/internal/repository/postgres"
	"github.com/sah4ez/dapr-example/internal/services/balance"
	"github.com/sah4ez/dapr-example/internal/services/user"
	"github.com/sah4ez/dapr-example/pkg/errors"

	daprd "github.com/dapr/go-sdk/service/http"
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

	s := daprd.NewService(config.Service().Bind)

	var storeUser repository.User
	storeUser = &postgres.Mock{}

	var ptrUser *user.Service
	{
		ptrUser = user.New(storeUser)
		err := s.AddServiceInvocationHandler("/api/v1/user", ptrUser.GetNameByIDHandler)
		errors.ExitOnError(log.Logger, err, "add user handler")
	}

	var ptrBalance *balance.Service
	{
		ptrBalance = balance.New(ptrUser)
		err := s.AddServiceInvocationHandler("/api/v1/balance/get/balance", ptrBalance.GetBalanceHandler)
		errors.ExitOnError(log.Logger, err, "add balance handler")
	}

	go func() {
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			errors.ExitOnError(log.Logger, err, "server failed")
		}
	}()

	<-shutdown
}
