package dberrors

import (
	"fmt"
)

type NotFoundError struct {
	Entity string
	ID     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Error: record not found: entity=%s id=%s", e.Entity, e.ID)
}
