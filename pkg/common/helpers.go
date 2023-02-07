package common

import (
	"errors"
	"html"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"

	"github.com/aaryanraj/ebook-store-api/pkg/clients"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BeforeSave(pas string) (string, error) {
	hashedPassword, err := Hash(pas)
	if err != nil {
		return "", err
	}
	password := string(hashedPassword)
	return password, nil
}

func Prepare(u *clients.User) {
	u.FullName = html.EscapeString(strings.TrimSpace(u.FullName))
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func Validate(u *clients.User, action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.FullName == "" {
			return errors.New("Required Full Name")
		}
		if u.UserName == "" {
			return errors.New("Required User Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.UserName == "" {
			return errors.New("Required User Name")
		}
		if u.FullName == "" {
			return errors.New("Required Full Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}
