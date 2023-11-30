package mailer

import (
	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/mailer/pkg/store"
)

type mailer struct {
	store  store.Store
	logger *logrus.Logger
}

func (m *mailer) Insert(email string) error {
	return m.store.InsertRecord(email)
}

type Service interface {
	Insert(email string) error
}

func New(store store.Store, logger *logrus.Logger) Service {
	return &mailer{
		store:  store,
		logger: logger,
	}
}
