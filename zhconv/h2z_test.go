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
			expected: "ＡＢＣｄｅｆ　ＸＹＺ！＃＄％＆’（）＊＋，－．／：；＜＝＞？＠［¥］＾＿‘｛｜｝～　￥",
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
			expected: "カﾞキﾞクﾞケﾞコﾞサﾞシﾞスﾞセﾞソﾞタﾞチﾞツﾞテﾞトﾞハﾞヒﾞフﾞヘﾞホﾞウﾞ",
		},
		{
			name:     "Hankaku Katakana with Handakuten to Zenkaku",
			input:    "ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ",
			expected: "ハﾟヒﾟフﾟヘﾟホﾟ",
		},
		{
			name:     "Mixed Hankaku/Zenkaku/Other",
			input:    "これはﾃｽﾄです｡123 ABC ｱｲｳ ｶﾞｷﾞｸﾞ ﾊﾟﾋﾟﾌﾟ!",
			expected: "これはテストです。１２３　ＡＢＣ　アイウ　カﾞキﾞクﾞ　ハﾟヒﾟフﾟ！",
		},
		{
			name:     "Hankaku Space",
			input:    " ｽﾍﾟｰｽ ",
			expected: "　スヘﾟース　",
		},
		{
			name:     "Dakuten/Handakuten cannot be applied",
			input:    "ｱﾞｲﾟﾝﾞ", // アに濁点、イに半濁点、ンに濁点は付けられない
			expected: "ア゛イ゜ン゛", // 分離されたまま全角に変換される (ｱ->ア, ﾞ->ﾞ)
		},
		{
			name:     "Edge case: Dakuten at end",
			input:    "ﾃｽﾄｶﾞ",
			expected: "テストカﾞ",
		},
		{
			name:     "Edge case: Handakuten at end",
			input:    "ﾃｽﾄﾊﾟ",
			expected: "テストハﾟ",
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
			expected: "１ハﾞイト文字と２ハﾞイト文字が混在するテキスト。　ＡＢＣ　１２３　カﾞキﾞクﾞハﾟヒﾟフﾟ！？",
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