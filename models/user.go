package models

// The user represents a reference to a oauth user id
// swagger:model User
type User struct {
	Base
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsMature  bool `json:"is_mature"`
}
