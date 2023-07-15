package model

import (
	"github.com/google/uuid"
	"time"
)

type Office struct {
	Uuid      uuid.UUID `db:"uuid" `
	Name      string    `db:"name"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
}
