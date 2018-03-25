package entity

// Player struct
type Player struct {
	ID          string `db:"id"`
	DisplayName string `db:"display_name"`
}
