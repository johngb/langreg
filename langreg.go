package langreg

// IsValidLangRegCode returns true if the string s is a valid ISO 639-1 language
// and ISO1366-1_alpa-2 region code separated by an underscore.  E.g. "en_US".
func IsValidLangRegCode(s string) bool {

	// all valid codes are 5 characters long
	if len(s) != 5 {
		return false
	}

	// the middle (third) character must be a '_' uint8 char
	if s[2] != '_' {
		return false
	}

	// check the language code, which should be the first two characters in s
	if !IsValidLanguageCode(s[:2]) {
		return false
	}

	// check the region code, which should be the last two characters in s
	if !IsValidRegionCode(s[3:]) {
		return false
	}
	return true
}
