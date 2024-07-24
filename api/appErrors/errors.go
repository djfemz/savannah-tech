package appErrors

const (
	COMMIT_NOT_FOUND      = "requested commit not found"
	AUTHORS_NOT_FOUND     = "authors not found"
	REPOSITORY_NOT_FOUND  = "repository not found"
	INVALID_TIME_SUPPLIED = "invalid time format, supply time in format, MM-DD-YYYY"
)

type commitNotFound struct {
}

type authorsNotFound struct {
}

type repositoryNotFound struct {
}

type timeFormatError struct {
}

func NewTimeFormatError() error {
	return &timeFormatError{}
}

func NewCommitNotFoundError() error {
	return &commitNotFound{}
}

func NewAuthorNotFoundError() error {
	return &authorsNotFound{}
}

func NewRepositoryNotFoundError() error {
	return &repositoryNotFound{}
}

func (commitNotFound *commitNotFound) Error() string {
	return COMMIT_NOT_FOUND
}

func (authorsNotFound *authorsNotFound) Error() string {
	return AUTHORS_NOT_FOUND
}

func (repositoryNotFound *repositoryNotFound) Error() string {
	return REPOSITORY_NOT_FOUND
}

func (timeFormatError *timeFormatError) Error() string {
	return INVALID_TIME_SUPPLIED
}
