package util

import "github.com/google/uuid"

func UUIDMustParse(id string) string {
	return uuid.MustParse(id).String()
}
