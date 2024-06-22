package utils

import (
	"strings"

	"github.com/google/uuid"
)

const (
	serviceNamePrefix       = "dlg"
	dialogEntityPrefix      = "d"
	participantEntityPrefix = "p"
	messageEntityPrefix     = "m"
)

func GenerateDUID() string {
	return generateUID(dialogEntityPrefix)
}

func GeneratePUID() string {
	return generateUID(participantEntityPrefix)
}

func GenerateMUID() string {
	return generateUID(messageEntityPrefix)
}

func generateUID(entityPrefix string) string {
	return serviceNamePrefix + entityPrefix + strings.Replace(uuid.New().String(), "-", "", -1)
}
