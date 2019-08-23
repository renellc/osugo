package osugo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
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

// GetUser gets the osu! data for a specified user.
func (c OsuClient) GetUser(q UserQuery) (*User, error) {
	res, err := c.sendRequest("get_user", q)
	if err != nil {
		return nil, err
	}

	var users []User
	jErr := json.Unmarshal(res, &users)
	if jErr != nil {
		return nil, jErr
	}

	// get_users returns an empty array when the User value provided in the query is not a valid
	// user or if the Type value given conflicts with the value given in User. In this case,
	// return nil for the *User value
	if len(users) == 0 {
		return nil, nil
	}

	// get_users returns an array, but the array will always have one result so we just return the
	// first result in the array.
	return &users[0], nil
}

// Print prints to the console the user's information in a digestible format.
func (u User) Print() {
	fmt.Printf("User\n    Name: %s\n    ID: %s\n", u.Username, u.UserID)
	fmt.Printf("Joined: %s\n", u.JoinDate)
	fmt.Printf("Hits\n    300: %d\n    100: %d\n    50: %d\n", u.Hits300, u.Hits100, u.Hits50)
	fmt.Printf("Play Count: %d\n", u.PlayCount)
	fmt.Printf("Score\n    Ranked: %d\n    Total: %d\n", u.RankedScore, u.TotalScore)
	fmt.Printf("Rank: %d\nPP: %f\nLevel: %f\n", u.Rank, u.PP, u.Level)
	fmt.Printf("Accuracy: %f\n", u.Accuracy)
	fmt.Printf("Letter Count\n    SS/SSH: %d/%d\n    S/SH: %d/%d\n    A: %d\n",
		u.RankSSCount, u.RankSSHCount, u.RankSCount, u.RankSHCount, u.RankACount)
	fmt.Printf("Country\n    Location: %s\n    Rank: %d\n", u.Country, u.CountryRank)
	fmt.Printf("Total Seconds Played: %d\n", u.TotalSecondsPlayed)
}

// UserQuery is used to fetch user information.
type UserQuery struct {
	// REQUIRED - The user to get information from.
	User string
	// OPTIONAL - Specifies which game mode information to get from a user.
	Mode GameMode
	// OPTIONAL - Specifies the type of value (username or user ID) passed into the User field.
	Type UserType
	// OPTIONAL - Specifies the maximum number of days between now and the last event date.
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
	reqURL.Add("m", strconv.Itoa(int(u.Mode)))

	if u.Type != "" {
		reqURL.Add("type", string(u.Type))
	}

	if u.EventDays > 0 {
		reqURL.Add("event_days", strconv.Itoa(u.EventDays))
	}

	return reqURL.Encode(), nil
}

func (u UserQuery) validateQuery() error {
	var err error

	if u.User == "" {
		err = errors.New("UserQuery: User value must be provided")
	}

	if u.Mode > 3 {
		err = errors.New("UserQuery: GameMode provided is not supported by this query")
	}

	if u.EventDays < 0 || u.EventDays > 31 {
		err = errors.New("UserQuery: EventDays value not valid. Either leave EventDays blank or" +
			" set a value between 1 and 31")
	}

	return err
}
