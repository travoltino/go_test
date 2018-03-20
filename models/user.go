package models

// User sdfsd
type User struct {
	ID       int    `gorm:"column:id" json:"id"`
	UserName string `gorm:"column:user_name" json:"user_name"`
	DomainID int    `gorm:"column:domain_id" json:"domain_id"`
}

// TableName sdfgsd
func (User) TableName() string {
	return "user"
}

// Users dfd
type Users []User
