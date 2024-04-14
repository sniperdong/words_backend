package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"name"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"create_at;<-:false"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *User) Table() string {
	return "users"
}
