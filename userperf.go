package osugo

// BestScore represents a score in a user's top plays.
type BestScore struct {
	BeatmapID    string
	ScoreID      string
	AchievedDate string
	PP           float32
	scoreBase
}

// RecentScore represents a score in a user's recent plays.
type RecentScore struct {
	BeatmapID string
	scoreBase
}

// UserPerfQuery is a query that's used to get either a user's best scores or a user's recent
// scores.
type UserPerfQuery struct {
	User  string
	Mode  GameMode
	Limit int
	Type  UserType
}
