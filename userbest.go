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
func (c OsuClient) GetUserBest(q UserRecentQuery) ([]BestScore, error) {
	data, err := c.sendRequest("get_user_best", q)
	if err != nil {
		return nil, err
	}

	var scores []BestScore
	jErr := json.Unmarshal(data, &scores)
	if jErr != nil {
		return nil, jErr
	}

	return scores, nil
}

// UserBestQuery is a struct that contains various parameters that are applied to a GET request to
// get a user's best scores.
type UserBestQuery struct {
	// REQUIRED - Specifies a username or user ID to get plays from
	User string
	// OPTIONAL - The game mode. Defaults to osu!.
	Mode GameMode
	// OPTIONAL - The amount of results. Defaults to 10.
	Limit int
	// OPTIONAL - Specifies whether the value in `User` is a Username or and ID.
	Type UserType
}

func (ub UserBestQuery) constructQuery(key string) (string, error) {
	validateErr := ub.validateQuery()
	if validateErr != nil {
		return "", validateErr
	}

	reqURL := url.Values{}
	reqURL.Add("k", key)
	reqURL.Add("u", ub.User)
	reqURL.Add("m", string(ub.Mode))

	if ub.Limit > 0 {
		reqURL.Add("limit", strconv.Itoa(ub.Limit))
	}

	if ub.Type != "" {
		reqURL.Add("type", string(ub.Type))
	}

	return reqURL.Encode(), nil
}

func (ub UserBestQuery) validateQuery() error {
	var err error

	if ub.User == "" {
		err = errors.New("UserBestQuery: No User value provided")
	}

	if ub.Mode > 3 {
		err = errors.New("UserBestQuery: GameMode provided is not supported by this query")
	}

	if ub.Limit < 0 || ub.Limit > 50 {
		err = errors.New("UserBestQuery: Limit value must be between 1 and 50")
	}

	return err
}
