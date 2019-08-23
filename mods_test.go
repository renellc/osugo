package osugo_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/renellc/osugo"
)

func TestUnmarshalJSONIntoMods(t *testing.T) {
	// Setup test cases
	table := []struct {
		Input  string
		Output osugo.Mods
	}{
		{
			Input:  fmt.Sprintf("{\"enabled_mods\": \"%d\"}", osugo.ModHardRock|osugo.ModHidden),
			Output: osugo.Mods{"Hard Rock", "Hidden"},
		},
		{
			Input: fmt.Sprintf("{\"enabled_mods\": \"%d\"}",
				osugo.ModDoubleTime|osugo.ModNoFail|osugo.ModFlashlight),
			Output: osugo.Mods{"No Fail", "Double Time", "Flashlight"},
		},
	}

	for _, test := range table {
		// Convert to bytes so we can unmarshal
		data := []byte(test.Input)
		out := struct {
			EnabledMods osugo.Mods `json:"enabled_mods,string"`
		}{}
		err := json.Unmarshal(data, &out)
		if err != nil {
			t.Error(err)
		}

		// Checks if the mods generated match the desired output for this test case.
		m := createModMap(test.Output)
		for _, mod := range out.EnabledMods {
			if !m[mod] {
				t.Error("Generated mod that does not match the desired output")
			}
		}
	}
}

func createModMap(mods osugo.Mods) map[string]bool {
	m := make(map[string]bool)
	for _, mod := range mods {
		m[mod] = true
	}
	return m
}
