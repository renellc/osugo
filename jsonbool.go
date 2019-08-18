package osugo

// JSONBool represents a boolean value in an osu! API response. 1 and 0 were chosen for the true
// and false values respectively, so this type is used to unmarshal those values into Go's bool
// data type.
type JSONBool bool

func (b *JSONBool) UnmarshalJSON(data []byte) error {
	s := string(data)
	*b = false

	if s == "1" {
		*b = true
	}

	return nil
}
