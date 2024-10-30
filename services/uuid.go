package services

import "github.com/google/uuid"

func generateId() string {
	return uuid.NewString()
}
