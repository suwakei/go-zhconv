package zhconv

import (
	"slices"
)

// H2zAt returns string that converted from half width to full width.
// Conversion string can be selected with the second argument.
func H2zAt(str string, at ...int) string {
	if str == "" {
		return ""
	}

	atLen := len(at)
	if atLen == 0 {
		return str
	}

	runes := []rune(str)
	runeLen := len(runes)

	if runeLen < slices.Max(at) {
		return str
	}

	dakutenFlag := false

	convMap := make(map[int]rune, atLen)

	for _, a := range at {
		target := runes[a]
		if a+1 < runeLen {
			next := runes[a+1]

			if next == '゛' {
				if zenkakuDakuten, ok := convTables.KANA_H2Z_DAKUTEN_MAP[target]; ok {
					convMap[a] = zenkakuDakuten
					convMap[a+1] = '゛'
					dakutenFlag = true
					continue
				}
			} else if next == 'ﾟ' {
				if zenkakuHandakuten, ok := convTables.KANA_H2Z_MARU_MAP[target]; ok {
					convMap[a] = zenkakuHandakuten
					convMap[a+1] = 'ﾟ'
					dakutenFlag = true
					continue
				}
			}
		}
		// If it's not a combined dakuten/handakuten case, perform standard conversions
		if c, ok := convTables.ASCII_H2Z_CHARS_MAP[target]; ok {
			convMap[a] = c
		} else if c, ok := convTables.KANA_H2Z_CHARS_MAP[target]; ok {
			convMap[a] = c
		} else if c, ok := convTables.DIGIT_H2Z_CHARS_MAP[target]; ok {
			convMap[a] = c
		} else {
			// Character is not convertible or a standalone dakuten/handakuten, etc.
			convMap[a] = target
		}
	}
	if !dakutenFlag {
		for atkey, atval := range convMap {
			runes[atkey] = atval
	}
	return string(runes)
	}

	resultRunes := make([]rune, 0, runeLen+len(convMap))
	for i := 0; i < runeLen; i++ {
		if val, ok := convMap[i]; ok {
			resultRunes = append(resultRunes, val)
		} else {
			resultRunes = append(resultRunes, runes[i])
		}
		}
	return string(resultRunes)
}