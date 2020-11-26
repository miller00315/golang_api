package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users is a repository od users
type Users struct {
	db *sql.DB
}

// NewUserRepository generate a new repository of user
func NewUserRepository(db *sql.DB) *Users {

	return &Users{db}
}

// Insert a new user in database
func (repository Users) Create(user models.User) (uint64, error) {

	statement, error := repository.db.Prepare("insert into users (name, nick, email, password) values(?,?,?,?)")

	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if error != nil {
		return 0, error
	}

	lastInsertedID, error := result.LastInsertId()

	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedID), nil
}

// Search find all users that has the parameter userQuery
func (repository Users) Search(userQuery string) ([]models.User, error) {
	userQuery = fmt.Sprintf("%%%s%%", userQuery) // %% serve para o escape de caracteres

	lines, error := repository.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?",
		userQuery, userQuery)

	if error != nil {
		return nil, error
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

// Get get a user by id
func (repository Users) Get(ID uint64) (models.User, error) {
	lines, error := repository.db.Query("SELECT id, name, nick, email, createdAt from users where id = ?",
		ID)

	if error != nil {
		return models.User{}, error
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

// Update update a user
func (repository Users) Update(ID uint64, user models.User) error {

	statement, error := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? where id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, ID); error != nil {
		return error
	}

	return nil
}

// Delete delete a user from database
func (repository Users) Delete(ID uint64) error {
	statement, error := repository.db.Prepare("DELETE FROM users WHERE id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(ID); error != nil {
		return error
	}

	return nil
}

// SearchByEmail get a user by email 
func (repository Users) SearchByEmail(email string) (models.User, error) {
	line, error := repository.db.Query("SELECT id, password FROM users where email = ?", email)

	if error != nil {
		return models.User{}, error
	}

	defer line.Close()

	var user models.User

	if line.Next() {
		if error = line.Scan(&user.ID, & user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}
