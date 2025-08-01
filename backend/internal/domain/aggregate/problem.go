package aggregate

type ProblemInfo struct {
	ID          int
	Title       string
	Difficulty  string   // "Easy", "Medium", "Hard"
	Categories  []string // e.g., ["Array", "DP"]
	Description string
	Likes       int
	Solved      bool
}
