package osugo

// scoreBase is used internally amongst the Score, RecentScore, BestScore, and MultiScore structs.
// Because all of the response from get_user_best, get_user_recent, get_match, and get_scores share
// virtually the same fields for the score, this struct was created to be used amongts all of them.
type scoreBase struct {
	UserID      string `json:"user_id"`
	Score       int    `json:"score,string"`
	MaxCombo    int    `json:"maxcombo,string"`
	Rank        string `json:"rank"`
	HitsGeki    int    `json:"countgeki,string"`
	HitsKatu    int    `json:"countkatu,string"`
	Hits300     int    `json:"count300,string"`
	Hits100     int    `json:"count100,string"`
	Hits50      int    `json:"count50,string"`
	HitsMiss    int    `json:"countmiss,string"`
	FullCombo   string `json:"perfect"`
	EnabledMods string `json:"enabled_mods"`
}