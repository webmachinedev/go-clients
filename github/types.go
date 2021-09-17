package github

type Tree struct {
	Owner, Repo, Branch string
}

type Client struct {
	Name string
	Token string
}

type FileSystem struct {
	Tree *Tree
	Client *Client
}
