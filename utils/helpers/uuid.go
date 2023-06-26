package helpers

import "github.com/google/uuid"

func GenerateNewId() string {
	id := uuid.New()
	return id.String()
}