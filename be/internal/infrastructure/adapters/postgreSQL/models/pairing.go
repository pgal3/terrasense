package pg_models

import "time"

type Pairing struct {
	UserID            string `gorm:"primaryKey;not null"`
	ChipID            int32  `gorm:"primaryKey;autoIncrement:false;not null"`
	NotifyMe          bool   `gorm:"default:false"`
	LowLevelThreshold int64  `gorm:"default:400"`
	CreatedAt         time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt         time.Time `gorm:"not null;default:current_timestamp"`
}
