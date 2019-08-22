package osugo

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"
)

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

func (c OsuClient) GetBeatmaps(q BeatmapQuery) ([]Beatmap, error) {
	res, err := c.sendRequest("get_beatmaps", q)
	if err != nil {
		return nil, err
	}

	maps := []Beatmap{}
	jErr := json.Unmarshal(res, &maps)
	if jErr != nil {
		return nil, jErr
	}

	return maps, nil
}

// BeatmapQuery is used to get a list of beatmaps.
type BeatmapQuery struct {
	Since            time.Time
	BeatmapSetID     string
	BeatmapID        string
	User             string
	Type             UserType
	Mode             GameMode
	IncludeConverted bool
	BeatmapHash      string
	Limit            int
	// Not entirely sure what this parameter is for... Tested a couple of calls to API adding a
	// value to this parameter and it didn't change the response. If someone knows what this does,
	// please tell me lol
	Mods int
}

func (b BeatmapQuery) constructQuery(key string) (string, error) {
	validateErr := b.validateQuery()
	if validateErr != nil {
		return "", nil
	}

	reqURL := url.Values{}
	reqURL.Add("k", key)

	if !b.Since.IsZero() {
		since := b.Since.Format("2006-01-22 21:30:00")
		reqURL.Add("since", since)
	}

	if b.BeatmapSetID != "" {
		reqURL.Add("s", b.BeatmapSetID)
	}

	if b.BeatmapID != "" {
		reqURL.Add("b", b.BeatmapID)
	}

	if b.User != "" {
		reqURL.Add("u", b.User)
	}

	if b.Type != "" {
		reqURL.Add("type", string(b.Type))
	}

	if b.Mode != Any {
		reqURL.Add("m", string(b.Mode))
	}

	if b.IncludeConverted {
		reqURL.Add("a", "true")
	}

	if b.BeatmapHash != "" {
		reqURL.Add("h", b.BeatmapHash)
	}

	if b.Limit > 0 {
		reqURL.Add("limit", string(b.Limit))
	}

	if b.Mods > 0 {
		reqURL.Add("mods", string(b.Mods))
	}

	return reqURL.Encode(), nil
}

func (b BeatmapQuery) validateQuery() error {
	var err error
	if b.Limit < 1 || b.Limit > 500 {
		err = errors.New("BeatmapQuery: limit value provided is invalid. " +
			"Don't set value or set a value between 1 and 500")
	}
	return err
}