package brand

import "time"

type Brand struct {
	Brd_Id    int `gorm:"primary_key"`
	Brd_Name  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
