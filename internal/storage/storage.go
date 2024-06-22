package storage

import (
	"context"

	"github.com/syth0le/dialog-service/internal/model"
)

type Storage interface {
	Dialog() DialogRepository
}

type DialogRepository interface {
	CreateDialog(ctx context.Context, id model.DialogID) error
	GetDialogParticipants(ctx context.Context, dialogID model.DialogID) ([]*model.Participant, error)
	AddParticipants(ctx context.Context, dialogID model.DialogID, participants []*model.Participant) error
	GetDialogMessages(ctx context.Context, dialogID model.DialogID) ([]*model.Message, error) // TODO: pagination
	CreateMessage(ctx context.Context, params *model.Message) error
}
