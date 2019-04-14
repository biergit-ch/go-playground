package models

// A transaction will be used to store a beer owe
// swagger:model Transaction
type Transaction struct {
	Base
	Context   Group `gorm:"foreignkey:ContextID"`
	Source    User  `gorm:"foreignkey:SourceID"`
	Target    User  `gorm:"foreignkey:TargetID"`
	Amount    int
	SourceID  int `json:"-"`
	TargetID  int `json:"-"`
	ContextID int `json:"-"`
}
