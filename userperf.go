package osugo

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// BestScore represents a score in a user's top plays.
type BestScore struct {
	BeatmapID    string  `json:"beatmap_id"`
	ScoreID      string  `json:"score_id"`
	AchievedDate string  `json:"date"`
	PP           float32 `json:"pp,string"`
	Score
}

// GetUserBest gets the data on a specified user's top plays.
func (c OsuClient) GetUserBest(q UserPerfQuery) ([]BestScore, error) {
	data, err := c.sendRequest("get_user_best", q)
	if err != nil {
		return nil, err
	}

	scores := []BestScore{}
	jErr := json.Unmarshal(data, &scores)
	if jErr != nil {
		return nil, jErr
	}

	return scores, nil
}

// RecentScore represents a score in a user's recent plays.
type RecentScore struct {
	BeatmapID string
	Score
}

// GetUserRecent gets the data on a specified user's recent plays.
func (c OsuClient) GetUserRecent(q UserPerfQuery) ([]RecentScore, error) {
	data, err := c.sendRequest("get_user_recent", q)
	if err != nil {
		return nil, err
	}

	scores := []RecentScore{}
	jErr := json.Unmarshal(data, &scores)
	if jErr != nil {
		return nil, jErr
	}

	return scores, nil
}

// UserPerfQuery is a query that's used to get either a user's best scores or a user's recent
// scores.
type UserPerfQuery struct {
	// REQUIRED - Specifies a username or user ID to get plays from
	User string
	// OPTIONAL - The game mode. Defaults to osu!.
	Mode GameMode
	// OPTIONAL - The amount of results. Defaults to 10.
	Limit int
	// OPTIONAL - Specifies whether the value in `User` is a Username or and ID.
	Type UserType
}

func (upq UserPerfQuery) constructQuery(key string) (string, error) {
	validateErr := upq.validateQuery()
	if validateErr != nil {
		return "", validateErr
	}

	reqURL := url.Values{}
	reqURL.Add("k", key)
	reqURL.Add("u", upq.User)
	reqURL.Add("m", string(upq.Mode))

	if upq.Limit > 0 {
		reqURL.Add("limit", strconv.Itoa(upq.Limit))
	}

	if upq.Type != "" {
		reqURL.Add("type", string(upq.Type))
	}

	return reqURL.Encode(), nil
}

func (upq UserPerfQuery) validateQuery() error {
	var err error

	if upq.User == "" {
		err = errors.New("No User value provided")
	}

	if upq.Mode > 3 {
		err = errors.New("GameMode provided is not supported by this query")
	}

	if upq.Limit < 0 || upq.Limit > 50 {
		err = errors.New("Limit value must be between 1 and 50")
	}

	return err
}
