package model

import (
	"github.com/google/uuid"
	"gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"time"
)

type Order struct {
	Uuid       uuid.UUID             `json:"uuid" db:"uuid"`
	UserUuid   uuid.UUID             `json:"userUuid" db:"user_uuid"`
	OfficeUuid uuid.UUID             `json:"officeUuid"`
	Salads     []*customer.OrderItem `json:"salads,omitempty"`
	Garnishes  []*customer.OrderItem `json:"garnishes,omitempty"`
	Meats      []*customer.OrderItem `json:"meats,omitempty"`
	Soups      []*customer.OrderItem `json:"soups,omitempty"`
	Drinks     []*customer.OrderItem `json:"drinks,omitempty"`
	Desserts   []*customer.OrderItem `json:"desserts,omitempty"`
	CreatedAt  time.Time             `json:"created_at" db:"created_at"`
}
