package osugo

// Score is a struct that contains osu! data for a given score.
type Score struct {
	ScoreID         string   `json:"score_id"`
	Username        string   `json:"username"`
	AchievedDate    string   `json:"date"`
	PP              float32  `json:"pp,string"`
	ReplayAvailable JSONBool `json:"replay_available,string"`
	ScoreBase
}

// ScoresQuery is used to fetch the scores set for a specified beatmap.
type ScoresQuery struct {
	BeatmapID string
	User      string
	Mode      GameMode
	Mods      int
	Type      UserType
	Limit     int
}
