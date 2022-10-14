package user

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	UserName  string    `json:"user_name" gorm:"column:user_name"`
	Email     string    `json:"email" gorm:"column:email"`
	Dbo       time.Time `json:"dbo" gorm:"column:dbo"`
	Age       int       `json:"age" gorm:"column:age"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}
