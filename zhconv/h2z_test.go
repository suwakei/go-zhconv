package zhconv

import (
	"testing"
)

func BenchmarkH2z(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		H2z("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	}
}

func TestH2z(t *testing.T) {
	testCases := []struct {
		name     string // testcase name
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "No conversion needed (Zenkaku)",
			input:    "ＡＢＣ１２３アイウガパ",
			expected: "ＡＢＣ１２３アイウガパ",
		},
		{
			name:     "No conversion needed (Hiragana, Kanji)",
			input:    "あいうえお漢字１２３",
			expected: "あいうえお漢字１２３",
		},
		{
			name:     "Hankaku ASCII to Zenkaku",
			input:    "ABCdef XYZ!#$%&'()*+,-./:;<=>?@[¥]^_`{|}~ \\",
			expected: "ＡＢＣｄｅｆ　ＸＹＺ！＃＄％＆’（）＊＋，－．／：；＜＝＞？＠［￥］＾＿‘｛｜｝～　＼",
		},
		{
			name:     "Hankaku Digits to Zenkaku",
			input:    "0123456789",
			expected: "０１２３４５６７８９",
		},
		{
			name:     "Hankaku Katakana to Zenkaku",
			input:    "ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝ",
			expected: "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン",
		},
		{
			name:     "Hankaku Katakana (Small) to Zenkaku",
			input:    "ｧｨｩｪｫｯｬｭｮ",
			expected: "ァィゥェォッャュョ",
		},
		{
			name:     "Hankaku Katakana (Symbols) to Zenkaku",
			input:    "｡､･ｰ｢｣",
			expected: "。、・ー「」",
		},
		{
			name:     "Hankaku Katakana with Dakuten to Zenkaku",
			input:    "ｶﾞｷﾞｸﾞｹﾞｺﾞｻﾞｼﾞｽﾞｾﾞｿﾞﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞｳﾞ",
			expected: "ガギグゲゴザジズゼゾダヂヅデドバビブベボヴ",
		},
		{
			name:     "Hankaku Katakana with Handakuten to Zenkaku",
			input:    "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ",
			expected: "パピプペポ",
		},
		{
			name:     "Mixed Hankaku/Zenkaku/Other",
			input:    "これはﾃｽﾄです｡123 ABC ｱｲｳ ｶﾞｷﾞｸﾞ ﾊﾟﾋﾟﾌﾟ!",
			expected: "これはテストです。１２３　ＡＢＣ　アイウ　ガギグ　パピプ！",
		},
		{
			name:     "Hankaku Space",
			input:    " ｽﾍﾟｰｽ ", // 半角スペース
			expected: "　スペース　",  // 全角スペース
		},
		{
			name:     "Dakuten/Handakuten cannot be applied",
			input:    "ｱﾞｲﾟﾝﾞ", // アに濁点、イに半濁点、ンに濁点は付けられない
			expected: "ア゛イ゜ン゛", // 分離されたまま全角に変換される (ｱ->ア, ﾞ->ﾞ)
		},
		{
			name:     "Edge case: Dakuten at end",
			input:    "ﾃｽﾄｶﾞ",
			expected: "テストガ",
		},
		{
			name:     "Edge case: Handakuten at end",
			input:    "ﾃｽﾄﾊﾟ",
			expected: "テストパ",
		},
		{
			name:     "ASCII Edge Cases",
			input:    "azAZ09!@~ ",
			expected: "ａｚＡＺ０９！＠～　",
		},
		{
			name:     "Kana Edge Cases",
			input:    "ｱｦﾝｧｮ｡･｢｣ｰ",
			expected: "アヲンァョ。・「」ー",
		},
		{
			name:     "Not convertible symbols",
			input:    "①②③㈱㈲", // It is assumed that environment dependent characters will not be converted.
			expected: "①②③㈱㈲",
		},
		{
			name:     "Long string with various conversions",
			input:    "1ﾊﾞｲﾄ文字と2ﾊﾞｲﾄ文字が混在するﾃｷｽﾄ｡ ABC 123 ｶﾞｷﾞｸﾞﾊﾟﾋﾟﾌﾟ!?",
			expected: "１バイト文字と２バイト文字が混在するテキスト。　ＡＢＣ　１２３　ガギグパピプ！？",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := H2z(tc.input)
			if actual != tc.expected {
				t.Errorf("H2z(%q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}

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