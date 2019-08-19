package osugo

// MultiLobby represents a multiplayer lobby in osu!. This contains the match meta information, as
// well as the multiple games that are played within that lobby.
type MultiLobby struct {
	Info  MatchInfo   `json:"match"`
	Games []MultiGame `json:"games"`
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

// MultiMatchScore contains the information for a score set in a multiplayer lobby in osu!.
type MultiMatchScore struct {
	LobbySlot int      `json:"slot,string"`
	LobbyTeam int      `json:"team,string"`
	DidPass   JSONBool `json:"pass,string"`
	Score
}
