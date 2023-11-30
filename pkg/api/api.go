package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Api struct {
	logger         *logrus.Logger
	router         *chi.Mux
	mailController Mailer
}

func New(logger *logrus.Logger, router *chi.Mux, mailer Mailer) *Api {
	return &Api{
		logger:         logger,
		router:         router,
		mailController: mailer,
	}
}
