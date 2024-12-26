package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	TransactionTypeDeposit = "deposit"
)

type User struct {
	gorm.Model
	Email    string `faker:"email"`
	Password string `faker:"password"`
}

type Transaction struct {
	gorm.Model
	UserID uint
	User   User
	Amount int
	Type   string
}

type BalanceSnapshot struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint
	User             User
	LastKnownBalance int
	Timestamp        time.Time
}
