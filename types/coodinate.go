package types

type Coordinates struct {
	Latitude  float32 `msgp:"latitude" `
	Longitude float32 `msgp:"longitude" `
}
type UserCoordinateItem struct {
	UserID    uint    `msgp:"user_id" db:"user_id,string"`
	Latitude  float32 `msgp:"latitude" db:"latitude"`
	Longitude float32 `msgp:"longitude" db:"longitude"`
	Distance  float32 `msgp:"distance" db:"distance"`
}
