package object

import "time"

type User struct {
	ID        string
	Name      string
	Token     string
	HighScore int64
	Coin      int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
