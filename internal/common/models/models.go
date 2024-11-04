package models

import (
	"errors"
)

var (
	ErrForbidden = errors.New("forbidden: you don't have permission for this action")
)

type User struct {
	ID   string `json:"user_id"`
	Name string `json:"user_name"`
}
