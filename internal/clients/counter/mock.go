package counter

import (
	"context"

	"go.uber.org/zap"

	"github.com/syth0le/dialog-service/internal/model"
)

type ClientMock struct {
	logger *zap.Logger
}

func NewClientMock(logger *zap.Logger) *ClientMock {
	return &ClientMock{
		logger: logger,
	}
}

func (m *ClientMock) CreateDialogCounters(ctx context.Context, dialogID model.DialogID, participants []*model.Participant) error {
	m.logger.Sugar().Debug("create dialog counters through mock service")
	return nil
}

func (m *ClientMock) IncreaseDialogCounters(ctx context.Context, dialogID model.DialogID, senderID model.UserID, participants []*model.Participant) error {
	m.logger.Debug("increase dialog counters through mock service")
	return nil
}

func (m *ClientMock) FlushDialogCountersForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) error {
	m.logger.Debug("flush dialog counters for user through mock service")
	return nil
}
