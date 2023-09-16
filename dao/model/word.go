package model

import "time"

type Word struct {
	ID                uint      `gorm:"primaryKey"`
	Word              string    `gorm:"word"`
	Means             string    `gorm:"means"`
	Pronounce         string    `gorm:"pronounce"`
	Sentences         string    `gorm:"sentences"`
	Plural            string    `gorm:"plural"`
	PastTense         string    `gorm:"past_tense"`
	PastParticiple    string    `gorm:"past_participle"`
	PresentParticiple string    `gorm:"present_participle"`
	Property          int       `gorm:"property"`
	CreatedAt         time.Time `gorm:"create_at;<-:false"`
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Word) Table() string {
	return "words"
}
