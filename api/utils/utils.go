package utils

import (
	"database/sql"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDatabase     = errors.New("database error")
)

var (
	StatusPending    = "1"
	StatusCompleted  = "2"
	StatusCancelled  = "3"
	StatusRejected   = "4"
	StatusExpired    = "5"
	StatusProcessing = "6"
	StatusPartial    = "7"
	StatusFailed     = "8"

	StatusPendingText    = "Pending"
	StatusCompletedText  = "Completed"
	StatusCancelledText  = "Cancelled"
	StatusRejectedText   = "Rejected"
	StatusExpiredText    = "Expired"
	StatusProcessingText = "Processing"
	StatusPartialText    = "Partial"
	StatusFailedText     = "Failed"
)

func MapError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrUserNotFound
	default:
		return ErrDatabase
	}
}

// Switch Status From Database
func Status(status string) string {
	switch status {
	case StatusPending:
		return StatusPendingText
	case StatusCompleted:
		return StatusCompletedText
	case StatusCancelled:
		return StatusCancelledText
	case StatusRejected:
		return StatusRejectedText
	case StatusExpired:
		return StatusExpiredText
	case StatusProcessing:
		return StatusProcessingText
	case StatusPartial:
		return StatusPartialText
	case StatusFailed:
		return StatusFailedText
	default:
		return "Status Not Found"
	}
}

func UnPtr[A any](a *A) A {
	return *a
}

func Ptr[A any](a A) *A {
	return &a
}
