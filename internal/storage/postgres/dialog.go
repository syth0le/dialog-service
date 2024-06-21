package postgres

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	xerrors "github.com/syth0le/gopnik/errors"

	"github.com/syth0le/dialog-service/internal/model"
)

func (s *Storage) GetDialog(ctx context.Context, dialogID *model.DialogID) ([]*model.Message, error) {
	sql, args, err := sq.Select(messageFields...).From(MessageTable).
		Where(sq.Eq{
			fieldDeletedAt: nil,
		}).
		OrderBy(fieldCreatedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, xerrors.WrapInternalError(fmt.Errorf("incorrect sql"))
	}

	var entities []messageEntity
	err = sqlx.SelectContext(ctx, s.Slave(), &entities, sql, args...)
	if err != nil {
		return nil, xerrors.WrapSqlError(err)
	}

	return messageEntitiesToModels(entities), nil
}

// func (s *Storage) GetDialog(ctx context.Context, dialogID *model.DialogID) ([]*model.Message, error) {
// 	sql, args, err := sq.Select(
// 		tableField(MessageTable, fieldID),
// 		tableField(MessageTable, fieldDialogId),
// 		tableField(MessageTable, fieldSenderId),
// 		tableField(MessageTable, fieldRecipientId),
// 		tableField(MessageTable, fieldText),
// 		tableField(MessageTable, fieldCreatedAt),
// 		tableField(MessageTable, fieldUpdatedAt),
// 	).From(MessageTable).
// 		Join(
// 			joinString(MessageTable, fieldDialogId, DialogTable, fieldID),
// 		).
// 		Where(sq.Eq{
// 			tableField(MessageTable, fieldDeletedAt): nil,
// 		}).
// 		PlaceholderFormat(sq.Dollar).
// 		ToSql()
// 	if err != nil {
// 		return nil, xerrors.WrapInternalError(fmt.Errorf("incorrect sql"))
// 	}
//
// 	var entities []messageEntity
// 	err = sqlx.SelectContext(ctx, s.Slave(), &entities, sql, args...)
// 	if err != nil {
// 		return nil, xerrors.WrapSqlError(err)
// 	}
//
// 	return messageEntitiesToModels(entities), nil
// }

func (s *Storage) CreateMessage(ctx context.Context, params *model.Message) error {
	err := params.Validate()
	if err != nil {
		return fmt.Errorf("params validate: %w", err)
	}

	now := time.Now().Truncate(time.Millisecond)

	sql, args, err := sq.Insert(MessageTable).
		Columns(messageFields...).
		Values(
			params.ID.String(), params.DialogID.String(), params.SenderID, params.RecipientID,
			params.Text, now, now, nil,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return xerrors.WrapInternalError(fmt.Errorf("incorrect sql"))
	}

	_, err = s.Master().ExecContext(ctx, sql, args...)
	if err != nil {
		return xerrors.WrapSqlError(err)
	}

	return nil
}

// todo: create dialog method

type messageEntity struct {
	ID          string `db:"id"`
	DialogID    string `db:"dialog_id"`
	SenderID    string `db:"sender_id"`
	RecipientID string `db:"recipient_id"`
	Text        string `db:"text"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

func messageEntityToModel(entity messageEntity) *model.Message {
	return &model.Message{
		ID: model.MessageID(entity.ID),
	}
}

func messageEntitiesToModels(entities []messageEntity) []*model.Message {
	var friendModels []*model.Message
	for _, entity := range entities {
		friendModels = append(friendModels, messageEntityToModel(entity))
	}
	return friendModels
}
