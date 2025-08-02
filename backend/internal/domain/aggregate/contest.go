package aggregate

import "time"

type ContestInfo struct {
	ID          string
	Title       string
	Description string
	Difficulty  string // Easy, Medium, Hard
	StartTime   time.Time
	Duration    time.Duration
	Problems    int
	Prizes      []string // e.g., []string{"₹5000", "₹3000", "₹2000"}
	Registered  int
}