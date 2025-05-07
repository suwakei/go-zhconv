package tables

type ConversionTables struct {
	// ASCII_ZENKAKU_CHARS は全角ASCII文字の配列です。
	ASCII_ZENKAKU_CHARS []rune
	// ASCII_HANKAKU_CHARS は半角ASCII文字の配列です。
	ASCII_HANKAKU_CHARS []rune
	// KANA_ZENKAKU_CHARS は全角カナ文字の配列です
	KANA_ZENKAKU_CHARS []rune
	// KANA_HANKAKU_CHARS は半角カナ文字の配列です。
	KANA_HANKAKU_CHARS []rune
	// DIGIT_ZENKAKU_CHARS は全角数字の配列です
	DIGIT_ZENKAKU_CHARS []rune
	// DIGIT_HANKAKU_CHARS は半角数字の配列です。
	DIGIT_HANKAKU_CHARS []rune
	// KANA_TEN_MAP は全角カナの濁点マッピングです。
	KANA_ZENKAKU_DAKUTEN_MAP map[rune]rune
	// KANA_MARU_MAP は全角カナの丸点マッピングです。
	KANA_ZENKAKU_MARU_MAP map[rune]rune
	// KANA_HANKAKU_DAKUTEN_MAP は半角カナの濁点マッピングです。
	KANA_HANKAKU_DAKUTEN_MAP map[rune]rune
	// KANA_HANKAKU_MARU_MAP は半角カナの丸点マッピングです。
	KANA_HANKAKU_MARU_MAP map[rune]rune
}

