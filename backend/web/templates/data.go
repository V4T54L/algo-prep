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
