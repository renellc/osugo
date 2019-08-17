package osugo

type Mod int

const (
	ModNone   Mod = 0
	ModNoFail Mod = 1 << iota
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
	ModKeyMod            Mod = ModKey1 | ModKey2 | ModKey3 | ModKey4 | ModKey5 | ModKey6 | ModKey7 | ModKey8 | ModKey9 | ModKeyCoop
	ModFreeModAllowed    Mod = ModNoFail | ModEasy | ModHidden | ModHardRock | ModSuddenDeath | ModFlashlight | ModFadeIn | ModRelax | ModRelax2 | ModSpunOut | ModKeyMod
	ModScoreIncreaseMods Mod = ModHidden | ModHardRock | ModDoubleTime | ModFlashlight | ModFadeIn
)

// Getting the string representation of the Mods will be done somewhere else since scores can have
// more than 1 mod applied to it.
