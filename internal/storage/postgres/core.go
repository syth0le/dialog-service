package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	xstorage "github.com/syth0le/gopnik/db/postgres"
	"go.uber.org/zap"

	"github.com/syth0le/dialog-service/internal/storage"
)

type Storage struct {
	storage *xstorage.PGStorage
	hosts   []string
	salt    string
}

func NewStorage(logger *zap.Logger, config xstorage.StorageConfig, salt string) (*Storage, error) {
	postgresStorage, err := xstorage.NewPGStorage(logger, config)
	if err != nil {
		return nil, fmt.Errorf("new pg storage: %w", err)
	}
	return &Storage{
		storage: postgresStorage,
		hosts:   config.Hosts,
		salt:    salt,
	}, nil
}

func (s *Storage) Dialog() storage.DialogRepository {
	return s
}

func (s *Storage) Close() error {
	return s.storage.Close()
}

func (s *Storage) Master() sqlx.ExtContext {
	return s.storage.Master()
}

func (s *Storage) Slave() sqlx.ExtContext {
	return s.storage.Slave()
}

func (s *Storage) now() {
	// TODO: implement
}

func (s *Storage) Hosts() []string {
	return s.hosts
}

func (s *Storage) Salt() string {
	return s.salt
}
