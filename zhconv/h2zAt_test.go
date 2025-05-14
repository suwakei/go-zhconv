package zhconv

import (
	"testing"
)


func TestH2zAt(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		at       []int
		expected string
	}{
		{
			name:     "H2zAt: Empty string",
			input:    "",
			at:       []int{0},
			expected: "",
		},
		{
			name:     "H2zAt: No indices specified (should be same as H2z)",
			input:    "abc 123 ﾃｽﾄ ｶﾞ",
			at:       []int{},
			expected: "ａｂｃ　１２３　テスト　ガ",
		},
		{
			name:     "H2zAt: Convert ASCII at specified index",
			input:    "abc",
			at:       []int{1}, // 'b'
			expected: "aｂc",
		},
		{
			name:     "H2zAt: Convert Digit at specified index",
			input:    "123",
			at:       []int{0}, // '1'
			expected: "１23",
		},
		{
			name:     "H2zAt: Convert Kana at specified index",
			input:    "ｱｲｳ",
			at:       []int{1}, // 'ｲ'
			expected: "アＩウ", // 'ｲ' -> 'Ｉ' (KANA_H2Z_CHARS_MAPによる)
		},
		{
			name:     "H2zAt: Convert Kana with Dakuten at specified index",
			input:    "ｶﾞｷﾞ",
			at:       []int{0}, // 'ｶ' (and 'ﾞ')
			expected: "ガｷﾞ",
		},
		{
			name:     "H2zAt: Convert Kana with Handakuten at specified index",
			input:    "ﾊﾟﾋﾟ",
			at:       []int{0}, // 'ﾊ' (and 'ﾟ')
			expected: "パﾋﾟ",
		},
		{
			name:     "H2zAt: Multiple indices, mixed types",
			input:    "a1ｶb2ｷﾞc3ﾊﾟ", // Indices: 0,1,2,3,4,5,6,7,8
			at:       []int{0, 2, 5}, // 'a', 'ｶ', 'ｷ' (and 'ﾞ')
			expected: "ａ1カb2ギc3ﾊﾟ",
		},
		{
			name:     "H2zAt: Index at the end of string (no next char for dakuten)",
			input:    "abcｶ",
			at:       []int{3}, // 'ｶ'
			expected: "abcカ",
		},
		{
			name:     "H2zAt: Index for dakuten itself (should not convert dakuten alone)",
			input:    "ｶﾞ",
			at:       []int{1}, // 'ﾞ'
			expected: "ｶﾞ", // 'ﾞ' is not in ASCII/KANA/DIGIT maps, so remains unchanged.
		},
		{
			name:     "H2zAt: Non-convertible char at specified index",
			input:    "あいう",
			at:       []int{1}, // 'い'
			expected: "あいう",
		},
		{
			name:     "H2zAt: Index out of bounds (current impl might panic, test for graceful or expected error if modified)",
			input:    "abc",
			at:       []int{5},
			expected: "abc", // Expect no change if index is out of bounds and handled.
			// Note: Current H2zAt might panic if 'a' (index) is >= runeLen.
			// This test assumes a more robust implementation or tests the current behavior.
		},
		{
			name:     "H2zAt: Dakuten/Handakuten combination where base char is not convertible by itself",
			input:    "xﾞyﾟ", // Assume 'x' and 'y' are not in KANA_H2Z_DAKUTEN/MARU_MAP
			at:       []int{0, 2},
			expected: "xﾞyﾟ", // No conversion for 'x' or 'y', dakuten/handakuten remain.
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Note: The current H2zAt implementation has a potential panic if an index in 'at'
			// is near the end of the string, causing 'runes[a+1]' to go out of bounds.
			// These tests are written assuming the function either handles this gracefully
			// or the specific test inputs avoid such panics.
			actual := H2zAt(tc.input, tc.at...)
			if actual != tc.expected {
				t.Errorf("H2zAt(%q, %v) = %q, want %q", tc.input, tc.at, actual, tc.expected)
			}
		})
	}
}