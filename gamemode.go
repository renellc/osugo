package osugo

// GameMode represents the game mode for a given score, beatmap, or query.
type GameMode int

const (
	// Osu represents the osu! game mode.
	Osu GameMode = iota
	// Taiko represents the Taiko game mode.
	Taiko
	// CtB represents the Catch the Beat game mode.
	CtB
	// Mania represents the osu!mania game mode.
	Mania
	// Any represents all of the game modes.
	Any
)

// GetName gets the string representation for a GameMode.
func (m GameMode) GetName() string {
	modes := []string{"osu!", "Taiko", "Catch the Beat", "osu!mania"}
	return modes[m]
}
