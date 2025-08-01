package repository

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
	"sync"
)

// InMemoryProblemRepo is a thread-safe, in-memory implementation of a problem repository.
type InMemoryProblemRepo struct {
	mu       sync.RWMutex
	problems map[int]*aggregate.ProblemInfo
}

// NewInMemoryProblemRepo returns an in-memory repo preloaded with static problem data.
func NewInMemoryProblemRepo() repository.ProblemRepository {
	return &InMemoryProblemRepo{
		problems: map[int]*aggregate.ProblemInfo{
			1: {
				ID:          1,
				Title:       "Two Sum",
				Difficulty:  "Easy",
				Categories:  []string{"Array", "Hash Map"},
				Description: "Given an array of integers, return indices of the two numbers such that they add up to a specific target.",
				Likes:       24567,
				Solved:      true,
			},
			2: {
				ID:          2,
				Title:       "Longest Substring Without Repeating Characters",
				Difficulty:  "Medium",
				Categories:  []string{"String", "Sliding Window"},
				Description: "Find the length of the longest substring without repeating characters.",
				Likes:       19234,
				Solved:      false,
			},
			3: {
				ID:          3,
				Title:       "Median of Two Sorted Arrays",
				Difficulty:  "Hard",
				Categories:  []string{"Array", "Binary Search", "Divide and Conquer"},
				Description: "Given two sorted arrays, find the median of the combined sorted array in O(log (m+n)) time.",
				Likes:       28901,
				Solved:      false,
			},
			4: {
				ID:          4,
				Title:       "Climbing Stairs",
				Difficulty:  "Easy",
				Categories:  []string{"DP"},
				Description: "You are climbing a staircase. It takes n steps to reach the top. Each time you can climb 1 or 2 steps. How many distinct ways can you climb to the top?",
				Likes:       15832,
				Solved:      true,
			},
			5: {
				ID:          5,
				Title:       "Word Break",
				Difficulty:  "Medium",
				Categories:  []string{"String", "DP"},
				Description: "Given a string and a dictionary of words, determine if the string can be segmented into space-separated words in the dictionary.",
				Likes:       13374,
				Solved:      false,
			},
			6: {
				ID:          6,
				Title:       "LFU Cache",
				Difficulty:  "Hard",
				Categories:  []string{"Design", "Hash Map", "Heap"},
				Description: "Design and implement a data structure for Least Frequently Used (LFU) cache with O(1) time complexity for get and put.",
				Likes:       8993,
				Solved:      false,
			},
		},
	}
}

// List returns all problems from the in-memory repo.
func (r *InMemoryProblemRepo) List(ctx context.Context) ([]*aggregate.ProblemInfo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*aggregate.ProblemInfo, 0, len(r.problems))
	for _, p := range r.problems {
		result = append(result, p)
	}
	return result, nil
}
