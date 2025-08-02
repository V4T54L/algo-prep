package repository

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
	"sync"
	"time"
)

// InMemoryContestRepo is a thread-safe, in-memory implementation of a contest repository.
type InMemoryContestRepo struct {
	mu       sync.RWMutex
	contests map[string]*aggregate.ContestInfo
}

// NewInMemoryContestRepo returns an in-memory repo preloaded with static contest data.
func NewInMemoryContestRepo() repository.ContestRepository {
	return &InMemoryContestRepo{
		contests: map[string]*aggregate.ContestInfo{
			"1": {
				ID:          "1",
				Title:       "Spring Coding Challenge",
				Difficulty:  "Medium",
				Description: "A coding challenge for the spring season with a mix of algorithmic and problem-solving questions.",
				StartTime:   time.Date(2025, time.August, 5, 14, 0, 0, 0, time.UTC),
				Duration:    2 * time.Hour,
				Problems:    5,
				Prizes:      []string{"₹5000", "₹3000", "₹1000"},
				Registered:  120,
			},
			"2": {
				ID:          "2",
				Title:       "Summer Algorithmic Contest",
				Difficulty:  "Hard",
				Description: "An advanced contest for competitive programmers with challenging algorithmic problems.",
				StartTime:   time.Date(2025, time.August, 15, 18, 0, 0, 0, time.UTC),
				Duration:    3 * time.Hour,
				Problems:    8,
				Prizes:      []string{"₹10000", "₹5000", "₹2000"},
				Registered:  80,
			},
			"3": {
				ID:          "3",
				Title:       "Beginner’s Coding Tournament",
				Difficulty:  "Easy",
				Description: "A fun contest to introduce new programmers to competitive coding with easy problems.",
				StartTime:   time.Date(2025, time.August, 20, 10, 0, 0, 0, time.UTC),
				Duration:    1 * time.Hour,
				Problems:    4,
				Prizes:      []string{"₹2000", "₹1000", "₹500"},
				Registered:  250,
			},
		},
	}
}

// List returns all contests from the in-memory repo.
func (r *InMemoryContestRepo) List(ctx context.Context) ([]*aggregate.ContestInfo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*aggregate.ContestInfo, 0, len(r.contests))
	for _, c := range r.contests {
		result = append(result, c)
	}
	return result, nil
}
