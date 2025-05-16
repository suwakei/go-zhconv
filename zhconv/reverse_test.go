package zhconv

import (
	"testing"
)

// test/widthとのベンチマークも忘れずに書く

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Full-width ASCII to Half-width",
			input:    "ＡＢＣ　ＸＹＺ！",
			expected: "ABC XYZ!",
		},
		{
			name:     "Half-width ASCII to Full-width",
			input:    "abc xyz!",
			expected: "ａｂｃ　ｘｙｚ！",
		},
		{
			name:     "Full-width Digits to Half-width",
			input:    "０１２３４５",
			expected: "012345",
		},
		{
			name:     "Half-width Digits to Full-width",
			input:    "67890",
			expected: "６７８９０",
		},
		{
			name:     "Full-width Katakana (Seion) to Half-width",
			input:    "アイウエオ",
			expected: "ｱｲｳｴｵ",
		},
		{
			name:     "Half-width Katakana (Seion) to Full-width",
			input:    "ｶｷｸｹｺ",
			expected: "カキクケコ",
		},
		{
			name:     "Full-width Katakana (Dakuten) to Half-width",
			input:    "ガギグゲゴ",
			expected: "ｶﾞｷﾞｸﾞｹﾞｺﾞ",
		},
		{
			name:     "Half-width Katakana (Dakuten) to Full-width",
			input:    "ｻﾞｼﾞｽﾞｾﾞｿﾞ",
			expected: "サﾞシﾞスﾞセﾞソﾞ", // H2zの挙動に合わせる
		},
		{
			name:     "Full-width Katakana (Handakuten) to Half-width",
			input:    "パピプペポ",
			expected: "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ",
		},
		{
			name:     "Half-width Katakana (Handakuten) to Full-width",
			input:    "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ",
			expected: "ハﾟヒﾟフﾟヘﾟホﾟ", // H2zの挙動に合わせる
		},
		{
			name:     "Mixed Full-width to Half-width",
			input:    "Ｈｅｌｌｏ　Ｗｏｒｌｄ！　１２３　アイウガパ",
			expected: "Hello World! 123 ｱｲｳｶﾞﾊﾟ",
		},
		{
			name:     "Mixed Half-width to Full-width",
			input:    "Hello World! 123 ｱｲｳｶﾞﾊﾟ",
			expected: "Ｈｅｌｌｏ　Ｗｏｒｌｄ！　１２３　アイウカﾞハﾟ",
		},
		{
			name:     "Mixed Full-Half already, should reverse",
			input:    "ＡbcＤＥＦｇｈ　1２３ＸＹｚ　アｲウｴオ　カｷﾞクｹﾞコ　サｼﾞスｾﾞソ　タﾁヅﾃド",
			expected: "AｂｃDEFgh １23XYz ｱイｳエｵ ｶキﾞｸケﾞｺ ｻシﾞｽセﾞｿ ﾀチﾂﾞテﾄﾞ",
		},
		{
			name:     "Non-convertible characters (Hiragana, Kanji)",
			input:    "あいうえお漢字ＡＢＣ123ｱｲｳ",
			expected: "あいうえお漢字ABC１２３アイウ",
		},
		{
			name:     "String with only non-convertible characters",
			input:    "春夏秋冬",
			expected: "春夏秋冬",
		},
		{
			name:     "Full-width symbols to Half-width",
			input:    "。、・ー「」",
			expected: "｡､･ｰ｢｣",
		},
		{
			name:     "Half-width symbols to Full-width",
			input:    "｡､･ｰ｢｣",
			expected: "。、・ー「」",
		},
		{
			name:     "Complex mix",
			input:    "１ｓｔ「ＰＲＩＣＥ」ｉｓ　￥１，０００－　（ＴＡＸ　ＩＮ）　ｶﾞﾝﾊﾞﾚ！",
			expected: "1st｢PRICE｣is \\1,000- (TAX IN) カﾞンハﾞレ!",
		},
		{
			name:     "Half-width dakuten/handakuten at end of string",
			input:    "テストｶﾞ",
			expected: "ﾃｽﾄカﾞ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Reverse(tc.input)
			if actual != tc.expected {
				t.Errorf("Reverse(%q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}
