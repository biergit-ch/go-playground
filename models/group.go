package models

// A user
// swagger:model Group
type Group struct {
	Base
	GroupName string  `json:"group_name"`
}
