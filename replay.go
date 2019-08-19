package osugo

import (
	"encoding/json"
	"errors"
	"net/url"
)

// Replay is a struct that contains the data of an osu! replay file.
type Replay struct {
	Encoding string `json:"encoding"`
	Content  string `json:"content"`
}

// ReplayQuery is used to fetch the replay data of a user's score.
type ReplayQuery struct {
	// REQUIRED - The game mode the score was set on.
	Mode GameMode
	// REQUIRED - The beatmap ID from which the score was set.
	BeatmapID string
	// REQUIRED - The user that set the score.
	User string
	// OPTIONAL - Specifies the type of value (username or user ID) passed into the User field.
	Type UserType
	// OPTIONAL - The score to get with specific mods.
	Mods int
}

// GetReplay gets the replay data for a specified user on a specified beatmap.
func (c OsuClient) GetReplay(q ReplayQuery) (*Replay, error) {
	data, err := c.sendRequest("get_replay", q)
	if err != nil {
		return nil, err
	}

	replay := Replay{}
	jErr := json.Unmarshal(data, &replay)
	if jErr != nil {
		return nil, jErr
	}

	return &replay, nil
}

func (r ReplayQuery) constructQuery(key string) (string, error) {
	validErr := r.validateQuery()
	if validErr != nil {
		return "", validErr
	}

	q := url.Values{}
	q.Add("k", key)
	q.Add("m", string(r.Mode))
	q.Add("b", r.BeatmapID)
	q.Add("u", r.User)

	if r.Type != "" {
		q.Add("type", string(r.Type))
	}

	return q.Encode(), nil
}

func (r ReplayQuery) validateQuery() error {
	var err error

	if r.Mode > Mania {
		err = errors.New("GameMode provided is not supported by this query")
	}

	if r.BeatmapID == "" {
		err = errors.New("No BeatmapID value was provided")
	}

	if r.User == "" {
		err = errors.New("No user value was provided")
	}

	return err
}
