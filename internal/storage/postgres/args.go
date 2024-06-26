package postgres

import (
	"fmt"
	"strings"
)

const (
	DialogTable       = "dialog_table"
	MessageTable      = "message_table"
	ParticipantsTable = "participant_table"
)

const (
	returning = "RETURNING "
	separator = ","
)

const (
	fieldID = "id"

	fieldDialogId = "dialog_id"
	fieldText     = "text"
	fieldSenderId = "sender_id"

	fieldUserID = "user_id"

	fieldCreatedAt = "created_at"
	fieldUpdatedAt = "updated_at"
	fieldDeletedAt = "deleted_at"
)

var (
	dialogFields       = []string{fieldID, fieldCreatedAt}
	participantsFields = []string{fieldID, fieldDialogId, fieldUserID, fieldCreatedAt}
	messageFields      = []string{fieldID, fieldDialogId, fieldSenderId, fieldText, fieldCreatedAt, fieldUpdatedAt}

	returningMessage = returning + strings.Join(messageFields, separator)
)

func tableField(table, field string) string {
	return fmt.Sprintf("%s.%s", table, field)
}

func tableFields(table string, fields []string) []string {
	var respFields []string
	for _, field := range fields {
		respFields = append(respFields, tableField(table, field))
	}
	return respFields
}

func mergeFields(firstFields []string, secondFields ...string) []string {
	for _, field := range secondFields {
		firstFields = append(firstFields, field)
	}
	return firstFields
}

func joinString(sourceTable, sourceField, joinTable, joinField string) string {
	return fmt.Sprintf("%[1]s ON %[2]s.%[3]s = %[1]s.%[4]s", joinTable, sourceTable, sourceField, joinField)
}
