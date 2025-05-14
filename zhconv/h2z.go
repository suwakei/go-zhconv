package zhconv

import (
	"strings"
)

// H2z converts half-width characters (hankaku) in a string to full-width characters (zenkaku).
// It handles ASCII, Katakana, digits, and Katakana with dakuten/handakuten.
func H2z(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(str)

	i := 0
	runeLen := len(runes)

	var result strings.Builder
	result.Grow(runeLen * 3) // "* 3" is for corresponding to multi bytes capacity.

	for i < runeLen {
		char := runes[i]

		if i+1 < runeLen {
			nextChar := runes[i+1]

			if nextChar == 'ﾞ' {
				if zenkakuDakuten, ok := convTables.KANA_H2Z_DAKUTEN_MAP[char]; ok {
					result.WriteRune(zenkakuDakuten)
					result.WriteRune('ﾞ')
					i += 2   // Skip both the current character and the dakuten.
					continue
				}
			} else if nextChar == 'ﾟ' {
				if zenkakuHandakuten, ok := convTables.KANA_H2Z_MARU_MAP[char]; ok {
					result.WriteRune(zenkakuHandakuten)
					result.WriteRune('ﾟ')
					i += 2
					continue // Continue to the next iteration.
				}
			}
		}

		// If it's not a combined dakuten/handakuten case, perform standard conversions
		if c, ok := convTables.ASCII_H2Z_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else if c, ok := convTables.KANA_H2Z_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else if c, ok := convTables.DIGIT_H2Z_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else {
			// Character is not convertible or a standalone dakuten/handakuten, etc.
			result.WriteRune(char)
		}

		i++
	}
	return result.String()
}