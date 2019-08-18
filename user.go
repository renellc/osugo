package osugo

import (
	"errors"
	"net/url"
)

// Event is a struct that represents a recent event the user has done. This is related to a user
// setting a new top play (and maybe unlocking achievements?).
type Event struct {
	DisplayHTML  string `json:"display_html"`
	BeatmapID    string `json:"beatmap_id"`
	BeatmapSetID string `json:"beatmapset_id"`
	EventDate    string `json:"date"`
	EpicFactor   int    `json:"epicfactor,string"`
}

// User is a struct that contains the osu! profile data for a user.
type User struct {
	UserID             string  `json:"user_id"`
	Username           string  `json:"username"`
	JoinDate           string  `json:"join_date"`
	Hits300            uint    `json:"count300,string"`
	Hits100            uint    `json:"count100,string"`
	Hits50             uint    `json:"count50,string"`
	PlayCount          uint    `json:"playcount,string"`
	RankedScore        uint64  `json:"ranked_score,string"`
	TotalScore         uint64  `json:"total_score,string"`
	Rank               uint    `json:"pp_rank,string"`
	Level              float32 `json:"level,string"`
	PP                 float32 `json:"pp_raw,string"`
	Accuracy           float32 `json:"accuracy,string"`
	RankSSCount        uint    `json:"count_rank_ss,string"`
	RankSSHCount       uint    `json:"count_rank_ssh,string"`
	RankSCount         uint    `json:"count_rank_s,string"`
	RankSHCount        uint    `json:"count_rank_sh,string"`
	RankACount         uint    `json:"count_rank_a,string"`
	Country            string  `json:"country"`
	TotalSecondsPlayed uint    `json:"total_seconds_played,string"`
	CountryRank        uint    `json:"pp_country_rank,string"`
	Events             []Event `json:"events"`
}

// UserQuery is used to fetch user information.
type UserQuery struct {
	User      string
	Mode      GameMode
	Type      UserType
	EventDays int
}

func (u UserQuery) constructQuery(key string) (string, error) {
	validateErr := u.validateQuery()
	if validateErr != nil {
		return "", nil
	}

	reqURL := url.Values{}
	reqURL.Add("k", key)
	reqURL.Add("u", u.User)
	reqURL.Add("m", string(u.Mode))

	if u.Type != "" {
		reqURL.Add("type", string(u.Type))
	}

	if u.EventDays == 0 {
		reqURL.Add("event_days", string(1))
	} else {
		reqURL.Add("event_days", string(u.EventDays))
	}

	return reqURL.Encode(), nil
}

func (u UserQuery) validateQuery() error {
	var err error

	if u.User == "" {
		err = errors.New("User value must be provided")
	}

	if u.Mode > 3 {
		err = errors.New("GameMode provided is not supported by this query")
	}

	if u.EventDays < 0 || u.EventDays > 31 {
		err = errors.New("EventDays value not valid. Either leave EventDays blank or set a value" +
			"between 1 and 31")
	}

	return err
}
