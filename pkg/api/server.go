package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	configuration "github.com/hyperversalblocks/mailer/pkg/config"
	"github.com/hyperversalblocks/mailer/pkg/logger"
	"github.com/hyperversalblocks/mailer/pkg/mailer"
	"github.com/hyperversalblocks/mailer/pkg/store"
)

type Services struct {
	Config configuration.Config
	logger *logrus.Logger
	Api    *Api
}

func Init() error {
	services, err := bootstrapper(context.Background())
	if err != nil {
		return err
	}

	services.Api.Cors()
	services.Api.Routes()

	go func() {
		services.startServer()
	}()
	select {}
}

func (c *Services) startServer() {
	serverConf := c.Config.GetConfig().Server
	address := serverConf.Host + serverConf.PORT

	c.logger.Info("Starting Server at:", address)

	err := http.ListenAndServe(address, c.Api.GetRouter())
	if err != nil {
		c.logger.Error("error starting server at ", address, " with error: ", err)
		panic(err)
	}
}

func bootstrapper(ctx context.Context) (*Services, error) {
	conf, err := configuration.Init()
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping config: %w", err)
	}

	confInstance := conf.GetConfig()

	loggerInstance := logger.Init(confInstance.Logger.Level, confInstance.Logger.Env)
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping logger: %w", err)
	}

	storer, err := store.New()
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping store: %w", err)
	}

	mailerService := mailer.New(storer, loggerInstance)

	mailController := NewMailerController(loggerInstance, mailerService)

	apiService := InitAPI(loggerInstance,
		mailController)

	return &Services{
		Config: confInstance,
		logger: loggerInstance,
		Api:    apiService,
	}, nil
}

func InitAPI(
	loggerInstance *logrus.Logger,
	mailerService Mailer,
) *Api {
	return New(loggerInstance, chi.NewMux(), mailerService)
}
