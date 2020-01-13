package util

import (
	"strings"
)

// FormatString はテキストの整形処理を行う。
// 事前に整形を施すことで、全角/半角の違いによる余分な単語乱立を防ぐことができ、検索精度の向上に繋がる。
func FormatString(text string) string {
	// 改行, タブの削除
	text = strings.NewReplacer(
		"　", " ",  // 全角空白 => 半角空白
		"\n", "",   // 改行削除
		"\t", "",   // タブ削除
	).Replace(text)

	// 全角英数字 => 半角英数字
	text = zenkakuToHankaku(text)

	// 大文字 => 小文字
	strings.ToLower(text)

	return text
}

// zenkakuToHankaku は全角の英数字を半角英数字に変換する。
func zenkakuToHankaku(text string) string {
	// 全角英字 => 半角英字
	text = strings.NewReplacer(
		"ａ", "a",
		"Ａ", "A",
		"ｂ", "b",
		"Ｂ", "B",
		"ｃ", "c",
		"Ｃ", "C",
		"ｄ", "d",
		"Ｄ", "D",
		"ｅ", "e",
		"Ｅ", "E",
		"ｆ", "f",
		"Ｆ", "F",
		"ｇ", "g",
		"Ｇ", "G",
		"ｈ", "h",
		"Ｈ", "H",
		"ｉ", "i",
		"Ｉ", "I",
		"ｊ", "j",
		"Ｊ", "J",
		"ｋ", "k",
		"Ｋ", "K",
		"ｌ", "l",
		"Ｌ", "L",
		"ｍ", "m",
		"Ｍ", "M",
		"ｎ", "n",
		"Ｎ", "N",
		"ｏ", "o",
		"Ｏ", "O",
		"ｐ", "p",
		"Ｐ", "P",
		"ｑ", "q",
		"Ｑ", "Q",
		"ｒ", "r",
		"Ｒ", "R",
		"ｓ", "s",
		"Ｓ", "S",
		"ｔ", "t",
		"Ｔ", "T",
		"ｕ", "u",
		"Ｕ", "U",
		"ｖ", "v",
		"Ｖ", "V",
		"ｗ", "w",
		"Ｗ", "W",
		"ｘ", "x",
		"Ｘ", "X",
		"ｙ", "y",
		"Ｙ", "Y",
		"ｚ", "z",
		"Ｚ", "Z",
		"０", "0",
		"１", "1",
		"２", "2",
		"３", "3",
		"４", "4",
		"５", "5",
		"６", "6",
		"７", "7",
		"８", "8",
		"９", "9",
	).Replace(text)

	return text
}
