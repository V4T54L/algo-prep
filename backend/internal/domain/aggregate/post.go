package aggregate

import "time"

type PostInfo struct {
	ID           int
	ImageURL     string
	Tags         []string
	Title        string
	ShortContent string
	Views        int
	Comments     int
	CreatedAt    string
	ReadMoreLink string
	Author Author
}

type PostDetail struct {
	ID        int
	ImageURL  string
	Tags      []string
	Title     string
	Content   string
	Views     int
	CreatedAt string
	Comments  []Comment
	Author    Author
}

type Comment struct {
	Author    Author
	CreatedAt string
	Content   string
	Likes     int
	Dislikes  int
}

func NewPost(imageUrl, title, content string, tags []string) *PostDetail {
	return &PostDetail{
		ID:        int(time.Microsecond),
		ImageURL:  imageUrl,
		Tags:      tags,
		Title:     title,
		Content:   "",
		Views:     0,
		CreatedAt: "",
		Comments:  []Comment{},
	}
}
