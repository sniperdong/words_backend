package model

import "time"

type Videos struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"name"`
	Path      string    `gorm:"path"`
	Memo      string    `gorm:"memo"`
	Content   string    `gorm:"content"`
	Publish   int       `gorm:"publish"`
	CreatedAt time.Time `gorm:"create_at;<-:false"`
	UpdatedAt time.Time `gorm:"create_at;<-:false"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Videos) Table() string {
	return "videos"
}
