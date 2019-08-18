package osugo

import (
	"encoding/json"
	"errors"
	"net/url"
)

// Score is a struct that contains osu! data for a given score.
type Score struct {
	ScoreID         string   `json:"score_id"`
	Username        string   `json:"username"`
	AchievedDate    string   `json:"date"`
	PP              float32  `json:"pp,string"`
	ReplayAvailable JSONBool `json:"replay_available,string"`
	ScoreBase
}

// ScoresQuery is used to fetch the scores set for a specified beatmap.
type ScoresQuery struct {
	BeatmapID string
	User      string
	Mode      GameMode
	Mods      int
	Type      UserType
	Limit     int
}

// GetScores gets a list of scores for a specified beatmap.
func (c OsuClient) GetScores(q ScoresQuery) ([]Score, error) {
	res, err := c.sendRequest("get_scores", q)
	if err != nil {
		return nil, err
	}

	scores := []Score{}
	jErr := json.Unmarshal(res, &scores)
	if jErr != nil {
		return nil, jErr
	}

	return scores, nil
}

func (s ScoresQuery) constructQuery(key string) (string, error) {
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
		reqURL.Add("m", string(s.Mode))
	}

	if s.Type != "" {
		reqURL.Add("type", string(s.Type))
	}

	if s.Limit > 0 {
		reqURL.Add("limit", string(s.Limit))
	}

	return reqURL.Encode(), nil
}

func (s ScoresQuery) validateQuery() error {
	var err error

	if s.BeatmapID == "" {
		err = errors.New("No BeatmapID value provided")
	}

	if s.Mode > 3 {
		err = errors.New("GameMode provided is not supported by this query")
	}

	if s.Limit < 0 || s.Limit > 100 {
		err = errors.New("Limit value provided is invalid. Either leave the Limit field blank or " +
			"enter a value between 1 and 100")
	}

	return err
}
