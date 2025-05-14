package tables


type ConversionTables struct {
	// ASCII_Z2H_CHARS is a Full Width ASCII characters map.
	ASCII_Z2H_CHARS_MAP map[rune]rune
	// ASCII_H2Z_CHARS is a Half Width ASCII characters map.
	ASCII_H2Z_CHARS_MAP map[rune]rune
	// KANA_Z2H_CHARS is a Full Width KANA characters map.
	KANA_Z2H_CHARS_MAP map[rune]rune
	// KANA_H2Z_CHARS is a Half Width KANA characters map.
	KANA_H2Z_CHARS_MAP map[rune]rune
	// DIGIT_Z2H_CHARS is a Full Width number characters map.
	DIGIT_Z2H_CHARS_MAP map[rune]rune
	// DIGIT_H2Z_CHARS is a Half Width number characters map.
	DIGIT_H2Z_CHARS_MAP map[rune]rune
	// KANA_TEN_MAP is a Full Width DAKUTEN_KANA characters map.
	KANA_Z2H_DAKUTEN_MAP map[rune]rune
	// KANA_MARU_MAP is a Full Width HANDAKUTEN_KANA characters map.
	KANA_Z2H_MARU_MAP map[rune]rune
	// KANA_H2Z_DAKUTEN_MAP is a Half Width DAKUTEN_KANA characters map.
	KANA_H2Z_DAKUTEN_MAP map[rune]rune
	// KANA_H2Z_MARU_MAP is a Half Width HANDAKUTEN_KANA characters map.
	KANA_H2Z_MARU_MAP map[rune]rune
}

