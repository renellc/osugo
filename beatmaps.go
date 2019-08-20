package osugo

type Beatmap struct {
	Status              int      `json:"approved,string"`
	Submitted           string   `json:"submit_date"`
	Approved            string   `json:"approved_date"`
	LastUpdated         string   `json:"last_update"`
	Artist              string   `json:"artist"`
	BeatmapID           string   `json:"beatmap_id"`
	BeatmapSetID        string   `json:"beatmapset_id"`
	BPM                 float32  `json:"json:bpm,string"`
	Creator             string   `json:"creator"`
	CreatorID           string   `json:"creator_id"`
	RatingStar          float32  `json:"difficultyrating,string"`
	RatingAim           float32  `json:"diff_aim,string"`
	RatingSpeed         float32  `json:"diff_speed,string"`
	RatingCircleSize    float32  `json:"diff_size,string"`
	RatingOverall       float32  `json:"diff_overall,string"`
	RatingHPDrain       float32  `json:"diff_drain,string"`
	HitLength           int      `json:"hit_length,string"`
	SongSource          string   `json:"source"`
	Genre               Genre    `json:"genre_id,string"`
	Language            Language `json:"language_id,string"`
	Title               string   `json:"title"`
	MapLength           int      `json:"total_length,string"`
	DifficultyName      string   `json:"version"`
	FileMD5             string   `json:"file_md5"`
	Mode                GameMode `json:"mode,string"`
	Tags                Tags     `json:"tags,string"`
	FavoriteCount       int      `json:"favourite_count,string"`
	UserRating          float32  `json:"rating,string"`
	PlayCount           int      `json:"playcount,string"`
	PassCount           int      `json:"passcount,string"`
	NoteCount           int      `json:"count_normal,string"`
	SliderCount         int      `json:"count_slider,string"`
	SpinnerCount        int      `json:"count_spinner,string"`
	MaxCombo            int      `json:"max_combo,string"`
	DownloadUnavailable JSONBool `json:"download_unavailable,string"`
	AudioUnavailable    JSONBool `json:"audio_unavailable,string"`
}

func (b Beatmap) GetStatusName() string {
	names := []string{"Graveyard", "WIP", "Pending", "Ranked", "Approved", "Qualified", "Loved"}
	// Graveyard starts at -2, so add 2 to get the slice index.
	return names[b.Status+2]
}
