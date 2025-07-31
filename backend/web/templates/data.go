package templates

type navOption struct {
	name string
	href string
	icon string
}

var navigation = []navOption{
	{name: "Problems", href: "/problems", icon: "code"},
	{name: "Contests", href: "/contests", icon: "trophy"},
	{name: "Posts", href: "/posts", icon: "pen-tool"},
	{name: "Courses", href: "/courses", icon: "book-open"},
	{name: "Dashboard", href: "/dashboard", icon: "users"},
}

type Post struct {
	ImageURL     string
	Tags         []string
	Title        string
	Description  string
	Views        int
	Comments     int
	Date         string
	ReadMoreLink string
}

var featuredPostsData = []Post{
	{
		ImageURL:     "https://picsum.photos/600/300?random=1",
		Tags:         []string{"DSA", "Binary Search", "Interview Prep"},
		Title:        "Mastering Binary Search: From Basics to Advanced Patterns",
		Description:  "Learn the fundamental concepts of binary search and advanced patterns used in top tech company coding interviews.",
		Views:        234,
		Comments:     6,
		Date:         "1/12/2024",
		ReadMoreLink: "#",
	},
	{
		ImageURL:     "https://picsum.photos/600/300?random=2",
		Tags:         []string{"Dynamic Programming", "FAANG", "Patterns"},
		Title:        "Dynamic Programming Decoded: Top 15 Patterns You Must Know",
		Description:  "Comprehensive guide to dynamic programming patterns that appear frequently in FAANG interviews with practical examples.",
		Views:        183,
		Comments:     3,
		Date:         "1/6/2024",
		ReadMoreLink: "#",
	},
	{
		ImageURL:     "https://picsum.photos/600/300?random=3",
		Tags:         []string{"System Design", "Architecture", "Scalability"},
		Title:        "Designing Scalable Systems: A Deep Dive into System Design Interviews",
		Description:  "Understand trade-offs, scalability principles, and core system design patterns used in real-world interviews.",
		Views:        401,
		Comments:     11,
		Date:         "1/9/2024",
		ReadMoreLink: "#",
	},
	{
		ImageURL:     "https://picsum.photos/600/300?random=4",
		Tags:         []string{"Behavioral", "STAR Method", "Soft Skills"},
		Title:        "Behavioral Interview Questions: How to Answer Using the STAR Method",
		Description:  "Ace behavioral rounds with structured responses that highlight your experience and thought process effectively.",
		Views:        315,
		Comments:     8,
		Date:         "1/10/2024",
		ReadMoreLink: "#",
	},
	{
		ImageURL:     "https://picsum.photos/600/300?random=5",
		Tags:         []string{"Resume", "Job Hunt", "ATS"},
		Title:        "Resume Tips That Actually Work: Beat the ATS and Get Noticed",
		Description:  "Discover how to craft a clean, ATS-friendly resume that gets callbacks from top companies.",
		Views:        507,
		Comments:     4,
		Date:         "1/11/2024",
		ReadMoreLink: "#",
	},
	{
		ImageURL:     "https://picsum.photos/600/300?random=6",
		Tags:         []string{"Frontend", "Performance", "Best Practices"},
		Title:        "Optimizing Frontend Performance for Technical Interviews",
		Description:  "Get a solid grasp of performance bottlenecks, rendering patterns, and JS optimizations.",
		Views:        268,
		Comments:     5,
		Date:         "1/14/2024",
		ReadMoreLink: "#",
	},
}
