package osugo

// Genre represents the genre a given song is categorized under.
type Genre int

const (
	Any Genre = iota
	Unspecified
	VideoGame
	Anime
	Rock
	Pop
	Other
	Novelty
	HipHop Genre = iota + 1 // For some reason 8 was skipped. ??
	Electronic
)

// GetName gets the string representation for a Genre.
func (g Genre) GetName() string {
	genres := []string{
		"Any",
		"Unspecified",
		"Video Game",
		"Anime",
		"Rock",
		"Pop",
		"Other",
		"Novelty",
		"",
		"Hip Hop",
		"Electronic",
	}
	return genres[g]
}
