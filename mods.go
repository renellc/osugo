package osugo

import (
	"strconv"
)

type ModValue int

const (
	ModNone   ModValue = 0
	ModNoFail ModValue = 1 << iota
	ModEasy
	ModTouchDevice
	ModHidden
	ModHardRock
	ModSuddenDeath
	ModDoubleTime
	ModRelax
	ModHalfTime
	ModNightcore
	ModFlashlight
	ModAutoplay
	ModSpunOut
	ModRelax2
	ModPerfect
	ModKey4
	ModKey5
	ModKey6
	ModKey7
	ModKey8
	ModFadeIn
	ModRandom
	ModCinema
	ModTarget
	ModKey9
	ModKeyCoop
	ModKey1
	ModKey3
	ModKey2
	ModScoreV2
	ModLastMod
	ModKeyMod            ModValue = ModKey1 | ModKey2 | ModKey3 | ModKey4 | ModKey5 | ModKey6 | ModKey7 | ModKey8 | ModKey9 | ModKeyCoop
	ModFreeModAllowed    ModValue = ModNoFail | ModEasy | ModHidden | ModHardRock | ModSuddenDeath | ModFlashlight | ModFadeIn | ModRelax | ModRelax2 | ModSpunOut | ModKeyMod
	ModScoreIncreaseMods ModValue = ModHidden | ModHardRock | ModDoubleTime | ModFlashlight | ModFadeIn
)

// Mods represents a slice holding the string representations for mods applied to a score.
type Mods []string

var modNames = []string{
	"No Fail", "Easy", "Touch Device", "Hidden", "Hard Rock", "Sudden Death", "Double Time",
	"Relax", "Half Time", "Nightcore", "Flashlight", "Autoplay", "Spun Out", "Relax2",
	"Perfect", "Key 4", "Key 5", "Key 6", "Key 7", "Key 8", "Fade In", "Random", "Cinema",
	"Target", "Key 9", "Key Coop", "Key 1", "Key 3", "Key 2", "ScoreV2", "Last Mod",
}

func (m *Mods) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	// This case is for when no mods are applied to a score.
	if s == "0" {
		*m = append(*m, "None")
		return nil
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	asBinary := strconv.FormatInt(int64(v), 2)
	// We use a for loop starting from the end of the byte slice and build up the mod names from
	// there. This is so we can ensure we get the string representation of the mod in the modNames
	// string slice.
	maxIdx := len(asBinary) - 1
	for i := maxIdx; i >= 0; i-- {
		if asBinary[i] == '1' {
			// maxIdx - i gets the index of the byte if the byte array were in reverse order.
			modName := modNames[maxIdx-i]
			*m = append(*m, modName)
		}
	}

	return nil
}
