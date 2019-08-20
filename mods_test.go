package osugo_test

import (
	"flag"
	"testing"

	"github.com/renellc/osugo"
)

var key string

func init() {
	flag.StringVar(&key, "key", "", "API key")
	flag.Parse()
}

func TestUnmarshalJSONIntoMods(t *testing.T) {
	c := osugo.InitClient(key)

	scores, _ := c.GetBeatmapScores(osugo.ScoresQuery{
		BeatmapID: "129891", // Freedom Dive
		User:      "124493", // Cookiezi
		Type:      osugo.ID,
	})

	m := map[string]bool{"Hidden": true, "Hard Rock": true}
	for _, mod := range scores[0].EnabledMods {
		if !m[mod] {
			t.Errorf("Mod generated does not match desired output")
		}
	}
}
