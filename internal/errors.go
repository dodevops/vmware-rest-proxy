package internal

// MissingAuthorizationHeaderError is issued when a request doesn't hold an Authorization header
type MissingAuthorizationHeaderError struct{}

var _ error = MissingAuthorizationHeaderError{}

func (m MissingAuthorizationHeaderError) Error() string {
	return "Missing authentication header"
}
