package osugof

// UserType indicates if a given user's value was their username or user ID.
//
// This is used as an optional parameter for queries that ask for a user's name or ID.
type UserType string

const (
	// Username indicates the value used for the user parameter was a username.
	Username UserType = "string"
	// ID indicates the value used for the user parameter was a user ID.
	ID UserType = "id"
)
