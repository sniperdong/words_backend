package model

import "time"

type VideoLogs struct {
	ID        uint      `gorm:"primaryKey"`
	VideoID   uint      `gorm:"video_id"`
	Content   string    `gorm:"content"`
	CreatedAt time.Time `gorm:"create_at;<-:false"`
	UpdatedAt time.Time `gorm:"create_at;<-:false"`
	Likes     int       `gorm:"likes"`
	Stars     int       `gorm:"stars"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *VideoLogs) Table() string {
	return "video_logs"
}
