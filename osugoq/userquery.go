package osugoq

import "github.com/renellc/osugo/osugof"

// UserQuery is used to fetch user information.
type UserQuery struct {
	User      string
	Mode      osugof.GameMode
	Type      osugof.UserType
	EventDays int
}
