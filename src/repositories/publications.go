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
func (repository Publications) ListPublications(userID uint64) ([]models.Publication, error) {

	lines, error := repository.db.Query(`
	SELECT distinct p.*, u.nick from publications p 
	inner join users u on u.id = p.author_id 
	inner join followers f on p.author_id = f.user_id 
	where u.id = ? or f.follower_id = ?
	order by 1 desc`,
		userID, userID,
	)

	if error != nil {
		return nil, nil
	}

	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if error = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); error != nil {
			return nil, error
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

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
func (repository Publications) UpdatePublication(publication models.Publication, publicationID uint64) error {

	statement, error := repository.db.Prepare("update publications set title = ?, content = ? where id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error := statement.Exec(publication.Title, publication.Content, publicationID); error != nil {
		return error
	}

	return nil
}

// DeletePublication
func (repository Publications) DeletePublication(publicationID uint64) error {
	statement, error := repository.db.Prepare("DELETE FROM publications WHERE id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(publicationID); error != nil {
		return error
	}

	return nil
}

// ListUserPublications
func (repository Publications) ListUserPublications(userID uint64) ([]models.Publication, error) {
	lines, error := repository.db.Query(
		`select p.*, u.nick from publications p 
		join users u on u.id = p.author_id 
		where p.author_id = ?`,
		userID)

	if error != nil {
		return nil, nil
	}

	defer lines.Close()

	var publications []models.Publication

	for lines.Next() {
		var publication models.Publication

		if error = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); error != nil {
			return nil, error
		}

		publications = append(publications, publication)
	}

	return publications, nil

}

// LikePublication
func (repository Publications) LikePublication(publicationID uint64) error {
	statement, error := repository.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(publicationID); error != nil {
		return error
	}

	return nil
}

// UnLikePublication
func (repository Publications) UnLikePublication(publicationID uint64) error {
	statement, error := repository.db.Prepare(`
	UPDATE publications SET likes = 
	CASE 
		WHEN likes > 0 THEN likes - 1 
		ELSE likes 
	END
	WHERE id = ?`)

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(publicationID); error != nil {
		return error
	}

	return nil
}
