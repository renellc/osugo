package osugo

// Replay is a struct that contains the data of an osu! replay file.
type Replay struct {
	Encoding string `json:"encoding"`
	Content  string `json:"content"`
}

// ReplayQuery is used to fetch the replay data of a user's score.
type ReplayQuery struct {
	Mode      GameMode
	BeatmapID string
	User      string
	Type      UserType
	Mods      int
}
