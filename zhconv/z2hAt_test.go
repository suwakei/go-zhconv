package zhconv

import (
	"testing"
)

func TestZ2hAt(t *testing.T) {
	// Ensure convTables is initialized (it's done in init() in tables/convTable.go)
	// If tests are run package by package, init should be called.

	testCases := []struct {
		name     string
		input    string
		at       []int
		expected string
	}{
		{
			name:     "Z2hAt: Empty string",
			input:    "",
			at:       []int{0},
			expected: "",
		},
		{
			name:     "Z2hAt: No indices specified",
			input:    "ａｂｃ　１２３　テスト　ガ",
			at:       []int{},
			expected: "ａｂｃ　１２３　テスト　ガ", // Returns original string as per current implementation
		},
		{
			name:     "Z2hAt: Convert Full-width ASCII at specified index",
			input:    "ａｂｃ",
			at:       []int{1}, // 'ｂ'
			expected: "ａbｃ",
		},
		{
			name:     "Z2hAt: Convert Full-width Digit at specified index",
			input:    "１２３",
			at:       []int{0}, // '１'
			expected: "1２３",
		},
		{
			name:     "Z2hAt: Convert Full-width Kana (seion) at specified index",
			input:    "アイウ",
			at:       []int{1}, // 'イ'
			expected: "アｲウ",
		},
		{
			name:     "Z2hAt: Convert Full-width Kana with Dakuten at specified index",
			input:    "ガギグ", // 'ガ' at 0, 'ギ' at 1, 'グ' at 2
			at:       []int{0}, // 'ガ'
			expected: "ｶﾞギグ",
		},
		{
			name:     "Z2hAt: Convert Full-width Kana with Handakuten at specified index",
			input:    "パピプ",
			at:       []int{0}, // 'パ'
			expected: "ﾊﾟピプ",
		},
		{
			name:     "Z2hAt: Multiple indices, mixed types",
			input:    "ａ１カｂ２ギｃ３パ", // Indices: 0,1,2,3,4,5,6,7,8
			at:       []int{0, 2, 5}, // 'ａ', 'カ', 'ギ'
			expected: "a１ｶｂ２ｷﾞｃ３パ",
		},
		{
			name:     "Z2hAt: Index at the end of string",
			input:    "ａｂｃカ",
			at:       []int{3}, // 'カ'
			expected: "ａｂｃｶ",
		},
		{
			name:     "Z2hAt: Non-convertible char at specified index",
			input:    "あいう",
			at:       []int{1}, // 'い' (Hiragana, not in Z2H maps)
			expected: "あいう",
		},
		{
			name:     "Z2hAt: Index out of bounds (negative)",
			input:    "ａｂｃ",
			at:       []int{-1},
			expected: "ａｂｃ", // Returns original string
		},
		{
			name:     "Z2hAt: Index out of bounds (too large)",
			input:    "ａｂｃ",
			at:       []int{3},
			expected: "ａｂｃ", // Returns original string
		},
		{
			name:     "Z2hAt: Complex case with multiple conversions including decomposition",
			input:    "ＡＢＣ　１２３　テスト　ガギグ　パピプ",
			at:       []int{0, 4, 10, 13, 17}, // A, 1, ト, ギ, ピ
			expected: "AＢＣ　1２３　テスﾄ　ガｷﾞグ　パﾋﾟプ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Z2hAt(tc.input, tc.at...)
			if actual != tc.expected {
				t.Errorf("Z2hAt(%q, %v) = %q, want %q", tc.input, tc.at, actual, tc.expected)
			}
		})
	}
}