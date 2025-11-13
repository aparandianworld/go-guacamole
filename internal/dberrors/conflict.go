package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "Error: attempted to create a record with an existing key."
}
