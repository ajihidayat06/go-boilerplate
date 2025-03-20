package request

import (
	"errors"
	"time"
)

type AbstractRequest struct {
	UpdatedAt    time.Time 
	UpdatedAtStr string `json:"updated_at"`
	CreateddAt   time.Time 
	CreatedAtStr string `json:"created_at"`
}

func (a *AbstractRequest) ValidateUpdatedAt() error {
    parsedTime, err := time.Parse(time.RFC3339, a.UpdatedAtStr)
    if err != nil {
        return errors.New("invalid updated_at format, must be RFC3339")
    }

    a.UpdatedAt = parsedTime

    if a.UpdatedAt.IsZero() {
        return errors.New("updated_at cannot be empty")
    }

    return nil
}
