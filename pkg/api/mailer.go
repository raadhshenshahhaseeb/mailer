package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/mailer/pkg/mailer"
)

type email struct {
	Mail string `json:"mail"`
}

type mailController struct {
	logger  *logrus.Logger
	service mailer.Service
}

func (m *mailController) Add(w http.ResponseWriter, r *http.Request) {
	mail := new(email)
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		WriteJson(w, &outputDTO{
			Message:   "internal server error",
			Data:      nil,
			Timestamp: time.Now().String(),
		}, http.StatusInternalServerError)
		return
	}

	err = m.service.Insert(mail.Mail)
	if err != nil {
		WriteJson(w, &outputDTO{
			Message:   "internal server error",
			Data:      nil,
			Timestamp: time.Now().String(),
		}, http.StatusInternalServerError)
		return
	}

	WriteJson(w, &outputDTO{
		Message:   "success",
		Data:      nil,
		Timestamp: time.Now().String(),
	}, http.StatusOK)
}

func NewMailerController(logger *logrus.Logger, service mailer.Service) Mailer {
	return &mailController{
		logger:  logger,
		service: service,
	}
}

type Mailer interface {
	Add(w http.ResponseWriter, r *http.Request)
}
