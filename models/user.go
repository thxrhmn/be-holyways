package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"type: varchar(255)"`
}
