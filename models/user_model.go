package models

import (
	"time"
)

type User struct {
	Id        int64          `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"varchar(300); not null" json:"username"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"varchar(300);not null;size:255" json:"password"`
	Photos    []Photo `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"photos"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
