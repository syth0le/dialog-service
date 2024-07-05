package sharder

import (
	"context"
	"fmt"

	xerrors "github.com/syth0le/gopnik/errors"

	"github.com/syth0le/dialog-service/internal/model"
	"github.com/syth0le/dialog-service/internal/storage"
)

func (c *ConsistentSharder) CreateDialog(ctx context.Context, id model.DialogID) error {
	shard, err := c.Get(id.String())
	if err != nil {
		return xerrors.WrapInternalError(fmt.Errorf("cannot get shard: %w", err))
	}

	return shard.Dialog().CreateDialog(ctx, id)
}

func (c *ConsistentSharder) GetDialogParticipants(ctx context.Context, dialogID model.DialogID) ([]*model.Participant, error) {
	shard, err := c.Get(dialogID.String())
	if err != nil {
		return nil, xerrors.WrapInternalError(fmt.Errorf("cannot get shard: %w", err))
	}

	return shard.Dialog().GetDialogParticipants(ctx, dialogID)
}

func (c *ConsistentSharder) AddParticipants(ctx context.Context, dialogID model.DialogID, participants []*model.Participant) error {
	shard, err := c.Get(dialogID.String())
	if err != nil {
		return xerrors.WrapInternalError(fmt.Errorf("cannot get shard: %w", err))
	}

	return shard.Dialog().AddParticipants(ctx, dialogID, participants)
}

func (c *ConsistentSharder) GetDialogMessages(ctx context.Context, dialogID model.DialogID) ([]*model.Message, error) {
	shard, err := c.Get(dialogID.String())
	if err != nil {
		return nil, xerrors.WrapInternalError(fmt.Errorf("cannot get shard: %w", err))
	}

	return shard.Dialog().GetDialogMessages(ctx, dialogID)
}

func (c *ConsistentSharder) CreateMessage(ctx context.Context, params *model.Message) (*model.Message, error) {
	shard, err := c.Get(params.DialogID.String())
	if err != nil {
		return nil, xerrors.WrapInternalError(fmt.Errorf("cannot get shard: %w", err))
	}

	return shard.Dialog().CreateMessage(ctx, params)
}

func (c *ConsistentSharder) DeleteParticipants(ctx context.Context, dialogID model.DialogID, participants []*model.Participant) error {
	shard, err := c.Get(dialogID.String())
	if err != nil {
		return xerrors.WrapInternalError(fmt.Errorf("cannot get shard: %w", err))
	}

	return shard.Dialog().DeleteParticipants(ctx, dialogID, participants)
}

func (c *ConsistentSharder) DeleteDialog(ctx context.Context, dialogID model.DialogID) error {
	shard, err := c.Get(dialogID.String())
	if err != nil {
		return xerrors.WrapInternalError(fmt.Errorf("cannot get shard: %w", err))
	}

	return shard.Dialog().DeleteDialog(ctx, dialogID)
}

func (c *ConsistentSharder) DeleteMessage(ctx context.Context, dialogID model.DialogID, messageID model.MessageID) error {
	shard, err := c.Get(dialogID.String())
	if err != nil {
		return xerrors.WrapInternalError(fmt.Errorf("cannot get shard: %w", err))
	}

	return shard.Dialog().DeleteMessage(ctx, dialogID, messageID)
}

func (c *ConsistentSharder) Dialog() storage.DialogRepository {
	return c
}

func (c *ConsistentSharder) Salt() string {
	return ""
}
