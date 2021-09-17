package github

import "io/fs"

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

func (fsys *FileSystem) Open(name string) (fs.File, error) {
	return nil, nil
}
