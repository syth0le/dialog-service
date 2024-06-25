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

func (s *Storage) CreateDialog(ctx context.Context, id model.DialogID) error {
	now := time.Now().Truncate(time.Millisecond)

	sql, args, err := sq.Insert(DialogTable).
		Columns(dialogFields...).
		Values(
			id.String(), now,
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

func (s *Storage) GetDialogParticipants(ctx context.Context, dialogID model.DialogID) ([]*model.Participant, error) {
	sql, args, err := sq.Select(participantsFields...).From(ParticipantsTable).
		Where(sq.Eq{
			fieldDialogId:  dialogID.String(),
			fieldDeletedAt: nil,
		}).
		OrderBy(fieldCreatedAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, xerrors.WrapInternalError(fmt.Errorf("incorrect sql"))
	}

	var entities []participantEntity
	err = sqlx.SelectContext(ctx, s.Slave(), &entities, sql, args...)
	if err != nil {
		return nil, xerrors.WrapSqlError(err)
	}

	return participantEntitiesToModels(entities), nil
}

func (s *Storage) AddParticipants(ctx context.Context, dialogID model.DialogID, participants []*model.Participant) error {
	now := time.Now().Truncate(time.Millisecond)

	query := sq.Insert(ParticipantsTable).
		Columns(participantsFields...)

	for _, participant := range participants {
		query = query.Values(participant.ID, dialogID.String(), participant.UserID, now)
	}

	sql, args, err := query.
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

func (s *Storage) GetDialogMessages(ctx context.Context, dialogID model.DialogID) ([]*model.Message, error) {
	sql, args, err := sq.Select(messageFields...).From(MessageTable).
		Where(sq.Eq{
			fieldDialogId:  dialogID.String(),
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

func (s *Storage) CreateMessage(ctx context.Context, params *model.Message) (*model.Message, error) {
	err := params.Validate()
	if err != nil {
		return nil, fmt.Errorf("params validate: %w", err)
	}

	now := time.Now().Truncate(time.Millisecond)

	sql, args, err := sq.Insert(MessageTable).
		Columns(messageFields...).
		Values(
			params.ID.String(), params.DialogID.String(), params.SenderID, params.Text, now, now,
		).
		PlaceholderFormat(sq.Dollar).
		Suffix(returningMessage).
		ToSql()
	if err != nil {
		return nil, xerrors.WrapInternalError(fmt.Errorf("incorrect sql"))
	}

	var entity messageEntity
	err = sqlx.GetContext(ctx, s.Master(), &entity, sql, args...)
	if err != nil {
		return nil, xerrors.WrapSqlError(err)
	}

	return messageEntityToModel(entity), nil
}

func (s *Storage) DeleteParticipants(ctx context.Context, participants []*model.Participant) error {
	ids := make([]string, len(participants))
	for idx, val := range participants {
		ids[idx] = val.ID.String()
	}

	sql, args, err := sq.Delete(ParticipantsTable).
		Where(
			sq.Eq{fieldID: ids},
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

func (s *Storage) DeleteDialog(ctx context.Context, dialogID model.DialogID) error {
	sql, args, err := sq.Delete(DialogTable).
		Where(
			sq.Eq{fieldID: dialogID.String()},
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

func (s *Storage) DeleteMessage(ctx context.Context, messageID model.MessageID) error {
	sql, args, err := sq.Delete(MessageTable).
		Where(
			sq.Eq{fieldID: messageID.String()},
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
		ID:       model.MessageID(entity.ID),
		DialogID: model.DialogID(entity.DialogID),
		SenderID: model.UserID(entity.SenderID),
		Text:     entity.Text,
	}
}

func messageEntitiesToModels(entities []messageEntity) []*model.Message {
	var friendModels []*model.Message
	for _, entity := range entities {
		friendModels = append(friendModels, messageEntityToModel(entity))
	}
	return friendModels
}

type participantEntity struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`
}

func participantEntityToModel(entity participantEntity) *model.Participant {
	return &model.Participant{
		ID:     model.ParticipantID(entity.ID),
		UserID: model.UserID(entity.UserID),
	}
}

func participantEntitiesToModels(entities []participantEntity) []*model.Participant {
	var friendModels []*model.Participant
	for _, entity := range entities {
		friendModels = append(friendModels, participantEntityToModel(entity))
	}
	return friendModels
}
