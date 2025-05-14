package zhconv

import (
	"slices"
)

// Z2zAt returns string that converted from full-width to half-width.
// Conversion string can be selected with the second argument.
func Z2hAt(str string, at ...int) string {
	if str == "" {
		return ""
	}

	atLen := len(at)
	if atLen == 0 {
		return str
	}

	runes := []rune(str)
	runeLen := len(runes)

	// Validate indices
	if atLen > 0 && (slices.Min(at) < 0 || slices.Max(at) >= runeLen) {
		return str
	}

// Create a set of `at` indices for quick lookup
	atSet := make(map[int]struct{})
	for _, index := range at {
		atSet[index] = struct{}{}
	}

	var outputRunes []rune = make([]rune, 0, runeLen) // Estimate capacity
	var convertedRunes []rune

	for i := 0; i < runeLen; i++ {
		charToProcess := runes[i]
		if _, shouldConvert := atSet[i]; shouldConvert {
			convertedRunes = nil
			// Check for Katakana with dakuten/handakuten first for decomposition
			// KANA_Z2H_DAKUTEN_MAP maps 'ガ' -> 'ｶ' (hankaku base)
			if hankakuBase, ok := convTables.KANA_Z2H_DAKUTEN_MAP[charToProcess]; ok {
				convertedRunes = append(convertedRunes, []rune{hankakuBase, 'ﾞ'}...) // e.g. 'ガ' -> ['ｶ', 'ﾞ']
			} else if hankakuBase, ok := convTables.KANA_Z2H_MARU_MAP[charToProcess]; ok {
				convertedRunes = append(convertedRunes, []rune{hankakuBase, 'ﾟ'}...) // e.g. 'パ' -> ['ﾊ', 'ﾟ']
			}
			
			if c, ok := convTables.ASCII_Z2H_CHARS_MAP[charToProcess]; ok {
				convertedRunes = append(convertedRunes, c) // e.g. 'Ａ' -> ['A']
			} else if c, ok := convTables.KANA_Z2H_CHARS_MAP[charToProcess]; ok {
				convertedRunes = append(convertedRunes, c) // e.g. 'ア' -> ['ｱ']
			} else if c, ok := convTables.DIGIT_Z2H_CHARS_MAP[charToProcess]; ok {
				convertedRunes = append(convertedRunes, c) // e.g. '１' -> ['1']
			}
			if convertedRunes != nil {
				outputRunes = append(outputRunes, convertedRunes...)
			} else {
				outputRunes = append(outputRunes, charToProcess) // add original character without change
			}
		}else {
			outputRunes = append(outputRunes, charToProcess) // add original character without change
		}
	}
	return string(outputRunes)
}