var (
	ASCII_H2Z_CHARS_MAP = map[rune]rune{
		'a': 'ａ', 'b': 'ｂ', 'c': 'ｃ', 'd': 'ｄ', 'e': 'ｅ', 'f': 'ｆ', 'g': 'ｇ', 'h': 'ｈ', 'i': 'ｉ', 'j': 'ｊ', 'k': 'ｋ', 'l': 'ｌ', 'm': 'ｍ', 'n': 'ｎ', 'o': 'ｏ', 'p': 'ｐ', 'q': 'ｑ', 'r': 'ｒ', 's': 'ｓ', 't': 'ｔ', 'u': 'ｕ', 'v': 'ｖ', 'w': 'ｗ', 'x': 'ｘ', 'y': 'ｙ', 'z': 'ｚ',
		'A': 'Ａ', 'B': 'Ｂ', 'C': 'Ｃ', 'D': 'Ｄ', 'E': 'Ｅ', 'F': 'Ｆ', 'G': 'Ｇ', 'H': 'Ｈ', 'I': 'Ｉ', 'J': 'Ｊ', 'K': 'Ｋ', 'L': 'Ｌ', 'M': 'Ｍ', 'N': 'Ｎ', 'O': 'Ｏ', 'P': 'Ｐ', 'Q': 'Ｑ', 'R': 'Ｒ', 'S': 'Ｓ', 'T': 'Ｔ', 'U': 'Ｕ', 'V': 'Ｖ', 'W': 'Ｗ', 'X': 'Ｘ', 'Y': 'Ｙ', 'Z': 'Ｚ',
		'!': '！', '"': '”', '#': '＃', '$': '＄', '%': '％', '&': '＆', '\'': '’', '(': '（', ')': '）', '*': '＊', '+': '＋', ',': '，', '-': '－', '.': '．', '/': '／', ':': '：', ';': '；', '<': '＜', '=': '＝', '>': '＞', '?': '？', '@': '＠', '[': '［', '¥': '￥', ']': '］', '^': '＾',
		'_': '＿', '`': '‘', '{': '｛', '|': '｜', '}': '｝', '~': '～', ' ': '　', '\\': '＼',
	}

	ASCII_Z2H_CHARS_MAP = map[rune]rune{
		'ａ': 'a', 'ｂ': 'b', 'ｃ': 'c', 'ｄ': 'd', 'ｅ': 'e', 'ｆ': 'f', 'ｇ': 'g', 'ｈ': 'h', 'ｉ': 'i', 'ｊ': 'j', 'ｋ': 'k', 'ｌ': 'l', 'ｍ': 'm', 'ｎ': 'n', 'ｏ': 'o', 'ｐ': 'p', 'ｑ': 'q', 'ｒ': 'r', 'ｓ': 's', 'ｔ': 't', 'ｕ': 'u', 'ｖ': 'v', 'ｗ': 'w', 'ｘ': 'x', 'ｙ': 'y', 'ｚ': 'z',
		'Ａ': 'A', 'Ｂ': 'B', 'Ｃ': 'C', 'Ｄ': 'D', 'Ｅ': 'E', 'Ｆ': 'F', 'Ｇ': 'G', 'Ｈ': 'H', 'Ｉ': 'I', 'Ｊ': 'J', 'Ｋ': 'K', 'Ｌ': 'L', 'Ｍ': 'M', 'Ｎ': 'N', 'Ｏ': 'O', 'Ｐ': 'P', 'Ｑ': 'Q', 'Ｒ': 'R', 'Ｓ': 'S', 'Ｔ': 'T', 'Ｕ': 'U', 'Ｖ': 'V', 'Ｗ': 'W', 'Ｘ': 'X', 'Ｙ': 'Y', 'Ｚ': 'Z',
		'！': '!', '”': '"', '＃': '#', '＄': '$', '％': '%', '＆': '&', '’': '\'', '（': '(', '）': ')', '＊': '*', '＋': '+', '，': ',', '－': '-', '．': '.', '／': '/', '：': ':', '；': ';', '＜': '<', '＝': '=', '＞': '>', '？': '?', '＠': '@', '［': '[', '￥': '¥', '］': ']', '＾': '^',
		'＿': '_', '‘': '`', '｛': '{', '｜': '|', '｝': '}', '～': '~', '　': ' ', '＼': '\\',
	}

	KANA_H2Z_CHARS_MAP = map[rune]rune{
		'ｱ': 'ア', 'ｲ': 'イ', 'ｳ': 'ウ', 'ｴ': 'エ', 'ｵ': 'オ',
		'ｶ': 'カ', 'ｷ': 'キ', 'ｸ': 'ク', 'ｹ': 'ケ', 'ｺ': 'コ',
		'ｻ': 'サ', 'ｼ': 'シ', 'ｽ': 'ス', 'ｾ': 'セ', 'ｿ': 'ソ',
		'ﾀ': 'タ', 'ﾁ': 'チ', 'ﾂ': 'ツ', 'ﾃ': 'テ', 'ﾄ': 'ト',
		'ﾅ': 'ナ', 'ﾆ': 'ニ', 'ﾇ': 'ヌ', 'ﾈ': 'ネ', 'ﾉ': 'ノ',
		'ﾊ': 'ハ', 'ﾋ': 'ヒ', 'ﾌ': 'フ', 'ﾍ': 'ヘ', 'ﾎ': 'ホ',
		'ﾏ': 'マ', 'ﾐ': 'ミ', 'ﾑ': 'ム', 'ﾒ': 'メ', 'ﾓ': 'モ',
		'ﾔ': 'ヤ', 'ﾕ': 'ユ', 'ﾖ': 'ヨ',
		'ﾗ': 'ラ', 'ﾘ': 'リ', 'ﾙ': 'ル', 'ﾚ': 'レ', 'ﾛ': 'ロ',
		'ﾜ': 'ワ', 'ｦ': 'ヲ', 'ﾝ': 'ン',
		'ｧ': 'ァ', 'ｨ': 'ィ', 'ｩ': 'ゥ', 'ｪ': 'ェ', 'ｫ': 'ォ',
		'ｯ': 'ッ',
		'ｬ': 'ャ', 'ｭ': 'ュ', 'ｮ': 'ョ',
		'｡': '。', '､': '、', '･': '・', 'ﾞ': '゛', 'ﾟ': '゜',
		'｢': '「', '｣': '」', 'ｰ': 'ー',
	}

	KANA_Z2H_CHARS_MAP = map[rune]rune{
		'ア': 'ｱ', 'イ': 'ｲ', 'ウ': 'ｳ', 'エ': 'ｴ', 'オ': 'ｵ',
		'カ': 'ｶ', 'キ': 'ｷ', 'ク': 'ｸ', 'ケ': 'ｹ', 'コ': 'ｺ',
		'サ': 'ｻ', 'シ': 'ｼ', 'ス': 'ｽ', 'セ': 'ｾ', 'ソ': 'ｿ',
		'タ': 'ﾀ', 'チ': 'ﾁ', 'ツ': 'ﾂ', 'テ': 'ﾃ', 'ト': 'ﾄ',
		'ナ': 'ﾅ', 'ニ': 'ﾆ', 'ヌ': 'ﾇ', 'ネ': 'ﾈ', 'ノ': 'ﾉ',
		'ハ': 'ﾊ', 'ヒ': 'ﾋ', 'フ': 'ﾌ', 'ヘ': 'ﾍ', 'ホ': 'ﾎ',
		'マ': 'ﾏ', 'ミ': 'ﾐ', 'ム': 'ﾑ', 'メ': 'ﾒ', 'モ': 'ﾓ',
		'ヤ': 'ﾔ', 'ユ': 'ﾕ', 'ヨ': 'ﾖ',
		'ラ': 'ﾗ', 'リ': 'ﾘ', 'ル': 'ﾙ', 'レ': 'ﾚ', 'ロ': 'ﾛ',
		'ワ': 'ﾜ', 'ヲ': 'ｦ', 'ン': 'ﾝ',
		'ァ': 'ｧ', 'ィ': 'ｨ', 'ゥ': 'ｩ', 'ェ': 'ｪ', 'ォ': 'ｫ',
		'ッ': 'ｯ',
		'ャ': 'ｬ', 'ュ': 'ｭ', 'ョ': 'ｮ',
		'。': '｡', '、': '､', '・': '･', '゛': 'ﾞ', '゜': 'ﾟ',
		'「': '｢', '」': '｣', 'ー': 'ｰ',
	}

	DIGIT_H2Z_CHARS_MAP = map[rune]rune{
		'0': '０', '1': '１', '2': '２', '3': '３', '4': '４',
		'5': '５', '6': '６', '7': '７', '8': '８', '9': '９',
	}

	DIGIT_Z2H_CHARS_MAP = map[rune]rune{
		'０': '0', '１': '1', '２': '2', '３': '3', '４': '4',
		'５': '5', '６': '6', '７': '7', '８': '8', '９': '9',
	}

	KANA_H2Z_DAKUTEN_MAP = map[rune]rune{
		'ｶ': 'カ', 'ｷ': 'キ', 'ｸ': 'ク', 'ｹ': 'ケ', 'ｺ': 'コ',
		'ｻ': 'サ', 'ｼ': 'シ', 'ｽ': 'ス', 'ｾ': 'セ', 'ｿ': 'ソ',
		'ﾀ': 'タ', 'ﾁ': 'チ', 'ﾂ': 'ツ', 'ﾃ': 'テ', 'ﾄ': 'ト',
		'ﾊ': 'ハ', 'ﾋ': 'ヒ', 'ﾌ': 'フ', 'ﾍ': 'ヘ', 'ﾎ': 'ホ',
		'ｳ': 'ウ', 'ﾞ': 'ﾞ', 'ﾟ': 'ﾟ',
	}

	KANA_H2Z_MARU_MAP = map[rune]rune{
		'ﾊ': 'ハ', 'ﾋ': 'ﾋ', 'ﾌ': 'フ', 'ﾍ': 'ヘ', 'ﾎ': 'ホ',
	}

	KANA_Z2H_DAKUTEN_MAP = map[rune]rune{
		'ガ': 'カ', 'ギ': 'キ', 'グ': 'ク', 'ゲ': 'ケ', 'ゴ': 'コ',
		'ザ': 'サ', 'ジ': 'シ', 'ズ': 'ス', 'ゼ': 'セ', 'ゾ': 'ソ',
		'ダ': 'タ', 'ヂ': 'チ', 'ヅ': 'ツ', 'デ': 'テ', 'ド': 'ト',
		'バ': 'ハ', 'ビ': 'ヒ', 'ブ': 'フ', 'ベ': 'ヘ', 'ボ': 'ホ',
		'ヴ': 'ウ',
	}

	KANA_Z2H_MARU_MAP = map[rune]rune{
		'パ': 'ハ', 'ピ': 'ヒ', 'プ': 'フ', 'ペ': 'ヘ', 'ポ': 'ホ',
	}
)

func New() *ConversionTables {
	return &ConversionTables{
		ASCII_Z2H_CHARS_MAP:  ASCII_Z2H_CHARS_MAP,
		ASCII_H2Z_CHARS_MAP:  ASCII_H2Z_CHARS_MAP,
		KANA_Z2H_CHARS_MAP:   KANA_Z2H_CHARS_MAP,
		KANA_H2Z_CHARS_MAP:   KANA_H2Z_CHARS_MAP,
		DIGIT_Z2H_CHARS_MAP:  DIGIT_Z2H_CHARS_MAP,
		DIGIT_H2Z_CHARS_MAP:  DIGIT_H2Z_CHARS_MAP,
		KANA_Z2H_DAKUTEN_MAP: KANA_Z2H_DAKUTEN_MAP,
		KANA_Z2H_MARU_MAP:    KANA_Z2H_MARU_MAP,
		KANA_H2Z_DAKUTEN_MAP: KANA_H2Z_DAKUTEN_MAP,
		KANA_H2Z_MARU_MAP:    KANA_H2Z_MARU_MAP,
	}
}
