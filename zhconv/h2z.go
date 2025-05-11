package zhconv

import (
	"strings"

	"github.com/suwakei/go-zhconv/tables"
)

// H2z converts half-width characters (hankaku) in a string to full-width characters (zenkaku).
// It handles ASCII, Katakana, digits, and Katakana with dakuten/handakuten.
func H2z(str string) string {
	if str == "" {
		return ""
	}
	var result strings.Builder
	result.Grow(len(str) * 3) // "*3" is for corresponding to multi bytes capacity

	t := tables.New()

	runes := []rune(str)

	i := 0
	runeLen := len(runes)

	for i < runeLen {
		char := runes[i]

		if i+1 < runeLen {
			nextChar := runes[i+1]

			if nextChar == 'ﾞ' {
				if zenkakuDakuten, ok := t.KANA_H2Z_DAKUTEN_MAP[char]; ok {
					result.WriteRune(zenkakuDakuten)
					i += 2   // Skip both the current character and the dakuten
					continue // Continue to the next iteration
				}
			} else if nextChar == 'ﾟ' {

				if zenkakuHandakuten, ok := t.KANA_H2Z_MARU_MAP[char]; ok {
					result.WriteRune(zenkakuHandakuten)
					i += 2   // Skip both the current character and the handakuten
					continue // Continue to the next iteration
				}
			}
		}

		// If it's not a combined dakuten/handakuten case, perform standard conversions
		if c, ok := t.ASCII_H2Z_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else if c, ok := t.KANA_H2Z_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else if c, ok := t.DIGIT_H2Z_CHARS_MAP[char]; ok {
			result.WriteRune(c)
		} else {
			// Character is not convertible or a standalone dakuten/handakuten, etc.
			result.WriteRune(char)
		}

		i++ // Move to the next character in the rune slice
	}
	return result.String() // Return the built string
}

// H2zAt returns string that converted from half width to full width.
// Conversion string can be selected with the second argument.
// func H2zAt(str string, at ...int) string {
// 	if str == "" {
// 		return ""
// 	}

// 	if len(at) == 0 {
// 		return H2z(str)
// 	}

// 	var result strings.Builder
// 	result.Grow(len(str) * 3) //
// }
