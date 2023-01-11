package authorization

import (
	"errors"
	"net/mail"
	db "postgres/database"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var auth AuthorizedUser

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func validUser(id string, p string) bool {

	db.Connect.First(&auth, id)
	if auth.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, auth.Password) {
		return false
	}
	return true
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*AuthorizedUser, error) {

	if err := db.Connect.Where(&AuthorizedUser{Email: e}).Find(&auth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &auth, nil
}

func getUserByUsername(u string) (*AuthorizedUser, error) {
	if err := db.Connect.Where(&AuthorizedUser{Username: u}).Find(&auth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &auth, nil
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
