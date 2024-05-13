package models

type ErrorType string

func (e ErrorType) Error() string {
	return string(e)
}

const (
	ErrCommandNotFound    ErrorType = "command not found"
	ErrUpdateEmptyMessage ErrorType = "update has empty message"
)
