package models

import (
	"api/src/security"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User cria um usuário
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Password  string    `json:"password,omitEmpty"`
	CreatedAt time.Time `json:"CreatedAt,omitEmpty"`
}

// Prepare call the methods to validate and format user
func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	if error := user.format(step); error != nil {
		return error
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New(errorMessage("name"))
	}

	if user.Email == "" {
		return errors.New(errorMessage("email"))
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("Invalid email")
	}

	if user.Nick == "" {
		return errors.New(errorMessage("nick"))
	}

	if user.Password == "" && step == "register" {
		return errors.New(errorMessage("password"))
	}

	return nil
}

func errorMessage(field string) string {
	return fmt.Sprintf("Field %s cannot be empty", field)
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPassword, error := security.Hash(user.Password)

		if error != nil {
			return error
		}

		user.Password = string(hashPassword)
	}

	return nil
}
