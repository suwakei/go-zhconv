package zhconv

import (
	"strings" // strings パッケージをインポート

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

	t := tables.New() // Get conversion tables

	for _, char := range str {
		if idx := indexRune(t.ASCII_ZENKAKU_CHARS, char); idx != -1 {
			result.WriteRune(t.ASCII_HANKAKU_CHARS[idx])
		} else if idx := indexRune(t.DIGIT_ZENKAKU_CHARS, char); idx != -1 {
			result.WriteRune(t.DIGIT_HANKAKU_CHARS[idx])

			// Check for Katakana with dakuten (e.g., 'ガ')
			// Assumes KANA_DAKUTEN_MAP maps 'ガ' -> 'カ'
		} else if baseKana, ok := t.KANA_ZENKAKU_DAKUTEN_MAP[char]; ok {
			// Convert the base zenkaku kana ('カ') to hankaku kana ('ｶ')
			if idx := indexRune(t.KANA_ZENKAKU_CHARS, baseKana); idx != -1 {
				result.WriteRune(t.KANA_HANKAKU_CHARS[idx]) // Write 'ｶ'
				result.WriteRune('ﾞ')                       // Write hankaku dakuten 'ﾞ'
			} else {
				// Base kana not found in KANA_ZENKAKU_CHARS? Should not happen if tables are consistent.
				result.WriteRune(char) // Fallback: write original character
			}

			// Check for Katakana with handakuten (e.g., 'パ')
			// Assumes KANA_MARU_MAP maps 'パ' -> 'ハ'
		} else if baseKana, ok := t.KANA_ZENKAKU_MARU_MAP[char]; ok {
			// Convert the base zenkaku kana ('ハ') to hankaku kana ('ﾊ')
			if idx := indexRune(t.KANA_ZENKAKU_CHARS, baseKana); idx != -1 {
				result.WriteRune(t.KANA_HANKAKU_CHARS[idx]) // Write 'ﾊ'
				result.WriteRune('ﾟ')                       // Write hankaku handakuten 'ﾟ'
			} else {
				// Base kana not found in KANA_ZENKAKU_CHARS? Should not happen if tables are consistent.
				result.WriteRune(char) // Fallback: write original character
			}

			// Check for standard Katakana (without dakuten/handakuten)
		} else if idx := indexRune(t.KANA_ZENKAKU_CHARS, char); idx != -1 {
			result.WriteRune(t.KANA_HANKAKU_CHARS[idx])

		} else {
			// Character is not convertible (e.g., Hiragana, Kanji, symbols not covered)
			result.WriteRune(char)
		}
	}
	return result.String() // Return the built string
}
