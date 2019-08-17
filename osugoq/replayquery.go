package osugoq

import "github.com/renellc/osugo/osugof"

// ReplayQuery is used to fetch the replay data of a user's score.
type ReplayQuery struct {
	Mode      osugof.GameMode
	BeatmapID string
	User      string
	Type      osugof.UserType
	Mods      int
}
