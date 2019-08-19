package osugo

import (
	"encoding/json"
	"errors"
	"net/url"
)

// MultiLobby represents a multiplayer lobby in osu!. This contains the match meta information, as
// well as the multiple games that are played within that lobby.
type MultiLobby struct {
	Info  MatchInfo   `json:"match"`
	Games []MultiGame `json:"games"`
}

// GetMultiMatch gets information about a multiplayer match in osu!.
func (c OsuClient) GetMultiMatch(q MultiLobbyQuery) (*MultiLobby, error) {
	res, err := c.sendRequest("get_match", q)
	if err != nil {
		return nil, err
	}

	lobby := MultiLobby{}
	jErr := json.Unmarshal(res, &lobby)
	if jErr != nil {
		return nil, jErr
	}

	return &lobby, nil
}

// MatchInfo contains the meta information for a multiplayer lobby in osu!.
type MatchInfo struct {
	MatchID   string `json:"match_id"`
	LobbyName string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// MultiGame contains the information for a single game that is played within a multiplayer lobby
// in osu!. This means it containsinformation for the game as well as all the scores set by each of
// the players in the lobby.
type MultiGame struct {
	GameID      string            `json:"game_id"`
	StartTime   string            `json:"start_time"`
	EndTime     string            `json:"end_time"`
	BeatmapID   string            `json:"beatmap_id"`
	Mode        GameMode          `json:"play_mode,string"`
	MatchType   string            `json:"match_type"` // What is this? Not documented in osu! API wiki.
	WinCriteria int               `json:"scoring_type,string"`
	TeamType    int               `json:"team_type,string"`
	Mods        int               `json:"mods,string"`
	Scores      []MultiMatchScore `json:"scores"`
}

// GetWinCriteriaName gets the string value for a win criteria type.
func (m MultiGame) GetWinCriteriaName() string {
	crit := []string{"Score", "Accuracy", "Combo", "ScoreV2"}
	return crit[m.WinCriteria]
}

// GetTeamTypeName gets the string value for the type of match played in a multiplayer lobby.
func (m MultiGame) GetTeamTypeName() string {
	types := []string{"Head to Head", "Tag Co-op", "Team VS", "Tag Team VS"}
	return types[m.TeamType]
}

// MultiMatchScore contains the information for a score set in a multiplayer lobby in osu!.
type MultiMatchScore struct {
	LobbySlot int      `json:"slot,string"`
	LobbyTeam int      `json:"team,string"`
	DidPass   JSONBool `json:"pass,string"`
	Score
}

// MultiLobbyQuery contains the various parameters used to get the data for a multiplayer lobby in
// osu!.
type MultiLobbyQuery struct {
	LobbyID string
}

func (m MultiLobbyQuery) constructQuery(key string) (string, error) {
	validateErr := m.validateQuery()
	if validateErr != nil {
		return "", validateErr
	}

	reqURL := url.Values{}
	reqURL.Add("k", key)
	reqURL.Add("mp", m.LobbyID)
	return reqURL.Encode(), nil
}

func (m MultiLobbyQuery) validateQuery() error {
	if m.LobbyID == "" {
		return errors.New("No match ID value provided")
	}
	return nil
}
