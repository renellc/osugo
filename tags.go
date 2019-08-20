package osugo

import "strings"

type Tags struct {
	Values []string
}

func (t *Tags) UnmarshalJSON(data []byte) error {
	str := string(data)
	t.Values = strings.Split(str, " ")
	return nil
}
