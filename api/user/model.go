package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email     string    `gorm:"type:varchar(120);not null;unique"`
	Password  string    `gorm:"type:varchar(120);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;null;default:null"`
	DeletedAt time.Time `gorm:"type:timestamp;null;default:null"`
}

func (p *User) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
