package mailer

import (
	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/mailer/pkg/store"
)

type mailer struct {
	store  store.Store
	logger *logrus.Logger
}

func (m *mailer) Get() (*[]store.Emails, error) {
	return m.store.Get()
}

func (m *mailer) Insert(email string) error {
	return m.store.InsertRecord(email)
}

type Service interface {
	Insert(email string) error
	Get() (*[]store.Emails, error)
}

func New(store store.Store, logger *logrus.Logger) Service {
	return &mailer{
		store:  store,
		logger: logger,
	}
}
