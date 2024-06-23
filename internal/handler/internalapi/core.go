package internalapi

import (
	"context"
	"fmt"
	"net/http"

	xerrors "github.com/syth0le/gopnik/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/syth0le/dialog-service/internal/model"
	"github.com/syth0le/dialog-service/internal/service/dialog"
	inpb "github.com/syth0le/dialog-service/proto/internalapi"
)

type DialogHandler struct {
	inpb.UnimplementedDialogServiceServer

	DialogService dialog.Service
}

func (h *DialogHandler) CreateDialog(ctx context.Context, request *inpb.CreateDialogRequest) (*inpb.CreateDialogResponse, error) {
	participantsIDs := make([]model.UserID, len(request.ParticipantsIds))
	for idx, item := range request.ParticipantsIds {
		participantsIDs[idx] = model.UserID(item)
	}

	dialogModel, err := h.DialogService.CreateDialog(ctx, &dialog.CreateDialogParams{
		ParticipantsIDs: participantsIDs,
	})
	if err != nil {
		return nil, GRPCError(fmt.Errorf("get dialog: %w", err))
	}

	return &inpb.CreateDialogResponse{
		DialogId:        dialogModel.ID.String(),
		ParticipantsIds: request.ParticipantsIds,
	}, nil
}

func (h *DialogHandler) CreateMessage(ctx context.Context, request *inpb.CreateMessageRequest) (*emptypb.Empty, error) {
	err := h.DialogService.CreateMessage(ctx, &dialog.CreateMessageParams{
		DialogID: model.DialogID(request.DialogId),
		SenderID: model.UserID(request.SenderId),
		Text:     request.Text,
	})
	if err != nil {
		return nil, GRPCError(fmt.Errorf("get dialog: %w", err))
	}

	return &emptypb.Empty{}, nil
}

func (h *DialogHandler) GetDialogMessages(ctx context.Context, request *inpb.GetDialogMessagesRequest) (*inpb.GetDialogMessagesResponse, error) {
	messages, err := h.DialogService.GetDialogMessages(ctx, &dialog.GetDialogMessagesParams{
		UserID:   model.UserID(request.UserId),
		DialogID: model.DialogID(request.DialogId),
	})
	if err != nil {
		return nil, GRPCError(fmt.Errorf("get dialog: %w", err))
	}

	pbMessages := make([]*inpb.Message, len(messages))
	for idx, item := range messages {
		pbMessages[idx] = &inpb.Message{
			Id:       item.ID.String(),
			DialogId: item.DialogID.String(),
			SenderId: item.SenderID.String(),
			Text:     item.Text,
		}
	}

	return &inpb.GetDialogMessagesResponse{Messages: pbMessages}, nil
}

// GRPCError todo: move to gopnik and create server interceptor
func GRPCError(err error) error {
	resError, ok := xerrors.FromError(err)
	if !ok {
		return err
	}

	switch resError.StatusCode {
	case http.StatusForbidden:
		return status.Error(codes.PermissionDenied, resError.Msg)
	case http.StatusNotFound:
		return status.Error(codes.NotFound, resError.Msg)
	case http.StatusBadRequest:
		return status.Error(codes.InvalidArgument, resError.Msg)
	default:
		return status.Errorf(codes.Internal, resError.Msg)
	}
}
