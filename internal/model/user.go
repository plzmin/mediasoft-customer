package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Uuid       uuid.UUID `db:"uuid"`
	Name       string    `db:"name"`
	OfficeUuid uuid.UUID `db:"office_uuid"`
	OfficeName string    `db:"office_name"`
	CreatedAt  time.Time `db:"created_at"`
}
