package models

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Document  string    `bson:"document"`
	Email     string    `bson:"email"`
	Phone     string    `bson:"phone"`
	PwdHash   string    `bson:"pwdHash"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdateAt  time.Time `bson:"updatedAt"`
	DeletedAt time.Time `bson:"deletedAt"`
}
