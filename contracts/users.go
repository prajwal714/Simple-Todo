package contracts

import "time"

type User struct {
	ID        int64      `json:"id"`
	UUID      string     `json:"uuid"`
	Email     string     `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
