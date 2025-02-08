package models

type Route struct {
	Id    int      `json:"id" db:"id"`
	Name  string   `json:"name" db:"name"`
	Stops []string `json:"stops" db:"stops"`
}
