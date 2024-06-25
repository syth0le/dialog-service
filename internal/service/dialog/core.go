package dialog

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/syth0le/dialog-service/internal/clients/counter"
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
	logger        *zap.Logger
	storage       storage.Storage
	counterClient counter.Client
}

func NewServiceImpl(logger *zap.Logger, storage storage.Storage, client counter.Client) *ServiceImpl {
	return &ServiceImpl{
		logger:        logger,
		storage:       storage,
		counterClient: client,
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

	err = s.counterClient.CreateDialogCounters(ctx, dialogID, participants)
	if err != nil {
		return nil, s.compensationCreateDialogOperation(
			ctx, dialogID, participants, fmt.Errorf("create dialog counters: %w", err),
		)
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
	message, err := s.storage.Dialog().CreateMessage(ctx, &model.Message{
		ID:       model.MessageID(utils.GenerateMUID()),
		DialogID: params.DialogID,
		SenderID: params.SenderID,
		Text:     params.Text,
	})
	if err != nil {
		return fmt.Errorf("create message: %w", err)
	}

	participants, err := s.storage.Dialog().GetDialogParticipants(ctx, params.DialogID)
	if err != nil {
		return s.compensationCreateMessageOperation(ctx, message.ID, fmt.Errorf("get dialog participants: %w", err))
	}

	err = s.counterClient.IncreaseDialogCounters(ctx, params.DialogID, params.SenderID, participants)
	if err != nil {
		return s.compensationCreateMessageOperation(ctx, message.ID, fmt.Errorf("increase dialog counters: %w", err))
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

	err = s.counterClient.FlushDialogCountersForUser(ctx, params.DialogID, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("flush dialog counters for user: %w", err)
	}

	return messages, nil
}

func (s *ServiceImpl) compensationCreateDialogOperation(ctx context.Context, dialogID model.DialogID, participants []*model.Participant, originErr error) error {
	// умышленно отказался от каскадного удаления в sql, потому что это плохая стратегия поведения при наличии пользовательских данных
	// для уверенности можно явно запретить удалять родителя ON DELETE RESTRICT, пока есть потомки
	err := s.storage.Dialog().DeleteParticipants(ctx, participants)
	if err != nil {
		return fmt.Errorf("origin %w: delete participants: %w", originErr, err)
	}

	err = s.storage.Dialog().DeleteDialog(ctx, dialogID)
	if err != nil {
		return fmt.Errorf("origin %w: delete dialog: %w", originErr, err)
	}

	return originErr
}

func (s *ServiceImpl) compensationCreateMessageOperation(ctx context.Context, messageID model.MessageID, originErr error) error {
	err := s.storage.Dialog().DeleteMessage(ctx, messageID)
	if err != nil {
		return fmt.Errorf("origin %w: delete message: %w", originErr, err)
	}

	return originErr
}
