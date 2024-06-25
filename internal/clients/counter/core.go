package counter

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	inpb "github.com/syth0le/counter-service/proto/internalapi"

	"github.com/syth0le/dialog-service/internal/model"
)

type Client interface {
	CreateDialogCounters(ctx context.Context, dialogID model.DialogID, participants []*model.Participant) error
	IncreaseDialogCounters(ctx context.Context, dialogID model.DialogID, senderID model.UserID, participants []*model.Participant) error
	FlushDialogCountersForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) error
}

type ClientImpl struct {
	logger *zap.Logger
	client inpb.CounterServiceClient
}

func NewClientImpl(logger *zap.Logger, conn *grpc.ClientConn) *ClientImpl {
	return &ClientImpl{
		logger: logger,
		client: inpb.NewCounterServiceClient(conn),
	}
}

func (c *ClientImpl) CreateDialogCounters(ctx context.Context, dialogID model.DialogID, participants []*model.Participant) error {
	participantsIDs := make([]string, len(participants))
	for idx, item := range participants {
		participantsIDs[idx] = item.UserID.String()
	}

	_, err := c.client.CreateDialogCounters(ctx, &inpb.CreateDialogCountersRequest{
		DialogId:     dialogID.String(),
		Participants: participantsIDs,
	})
	if err != nil {
		return fmt.Errorf("create dialog counters: %w", err)
	}

	return nil
}

func (c *ClientImpl) IncreaseDialogCounters(ctx context.Context, dialogID model.DialogID, senderID model.UserID, participants []*model.Participant) error {
	participantsIDs := make([]string, len(participants))
	for idx, item := range participants {
		participantsIDs[idx] = item.UserID.String()
	}

	_, err := c.client.IncreaseDialogCounters(ctx, &inpb.IncreaseDialogCountersRequest{
		DialogId:    dialogID.String(),
		SenderId:    senderID.String(),
		FollowersId: participantsIDs,
	})
	if err != nil {
		return fmt.Errorf("increase dialog counters: %w", err)
	}

	return nil
}

func (c *ClientImpl) FlushDialogCountersForUser(ctx context.Context, dialogID model.DialogID, userID model.UserID) error {
	_, err := c.client.FlushDialogCountersForUser(ctx, &inpb.FlushDialogCountersForUserRequest{
		DialogId: dialogID.String(),
		UserId:   userID.String(),
	})
	if err != nil {
		return fmt.Errorf("flusf dialog counters for user: %w", err)
	}

	return nil
}
