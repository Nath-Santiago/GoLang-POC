package authorization

import "gorm.io/gorm"

type AuthorizedUser struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Names    string `json:"names"`
}

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
type RegisterUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type UserData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
