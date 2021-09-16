package github

func WriteFile(owner, repo, branch, filename, file, commitmessage, githubkey string) error {
	return WriteFiles(owner, repo, branch, map[string]string{filename: file}, commitmessage, githubkey)
}
