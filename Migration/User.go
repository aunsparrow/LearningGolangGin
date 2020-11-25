package Migration

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FirstName string    `gorm:"type:varchar(100)"`
	LastName  string    `gorm:"type:varchar(100)"`
	Age       int
	Active    int
	CreatedAt time.Time `gorm:"type:timestamp"`
}
