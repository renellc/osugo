package osugof

// Language is an ID that represents a given language.
type Language int

const (
	AnyLanguage Language = iota
	OtherLanguage
	English
	Japanese
	Chinese
	Instrumental
	Korean
	French
	German
	Swedish
	Spanish
	Italian
)

// GetName gets the string representation for a Language.
func (l Language) GetName() string {
	langs := []string{
		"Any Language",
		"Other Language",
		"English",
		"Japanese",
		"Chinese",
		"Instrumental",
		"Korean",
		"French",
		"German",
		"Swedish",
		"Spanish",
		"Italian",
	}
	return langs[l]
}
