package zhconv

import (
	"strings"
)

func Reverse(str string) string {
		if str == "" {
		return ""
	}

	var sb strings.Builder
	runes := []rune(str)
	n := len(runes)
	// Pre-allocate memory; actual length might differ due to multi-rune characters
	// or single rune to multi-rune conversions (e.g., 'ガ' -> 'ｶ','ﾞ').
	sb.Grow(n + n/2) // Estimate a bit more capacity

	i := 0
	for i < n {
		char := runes[i]
		converted := false

		// --- 1. Try to convert from Full-width to Half-width ---

		// Full-width Katakana with Dakuten (e.g., 'ガ' -> 'ｶﾞ')
		if halfBase, ok := convTables.KANA_Z2H_DAKUTEN_MAP[char]; ok {
			sb.WriteRune(halfBase) // e.g., 'ｶ'
			sb.WriteRune('ﾞ')     // Half-width Dakuten mark
			i++
			converted = true
		} else if halfBase, ok := convTables.KANA_Z2H_MARU_MAP[char]; ok {
			sb.WriteRune(halfBase)
			sb.WriteRune('ﾟ')     // Half-width Handakuten mark
			i++
			converted = true
		} else if half, ok := convTables.ASCII_Z2H_CHARS_MAP[char]; ok { // Full-width ASCII (e.g., 'Ａ' -> 'A', '　' -> ' ')
			sb.WriteRune(half)
			i++
			converted = true
		} else if half, ok := convTables.DIGIT_Z2H_CHARS_MAP[char]; ok { // Full-width Digits (e.g., '１' -> '1')
			sb.WriteRune(half)
			i++
			converted = true
		} else if half, ok := convTables.KANA_Z2H_CHARS_MAP[char]; ok { // Full-width Katakana (standard, e.g., 'ア' -> 'ｱ')
			sb.WriteRune(half)
			i++
			converted = true
		}

		if converted {
			continue
		}

		// --- 2. Try to convert from Half-width to Full-width (if not converted above) ---

		// Check for Half-width Katakana with Dakuten/Handakuten (2 runes sequence)
		if i+1 < n {
			nextChar := runes[i+1]
			if nextChar == 'ﾞ' { // Half-width Dakuten mark
				if zenkakuBase, ok := convTables.KANA_H2Z_DAKUTEN_MAP[char]; ok {
					sb.WriteRune(zenkakuBase)
					sb.WriteRune('ﾞ')
					i += 2
					converted = true
				}
			} else if nextChar == 'ﾟ' {
				if zenkakuBase, ok := convTables.KANA_H2Z_MARU_MAP[char]; ok {
					sb.WriteRune(zenkakuBase)
					sb.WriteRune('ﾟ')
					i += 2
					converted = true
				}
			}
		}

		if converted {
			continue
		}

		// Standard Half-width to Full-width conversions (single rune)
		if full, ok := convTables.ASCII_H2Z_CHARS_MAP[char]; ok {
			sb.WriteRune(full)
			i++
		} else if full, ok := convTables.DIGIT_H2Z_CHARS_MAP[char]; ok {
			sb.WriteRune(full)
			i++
		} else if full, ok := convTables.KANA_H2Z_CHARS_MAP[char]; ok {
			sb.WriteRune(full)
			i++
		} else {
			// If no conversion is applicable, append the original rune
			sb.WriteRune(char)
			i++
		}
	}
	return sb.String()
}