package osugoq

import (
	"time"

	"github.com/renellc/osugo/osugof"
)

// BeatmapQuery is used to get a list of beatmaps.
type BeatmapQuery struct {
	Since            time.Time
	BeatmapSetID     string
	BeatmapID        string
	User             string
	Type             osugof.UserType
	Mode             osugof.GameMode
	IncludeConverted bool
	BeatmapHash      string
	Limit            int
	Mods             int
}