var (
	ASCII_ZENKAKU_CHARS = []rune{
		'ａ', 'ｂ', 'ｃ', 'ｄ', 'ｅ', 'ｆ', 'ｇ', 'ｈ', 'ｉ', 'ｊ', 'ｋ', 'ｌ', 'ｍ', 'ｎ', 'ｏ', 'ｐ', 'ｑ', 'ｒ', 'ｓ', 'ｔ', 'ｕ', 'ｖ', 'ｗ', 'ｘ', 'ｙ', 'ｚ',
		'Ａ', 'Ｂ', 'Ｃ', 'Ｄ', 'Ｅ', 'Ｆ', 'Ｇ', 'Ｈ', 'Ｉ', 'Ｊ', 'Ｋ', 'Ｌ', 'Ｍ', 'Ｎ', 'Ｏ', 'Ｐ', 'Ｑ', 'Ｒ', 'Ｓ', 'Ｔ', 'Ｕ', 'Ｖ', 'Ｗ', 'Ｘ', 'Ｙ', 'Ｚ',
		'！', '”', '＃', '＄', '％', '＆', '’', '（', '）', '＊', '＋', '，', '－', '．', '／', '：', '；', '＜', '＝', '＞', '？', '＠', '［', '￥', '］', '＾', '＿', '‘', '｛', '｜', '｝', '～', '　', '＼',
	}

	ASCII_HANKAKU_CHARS = []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '¥', ']', '^', '_', '`', '{', '|', '}', '~', ' ', '\\',
	}

	KANA_ZENKAKU_CHARS = []rune{
		'ア', 'イ', 'ウ', 'エ', 'オ', 'カ', 'キ', 'ク', 'ケ', 'コ', 'サ', 'シ', 'ス', 'セ', 'ソ', 'タ', 'チ', 'ツ', 'テ', 'ト', 'ナ', 'ニ', 'ヌ', 'ネ', 'ノ', 'ハ', 'ヒ', 'フ', 'ヘ', 'ホ', 'マ', 'ミ', 'ム', 'メ', 'モ', 'ヤ', 'ユ', 'ヨ', 'ラ', 'リ', 'ル', 'レ', 'ロ', 'ワ', 'ヲ', 'ン', 'ァ', 'ィ', 'ゥ', 'ェ', 'ォ', 'ッ', 'ャ', 'ュ', 'ョ', '。', '、', '・', '゛', '゜', '「', '」', 'ー',
	}

	KANA_HANKAKU_CHARS = []rune{
		'ｱ', 'ｲ', 'ｳ', 'ｴ', 'ｵ', 'ｶ', 'ｷ', 'ｸ', 'ｹ', 'ｺ', 'ｻ', 'ｼ', 'ｽ', 'ｾ', 'ｿ', 'ﾀ', 'ﾁ', 'ﾂ', 'ﾃ', 'ﾄ', 'ﾅ', 'ﾆ', 'ﾇ', 'ﾈ', 'ﾉ', 'ﾊ', 'ﾋ', 'ﾌ', 'ﾍ', 'ﾎ', 'ﾏ', 'ﾐ', 'ﾑ', 'ﾒ', 'ﾓ', 'ﾔ', 'ﾕ', 'ﾖ', 'ﾗ', 'ﾘ', 'ﾙ', 'ﾚ', 'ﾛ', 'ﾜ', 'ｦ', 'ﾝ', 'ｧ', 'ｨ', 'ｩ', 'ｪ', 'ｫ', 'ｯ', 'ｬ', 'ｭ', 'ｮ', '｡', '､', '･', 'ﾞ', 'ﾟ', '｢', '｣', 'ｰ',
	}

	DIGIT_ZENKAKU_CHARS = []rune{'０', '１', '２', '３', '４', '５', '６', '７', '８', '９'}

	DIGIT_HANKAKU_CHARS = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	KANA_HANKAKU_DAKUTEN_MAP = map[rune]rune{
		'ｶ': 'ガ', 'ｷ': 'ギ', 'ｸ': 'グ', 'ｹ': 'ゲ', 'ｺ': 'ゴ',
		'ｻ': 'ザ', 'ｼ': 'ジ', 'ｽ': 'ズ', 'ｾ': 'ゼ', 'ｿ': 'ゾ',
		'ﾀ': 'ダ', 'ﾁ': 'ヂ', 'ﾂ': 'ヅ', 'ﾃ': 'デ', 'ﾄ': 'ド',
		'ﾊ': 'バ', 'ﾋ': 'ビ', 'ﾌ': 'ブ', 'ﾍ': 'ベ', 'ﾎ': 'ボ',
		'ｳ': 'ヴ', 'ﾞ': 'ﾞ', 'ﾟ': 'ﾟ',
	}

	KANA_HANKAKU_MARU_MAP = map[rune]rune{
		'ﾊ': 'パ', 'ﾋ': 'ピ', 'ﾌ': 'プ', 'ﾍ': 'ペ', 'ﾎ': 'ポ',
	}

	KANA_ZENKAKU_DAKUTEN_MAP = map[rune]rune{
		'ガ': 'カ', 'ギ': 'キ', 'グ': 'ク', 'ゲ': 'ケ', 'ゴ': 'コ',
		'ザ': 'サ', 'ジ': 'シ', 'ズ': 'ス', 'ゼ': 'セ', 'ゾ': 'ソ',
		'ダ': 'タ', 'ヂ': 'チ', 'ヅ': 'ツ', 'デ': 'テ', 'ド': 'ト',
		'バ': 'ハ', 'ビ': 'ヒ', 'ブ': 'フ', 'ベ': 'ヘ', 'ボ': 'ホ',
		'ヴ': 'ウ', // ヴ -> ウ -> ｳ
	}
	KANA_ZENKAKU_MARU_MAP = map[rune]rune{
		'パ': 'ハ', 'ピ': 'ヒ', 'プ': 'フ', 'ペ': 'ヘ', 'ポ': 'ホ',
	}
)

func New() *ConversionTables {
	return &ConversionTables{
		ASCII_ZENKAKU_CHARS:      ASCII_ZENKAKU_CHARS,
		ASCII_HANKAKU_CHARS:      ASCII_HANKAKU_CHARS,
		KANA_ZENKAKU_CHARS:       KANA_ZENKAKU_CHARS,
		KANA_HANKAKU_CHARS:       KANA_HANKAKU_CHARS,
		DIGIT_ZENKAKU_CHARS:      DIGIT_ZENKAKU_CHARS,
		DIGIT_HANKAKU_CHARS:      DIGIT_HANKAKU_CHARS,
		KANA_ZENKAKU_DAKUTEN_MAP: KANA_ZENKAKU_DAKUTEN_MAP,
		KANA_ZENKAKU_MARU_MAP:    KANA_ZENKAKU_MARU_MAP,
		KANA_HANKAKU_DAKUTEN_MAP: KANA_HANKAKU_DAKUTEN_MAP,
		KANA_HANKAKU_MARU_MAP:    KANA_HANKAKU_MARU_MAP,
	}
}
