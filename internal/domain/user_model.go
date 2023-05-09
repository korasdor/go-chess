package domain

import (
	"time"
)

type UserData struct {
	Id           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	Phone        string    `json:"phone" db:"phone"`
	Password     string    `json:"-" db:"password"`
	RegisteredAt time.Time `json:"registeredAt" db:"registered_at"`
	LastVisitAt  time.Time `json:"lastVisitAt" db:"last_visit_at"`
	Wallets      []Wallet  `json:"wallets"`
}

type Wallet struct {
	Id      string `json:"id" db:"id"`
	Balance string `json:"balance" db:"balance"`
}
