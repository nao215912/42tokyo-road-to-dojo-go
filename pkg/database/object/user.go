package object

import "time"

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Token     string    `db:"token"`
	HighScore int64     `db:"high_score"`
	Coin      int64     `db:"coin"`
	CreatedAt time.Time `db:"create_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
