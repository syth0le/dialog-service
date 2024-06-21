package postgres

import (
	"fmt"
)

const (
	DialogTable  = "dialog_table"
	MessageTable = "message_table"
)

const (
	returning = "RETURNING "
	separator = ","
)

const (
	fieldID = "id"

	fieldFirstUserID  = "first_user_id"
	fieldSecondUserID = "second_user_id"

	fieldDialogId    = "dialog_id"
	fieldText        = "text"
	fieldSenderId    = "sender_id"
	fieldRecipientId = "recipient_id"

	fieldCreatedAt = "created_at"
	fieldUpdatedAt = "updated_at"
	fieldDeletedAt = "deleted_at"
)

var (
	dialogFields = []string{
		fieldID, fieldFirstUserID, fieldSecondUserID, fieldCreatedAt,
	}
	messageFields = []string{fieldID, fieldDialogId, fieldSenderId, fieldRecipientId, fieldText, fieldCreatedAt, fieldUpdatedAt}
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
