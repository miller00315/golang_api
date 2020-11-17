package models

import "time"

// User cria um usuário
type User struct {
	Id        uint      `json:"id,omitempty'`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Password  string    `json:"password,omitEmpty"`
	CreatedAt time.Time `json:"CreatedAt,omitEmpty"`
}
