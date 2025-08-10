package utils

func IsAlphabeticOrUnderscore(ch byte) bool {
	return IsAlphabetic(ch) || IsUnderscore(ch)
}
