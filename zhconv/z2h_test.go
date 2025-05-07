package zhconv

import (
	"testing"
)

func TestZ2h(t *testing.T) {
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
			name:     "No conversion needed (Hankaku)",
			input:    "abcABC123ｱｲｳｴｵｶﾞ",
			expected: "abcABC123ｱｲｳｴｵｶﾞ",
		},
		{
			name:     "No conversion needed (Hiragana, Kanji)",
			input:    "あいうえお漢字",
			expected: "あいうえお漢字",
		},
		{
			name:     "Zenkaku ASCII to Hankaku",
			input:    "ＡＢＣｄｅｆ　ＸＹＺ！＃＄％＆’（）＊＋，－．／：；＜＝＞？＠［￥］＾＿‘｛｜｝～",
			expected: "ABCdef XYZ!#$%&'()*+,-./:;<=>?@[¥]^_`{|}~",
		},
		{
			name:     "Zenkaku Digits to Hankaku",
			input:    "０１２３４５６７８９",
			expected: "0123456789",
		},
		{
			name:     "Zenkaku Katakana (Seion) to Hankaku",
			input:    "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン",
			expected: "ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝ",
		},
		{
			name:     "Zenkaku Katakana (Small) to Hankaku",
			input:    "ァィゥェォッャュョ", // 	ヮ is not converted because there is no corresponding character for half-width
			expected: "ｧｨｩｪｫｯｬｭｮ",
		},
		{
			name:     "Zenkaku Katakana (Symbols) to Hankaku",
			input:    "。、・ー「」",
			expected: "｡､･ｰ｢｣",
		},
		{
			name:     "Zenkaku Katakana (Dakuten) to Hankaku",
			input:    "ガギグゲゴザジズゼゾダヂヅデドバビブベボヴ",
			expected: "ｶﾞｷﾞｸﾞｹﾞｺﾞｻﾞｼﾞｽﾞｾﾞｿﾞﾀﾞﾁﾞﾂﾞﾃﾞﾄﾞﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞｳﾞ",
		},
		{
			name:     "Zenkaku Katakana (Handakuten) to Hankaku",
			input:    "パピプペポ",
			expected: "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ",
		},
		{
			name:     "Mixed Zenkaku/Hankaku/Other",
			input:    "これは ﾃｽﾄ です。１２３ ＡＢＣ アイウ ガギグ パピプ！",
			expected: "これは ﾃｽﾄ です｡123 ABC ｱｲｳ ｶﾞｷﾞｸﾞ ﾊﾟﾋﾟﾌﾟ!",
		},
		{
			name:     "Zenkaku Space",
			input:    "　スペース　", //  Full Width
			expected: " ｽﾍﾟｰｽ ",   // Half Width
		},
		{
			name:     "ASCII Edge Cases",
			input:    "ａｚＡＺ０９！＠～",
			expected: "azAZ09!@~",
		},
		{
			name:     "Kana Edge Cases",
			input:    "アヲンァョ。・「」ー",
			expected: "ｱｦﾝｧｮ｡･｢｣ｰ",
		},
		{
			name:     "Not convertible symbols",
			input:    "①②③㈱㈲", // It is assumed that environment dependent characters will not be converted.
			expected: "①②③㈱㈲",
		},
	}


	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Z2h(tc.input)
			if actual != tc.expected {
				t.Errorf("Z2h(%q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}
