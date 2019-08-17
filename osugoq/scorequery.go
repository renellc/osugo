package osugoq

import "github.com/renellc/osugo/osugof"

// ScoresQuery is used to fetch the scores set for a specified beatmap.
type ScoresQuery struct {
	BeatmapID string
	User      string
	Mode      osugof.GameMode
	Mods      int
	Type      osugof.UserType
	Limit     int
}
