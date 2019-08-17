package osugoq

import "github.com/renellc/osugo/osugof"

// UserPerfQuery is a query that's used to get either a user's best scores or a user's recent
// scores.
type UserPerfQuery struct {
	User  string
	Mode  osugof.GameMode
	Limit int
	Type  osugof.UserType
}
