package zhconv

import (
	"strings"
)

// Z2h converts full-width characters (zenkaku) in a string to half-width characters (hankaku).
// It handles ASCII, Katakana, digits, and Katakana with dakuten/handakuten.
func Z2h(str string) string {
	if str == "" {
		return ""
	}
	var result strings.Builder
	result.Grow(len(str))

	for _, char := range str {
		if c, ok := convTables.ASCII_Z2H_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else if c, ok := convTables.DIGIT_Z2H_CHARS_MAP[char]; ok {
			result.WriteRune(c)

			// Check for Katakana with dakuten (e.g., 'ガ')
			// Assumes KANA_DAKUTEN_MAP maps 'ガ' -> 'カ'
		} else if baseKana, ok := convTables.KANA_Z2H_DAKUTEN_MAP[char]; ok {
			result.WriteRune(baseKana)
			result.WriteRune('ﾞ')

			// Check for Katakana with handakuten (e.g., 'パ')
			// Assumes KANA_MARU_MAP maps 'パ' -> 'ハ'
		} else if baseKana, ok := convTables.KANA_Z2H_MARU_MAP[char]; ok {
			result.WriteRune(baseKana) // Write the hankaku base kana (e.g., 'ﾊ')
			result.WriteRune('ﾟ')     // Write hankaku handakuten 'ﾟ'
			// Check for standard Katakana (without dakuten/handakuten)
		} else if c, ok := convTables.KANA_Z2H_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else {
			// Character is not convertible (e.g., Hiragana, Kanji, symbols not covered)
			result.WriteRune(char)
		}
	}
	return result.String()
}
