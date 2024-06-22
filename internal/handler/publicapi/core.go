package publicapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	xerrors "github.com/syth0le/gopnik/errors"
	"go.uber.org/zap"

	"github.com/go-http-utils/headers"

	"github.com/syth0le/dialog-service/internal/service/dialog"
)

type Handler struct {
	Logger        *zap.Logger
	DialogService dialog.Service
}

func writeResponse(w http.ResponseWriter, response any) {
	w.Header().Set(headers.ContentType, "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, xerrors.InternalErrorMessage, http.StatusInternalServerError) // TODO: make error mapping
	}
}

func (h *Handler) writeError(ctx context.Context, w http.ResponseWriter, err error) {
	h.Logger.Sugar().Warnf("http response error: %v", err)

	w.Header().Set(headers.ContentType, "application/json")
	errorResult, ok := xerrors.FromError(err)
	if !ok {
		h.Logger.Sugar().Errorf("cannot write log message: %v", err)
		return
	}
	w.WriteHeader(errorResult.StatusCode)
	err = json.NewEncoder(w).Encode(
		map[string]any{
			"message": errorResult.Msg,
			"code":    errorResult.StatusCode,
		})

	if err != nil {
		http.Error(w, xerrors.InternalErrorMessage, http.StatusInternalServerError) // TODO: make error mapping
	}
}

func parseJSONRequest[T createMessageRequest | createDialogRequest](r *http.Request) (*T, error) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("read body: %w", err)
		return nil, xerrors.WrapInternalError(err)
	}

	var request T
	err = json.Unmarshal(body, &request)
	if err != nil {
		err = fmt.Errorf("unmarshal request body: %w", err)
		return nil, xerrors.WrapValidationError(err)
	}
	return &request, nil
}
