
package aggregate

import "time"

type Post struct {
	ID        string
	Title     string
	Content   string
	AuthorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPost(id, title, content, authorID string) *Post {
	return &Post{
		ID:        id,
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
