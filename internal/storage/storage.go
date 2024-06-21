package storage

import (
	"context"

	"github.com/syth0le/dialog-service/internal/model"
)

type Storage interface {
	Dialog() DialogRepository
}

type DialogRepository interface {
	GetDialog(ctx context.Context, dialogID *model.DialogID) ([]*model.Message, error) // TODO: pagination
	CreateMessage(ctx context.Context, msg *model.Message) error
}
