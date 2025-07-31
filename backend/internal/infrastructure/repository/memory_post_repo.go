package repository

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
	"errors"
	"fmt"
	"sync"
)

type InMemoryPostRepo struct {
	mu    sync.RWMutex
	posts map[int]*aggregate.PostDetail
}

func NewInMemoryPostRepo() repository.PostRepository {

	defaultAuthor := aggregate.Author{
		ID:          1,
		Username:    "jdoe",
		Name:        "John Doe",
		AvatarURL:   "https://i.pravatar.cc/150?img=1",
		Designation: "Software Engineer",
	}

	commentAuthor := aggregate.Author{
		ID:          2,
		Username:    "alice123",
		Name:        "Alice Smith",
		AvatarURL:   "https://i.pravatar.cc/150?img=2",
		Designation: "Tech Writer",
	}

	comments := []aggregate.Comment{
		{
			Author:    commentAuthor,
			CreatedAt: "1/15/2024",
			Content:   "Great post! Super helpful.",
			Likes:     12,
			Dislikes:  0,
		},
		{
			Author:    commentAuthor,
			CreatedAt: "1/16/2024",
			Content:   "Can you elaborate on the third pattern?",
			Likes:     5,
			Dislikes:  1,
		},
	}

	repo := &InMemoryPostRepo{
		posts: map[int]*aggregate.PostDetail{

			1: &aggregate.PostDetail{
				ID:        1,
				ImageURL:  "https://picsum.photos/600/300?random=1",
				Tags:      []string{"DSA", "Binary Search", "Interview Prep"},
				Title:     "Mastering Binary Search: From Basics to Advanced Patterns",
				Content:   "Learn the fundamental concepts of binary search and advanced patterns used in top tech company coding interviews.",
				Views:     234,
				CreatedAt: "1/12/2024",
				Comments:  comments,
				Author:    defaultAuthor,
			},

			2: &aggregate.PostDetail{
				ID:        2,
				ImageURL:  "https://picsum.photos/600/300?random=2",
				Tags:      []string{"Dynamic Programming", "FAANG", "Patterns"},
				Title:     "Dynamic Programming Decoded: Top 15 Patterns You Must Know",
				Content:   "Comprehensive guide to dynamic programming patterns that appear frequently in FAANG interviews with practical examples.",
				Views:     183,
				CreatedAt: "1/6/2024",
				Comments:  comments,
				Author:    defaultAuthor,
			},

			3: &aggregate.PostDetail{
				ID:        3,
				ImageURL:  "https://picsum.photos/600/300?random=3",
				Tags:      []string{"System Design", "Architecture", "Scalability"},
				Title:     "Designing Scalable Systems: A Deep Dive into System Design Interviews",
				Content:   "Understand trade-offs, scalability principles, and core system design patterns used in real-world interviews.",
				Views:     401,
				CreatedAt: "1/9/2024",
				Comments:  comments,
				Author:    defaultAuthor,
			},

			4: &aggregate.PostDetail{
				ID:        4,
				ImageURL:  "https://picsum.photos/600/300?random=4",
				Tags:      []string{"Behavioral", "STAR Method", "Soft Skills"},
				Title:     "Behavioral Interview Questions: How to Answer Using the STAR Method",
				Content:   "Ace behavioral rounds with structured responses that highlight your experience and thought process effectively.",
				Views:     315,
				CreatedAt: "1/10/2024",
				Comments:  comments,
				Author:    defaultAuthor,
			},

			5: &aggregate.PostDetail{
				ID:        5,
				ImageURL:  "https://picsum.photos/600/300?random=5",
				Tags:      []string{"Resume", "Job Hunt", "ATS"},
				Title:     "Resume Tips That Actually Work: Beat the ATS and Get Noticed",
				Content:   "Discover how to craft a clean, ATS-friendly resume that gets callbacks from top companies.",
				Views:     507,
				CreatedAt: "1/11/2024",
				Comments:  comments,
				Author:    defaultAuthor,
			},

			6: &aggregate.PostDetail{
				ID:        6,
				ImageURL:  "https://picsum.photos/600/300?random=6",
				Tags:      []string{"Frontend", "Performance", "Best Practices"},
				Title:     "Optimizing Frontend Performance for Technical Interviews",
				Content:   "Get a solid grasp of performance bottlenecks, rendering patterns, and JS optimizations.",
				Views:     268,
				CreatedAt: "1/14/2024",
				Comments:  comments,
				Author:    defaultAuthor,
			},
		},
	}

	return repo
}

func (r *InMemoryPostRepo) Save(ctx context.Context, post *aggregate.PostDetail) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.posts[post.ID] = post
	return nil
}

func (r *InMemoryPostRepo) FindByID(ctx context.Context, id int) (*aggregate.PostDetail, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	post, ok := r.posts[id]
	if !ok {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (r *InMemoryPostRepo) List(ctx context.Context) ([]*aggregate.PostInfo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	posts := make([]*aggregate.PostInfo, 0, len(r.posts))
	for _, postDetail := range r.posts {
		info := convertPostDetailToPostInfo(*postDetail)
		postCopy := info
		posts = append(posts, &postCopy)
	}

	return posts, nil
}

func convertPostDetailToPostInfo(detail aggregate.PostDetail) aggregate.PostInfo {
	shortContent := getShortContent(detail.Content, 150)
	readMoreLink := fmt.Sprintf("/posts/%d", detail.ID)

	return aggregate.PostInfo{
		ID:           detail.ID,
		ImageURL:     detail.ImageURL,
		Tags:         detail.Tags,
		Title:        detail.Title,
		ShortContent: shortContent,
		Views:        detail.Views,
		Comments:     len(detail.Comments),
		CreatedAt:    detail.CreatedAt,
		ReadMoreLink: readMoreLink,
		Author:       detail.Author,
	}
}

func getShortContent(content string, limit int) string {
	if len(content) <= limit {
		return content
	}
	return content[:limit] + "..."
}
