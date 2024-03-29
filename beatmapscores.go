package osugo

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// BeatmapScore is a struct that contains osu! data for a given score.
type BeatmapScore struct {
	ScoreID         string   `json:"score_id"`
	Username        string   `json:"username"`
	AchievedDate    string   `json:"date"`
	PP              float32  `json:"pp,string"`
	ReplayAvailable JSONBool `json:"replay_available,string"`
	Score
}

// BeatmapScoresQuery is used to fetch the scores set for a specified beatmap.
type BeatmapScoresQuery struct {
	// REQUIRED - Specifies the beatmap to get scores from.
	BeatmapID string
	// OPTIONAL - Specifies a user to get score data for.
	User string
	// OPTIONAL - The game mode to get scores for.
	Mode GameMode
	// OPTIONAL - The scores to get with specific mods.
	Mods *int
	// OPTIONAL - Specifies the type of value (username or user ID) passed into the User field.
	// This is a pointer value because mods aren't applied to the GET request (with the exception)
	// of GetBeatmaps.
	Type UserType
	// OPTIONAL - The amount of scores to get.
	Limit int
}

// GetBeatmapScores gets a list of scores for a specified beatmap.
func (c OsuClient) GetBeatmapScores(q BeatmapScoresQuery) ([]BeatmapScore, error) {
	res, err := c.sendRequest("get_scores", q)
	if err != nil {
		return nil, err
	}

	var scores []BeatmapScore
	jErr := json.Unmarshal(res, &scores)
	if jErr != nil {
		return nil, jErr
	}

	return scores, nil
}

func (s BeatmapScoresQuery) constructQuery(key string) (string, error) {
	validateErr := s.validateQuery()
	if validateErr != nil {
		return "", validateErr
	}

	reqURL := url.Values{}
	reqURL.Add("k", key)
	reqURL.Add("b", s.BeatmapID)

	if s.User != "" {
		reqURL.Add("u", s.User)
	}

	if s.Mode > 0 {
		reqURL.Add("m", strconv.Itoa(int(s.Mode)))
	}

	if s.Mods != nil {
		reqURL.Add("mods", strconv.Itoa(*s.Mods))
	}

	if s.Type != "" {
		reqURL.Add("type", string(s.Type))
	}

	if s.Limit > 0 {
		reqURL.Add("limit", strconv.Itoa(s.Limit))
	}

	return reqURL.Encode(), nil
}

func (s BeatmapScoresQuery) validateQuery() error {
	var err error

	if s.BeatmapID == "" {
		err = errors.New("BeatmapScoresQuery: No BeatmapID value provided")
	}

	if s.Mode > 3 {
		err = errors.New("BeatmapScoresQuery: GameMode provided is not supported by this query")
	}

	if s.Limit < 0 || s.Limit > 100 {
		err = errors.New("BeatmapScoresQuery: Limit value provided is invalid. Either leave the Limit " +
			"field blank or enter a value between 1 and 100")
	}

	return err
}
