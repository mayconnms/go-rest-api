package domain

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"Id"`
	Name         string         `json:"Name"`
	Email        *string        `json:"Email"`
	Age          uint8          `json:"Age"`
	Birthday     *time.Time     `json:"Birthday"`
	MemberNumber sql.NullString `json:"MemberNumber"`
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
