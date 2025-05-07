package zhconv

import (
	"strings"

	"github.com/suwakei/go-zhconv/tables"
)

// H2z converts half-width characters (hankaku) in a string to full-width characters (zenkaku).
// It handles ASCII, Katakana, digits, and Katakana with dakuten/handakuten.
// Assumes tables.KANA_HANKAKU_DAKUTEN_MAP and tables.KANA_HANKAKU_HANDAKUTEN_MAP are defined.
func H2z(str string) string {
	if str == "" {
		return ""
	}
	var result strings.Builder
	result.Grow(len(str) * 3) // "*3" is for corresponding to multi bytes capacity

	t := tables.New()    // Get conversion tables
	runes := []rune(str) // Convert string to rune slice for correct multi-byte character handling

	i := 0
	runeLen := len(runes)

	for i < runeLen {
		char := runes[i]

		// Check for potential dakuten (ﾞ) or handakuten (ﾟ) combination
		// Check if the next character existence and dakuten or handakuten
		if i+1 < runeLen {
			nextChar := runes[i+1]
			if nextChar == 'ﾞ' {
				// Check if the current character is a hankaku kana that can take a dakuten
				if zenkakuDakuten, ok := t.KANA_HANKAKU_DAKUTEN_MAP[char]; ok {
					result.WriteRune(zenkakuDakuten)
					i += 2   // Skip both the current character and the dakuten
					continue // Continue to the next iteration
				}
			} else if nextChar == 'ﾟ' {
				// Check if the current character is a hankaku kana that can take a handakuten
				if zenkakuHandakuten, ok := t.KANA_HANKAKU_MARU_MAP[char]; ok {
					result.WriteRune(zenkakuHandakuten)
					i += 2   // Skip both the current character and the handakuten
					continue // Continue to the next iteration
				}
			}
		}

		// If it's not a combined dakuten/handakuten case, perform standard conversions
		if idx := indexRune(t.ASCII_HANKAKU_CHARS, char); idx != -1 {
			result.WriteRune(t.ASCII_ZENKAKU_CHARS[idx])
		} else if idx := indexRune(t.KANA_HANKAKU_CHARS, char); idx != -1 {
			// This handles hankaku kana that are *not* followed by a dakuten/handakuten
			result.WriteRune(t.KANA_ZENKAKU_CHARS[idx])
		} else if idx := indexRune(t.DIGIT_HANKAKU_CHARS, char); idx != -1 {
			result.WriteRune(t.DIGIT_ZENKAKU_CHARS[idx])
		} else {
			// Character is not convertible or is a standalone dakuten/handakuten, etc.
			result.WriteRune(char)
		}

		i++ // Move to the next character in the rune slice
	}
	return result.String() // Return the built string
}