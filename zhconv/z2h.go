package zhconv

import (
	"strings"

	"github.com/suwakei/go-zhconv/tables"
)

// Z2h converts full-width characters (zenkaku) in a string to half-width characters (hankaku).
// It handles ASCII, Katakana, digits, and Katakana with dakuten/handakuten.
func Z2h(str string) string {
	if str == "" {
		return ""
	}
	var result strings.Builder
	result.Grow(len(str))

	t := tables.New()

	for _, char := range str {
		if c, ok := t.ASCII_Z2H_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else if c, ok := t.DIGIT_Z2H_CHARS_MAP[char]; ok {
			result.WriteRune(c)

			// Check for Katakana with dakuten (e.g., 'ガ')
			// Assumes KANA_DAKUTEN_MAP maps 'ガ' -> 'カ'
		} else if baseKana, ok := t.KANA_Z2H_DAKUTEN_MAP[char]; ok {
			// Convert the base Z2H kana ('カ') to hankaku kana ('ｶ')
			if c, ok := t.KANA_Z2H_CHARS_MAP[baseKana]; ok {
				result.WriteRune(c)   // Write 'ｶ'
				result.WriteRune('ﾞ') // Write hankaku dakuten 'ﾞ'
			} else {
				// Base kana not found in KANA_Z2H_CHARS_MAP Should not happen if tables are consistent.
				result.WriteRune(char) // write original character
			}

			// Check for Katakana with handakuten (e.g., 'パ')
			// Assumes KANA_MARU_MAP maps 'パ' -> 'ハ'
		} else if baseKana, ok := t.KANA_Z2H_MARU_MAP[char]; ok {
			// Convert the base Z2H kana ('ハ') to hankaku kana ('ﾊ')
			if c, ok := t.KANA_Z2H_CHARS_MAP[baseKana]; ok {
				result.WriteRune(c)   // Write 'ﾊ'
				result.WriteRune('ﾟ') // Write hankaku handakuten 'ﾟ'
			} else {
				// Base kana not found in KANA_Z2H_CHARS_MAP Should not happen if tables are consistent.
				result.WriteRune(char) // Fallback: write original character
			}
			// Check for standard Katakana (without dakuten/handakuten)
		} else if c, ok := t.KANA_Z2H_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else {
			// Character is not convertible (e.g., Hiragana, Kanji, symbols not covered)
			result.WriteRune(char)
		}
	}
	return result.String()
}
