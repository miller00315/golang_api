package models

import (
	"errors"
	"strings"
	"time"
)

// Publication represents a user publication
type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitEmpty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createAt,omitEmpty"`
}

// Prepare validate and format a publication
func (publication *Publication) Prepare() error {
	if error := publication.validate(); error != nil {
		return error
	}

	publication.format()

	return nil
}

func (publication *Publication) validate() error {

	if publication.Title == "" {
		return errors.New("Tittle cannot be empty")
	}

	if publication.Content == "" {
		return errors.New("Content cannot be empty")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
