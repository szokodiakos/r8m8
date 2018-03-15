package errors

import "fmt"

// RepoPlayerNotFoundByUniqueNameError struct
type RepoPlayerNotFoundByUniqueNameError struct {
	UniqueName string
}

func (e *RepoPlayerNotFoundByUniqueNameError) Error() string {
	return fmt.Sprintf("Repo Player Not Found By Unique Name: %s", e.UniqueName)
}

// NewRepoPlayerNotFoundByUniqueNameError factory
func NewRepoPlayerNotFoundByUniqueNameError(uniqueName string) error {
	return &RepoPlayerNotFoundByUniqueNameError{
		UniqueName: uniqueName,
	}
}
