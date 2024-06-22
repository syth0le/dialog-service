package dialog

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/syth0le/dialog-service/internal/model"
	"github.com/syth0le/dialog-service/internal/storage"
	"github.com/syth0le/dialog-service/internal/utils"
)

type Service interface {
	CreateDialog(ctx context.Context, params *CreateDialogParams) (*model.Dialog, error)
	CreateMessage(ctx context.Context, params *CreateMessageParams) error
	GetDialogMessages(ctx context.Context, params *GetDialogMessagesParams) ([]*model.Message, error)
}

type ServiceImpl struct {
	logger  *zap.Logger
	storage storage.Storage
}

func NewServiceImpl(logger *zap.Logger, storage storage.Storage) *ServiceImpl {
	return &ServiceImpl{
		logger:  logger,
		storage: storage,
	}
}

type CreateDialogParams struct {
	ParticipantsIDs []model.UserID
}

func (s *ServiceImpl) CreateDialog(ctx context.Context, params *CreateDialogParams) (*model.Dialog, error) {
	dialogID := model.DialogID(utils.GenerateDUID())
	// todo: one transaction
	err := s.storage.Dialog().CreateDialog(ctx, dialogID)
	if err != nil {
		return nil, fmt.Errorf("create dialog: %w", err)
	}

	participants := make([]*model.Participant, len(params.ParticipantsIDs))
	for idx, item := range params.ParticipantsIDs {
		participants[idx] = &model.Participant{
			ID:     model.ParticipantID(utils.GeneratePUID()),
			UserID: item,
		}
	}

	err = s.storage.Dialog().AddParticipants(ctx, dialogID, participants)
	if err != nil {
		return nil, fmt.Errorf("add participants: %w", err)
	}

	return &model.Dialog{
		ID:              dialogID,
		ParticipantsIDs: participants,
	}, nil
}

type CreateMessageParams struct {
	DialogID model.DialogID
	SenderID model.UserID
	Text     string
}

func (s *ServiceImpl) CreateMessage(ctx context.Context, params *CreateMessageParams) error {
	// todo: check user permissions params.UserID
	err := s.storage.Dialog().CreateMessage(ctx, &model.Message{
		ID:       model.MessageID(utils.GenerateMUID()),
		DialogID: params.DialogID,
		SenderID: params.SenderID,
		Text:     params.Text,
	})
	if err != nil {
		return fmt.Errorf("create message: %w", err)
	}

	return nil
}

type GetDialogMessagesParams struct {
	UserID   model.UserID
	DialogID model.DialogID
}

func (s *ServiceImpl) GetDialogMessages(ctx context.Context, params *GetDialogMessagesParams) ([]*model.Message, error) {
	// todo: check user permissions params.UserID
	messages, err := s.storage.Dialog().GetDialogMessages(ctx, params.DialogID)
	if err != nil {
		return nil, fmt.Errorf("get dialog messages: %w", err)
	}

	return messages, nil
}
