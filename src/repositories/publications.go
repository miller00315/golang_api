package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications represents a repository of publications
type Publications struct {
	db *sql.DB
}

// NewPublicationRepository returns a new publication repository
func NewPublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

// CreatePublication
func (repository Publications) CreatePublication(publication models.Publication) (uint64, error) {
	statement, error := repository.db.Prepare("insert into publications (title, content, author_id) values (?, ?, ?)")

	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(publication.Title, publication.Content, publication.AuthorID)

	if error != nil {
		return 0, error
	}

	lastInsertedId, error := result.LastInsertId()

	if error != nil {
		return 0, error
	}

	return uint64(lastInsertedId), nil
}

// ListPublications
func (repository Publications) ListPublications() {}

// GetPublication
func (repository Publications) GetPublication(publicationID uint64) (models.Publication, error) {
	line, error := repository.db.Query(
		`SELECT p.*, u.nick from 
		publications p inner join users u
		on u.id = p.author_id where p.id = ?`,
		publicationID)

	if error != nil {
		return models.Publication{}, error
	}

	defer line.Close()

	var publication models.Publication

	if line.Next() {
		if error = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); error != nil {
			return models.Publication{}, error
		}
	}

	return publication, nil
}

// UpdatePublication
func (repository Publications) UpdatePublication(publication models.Publication) {}

// DeletePublication
func (repository Publications) DeletePublication(publicationID uint64) {}
