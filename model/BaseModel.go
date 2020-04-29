package model

import "time"

type BaseModel struct {
	ID 			int	`gorm:"primary_key"`
	CreateAt	time.Time
	UpdateAt	time.Time
}
