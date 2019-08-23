package osugo

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// RecentScore represents a score in a user's recent plays.
type RecentScore struct {
	BeatmapID string
	Score
}

// GetUserRecent gets the data on a specified user's recent plays.
func (c OsuClient) GetUserRecent(q UserRecentQuery) ([]RecentScore, error) {
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

// UserRecentQuery is a query that's used to get either a user's best scores or a user's recent
// scores.
type UserRecentQuery struct {
	// REQUIRED - Specifies a username or user ID to get plays from
	User string
	// OPTIONAL - The game mode. Defaults to osu!.
	Mode GameMode
	// OPTIONAL - The amount of results. Defaults to 10.
	Limit int
	// OPTIONAL - Specifies whether the value in `User` is a Username or and ID.
	Type UserType
}

func (ur UserRecentQuery) constructQuery(key string) (string, error) {
	validateErr := ur.validateQuery()
	if validateErr != nil {
		return "", validateErr
	}

	reqURL := url.Values{}
	reqURL.Add("k", key)
	reqURL.Add("u", ur.User)
	reqURL.Add("m", string(ur.Mode))

	if ur.Limit > 0 {
		reqURL.Add("limit", strconv.Itoa(ur.Limit))
	}

	if ur.Type != "" {
		reqURL.Add("type", string(ur.Type))
	}

	return reqURL.Encode(), nil
}

func (ur UserRecentQuery) validateQuery() error {
	var err error

	if ur.User == "" {
		err = errors.New("UserRecentQuery: No User value provided")
	}

	if ur.Mode > 3 {
		err = errors.New("UserRecentQuery: GameMode provided is not supported by this query")
	}

	if ur.Limit < 0 || ur.Limit > 50 {
		err = errors.New("UserRecentQuery: Limit value must be between 1 and 50")
	}

	return err
}
