package util

import (
	"strings"
)

// FormatString はテキストの整形処理を行う。
// 事前に整形を施すことで、全角/半角の違いによる余分な単語乱立を防ぐことができ、検索精度の向上に繋がる。
func FormatString(text string) string {
	// 改行, タブの削除
	text = strings.NewReplacer(
		"\n", "",
		"\t", "",
	).Replace(text)

	// 全角数字 => 半角数字
	text = strings.NewReplacer(
		"１", "1",
		"２", "2",
		"３", "3",
		"４", "4",
		"５", "5",
		"６", "6",
		"７", "7",
		"８", "8",
		"９", "9",
		"０", "0",
	).Replace(text)

	// 全角英字 => 半角英字
	text = strings.NewReplacer(
		"ａ", "a",
		"ｂ", "b",
		"ｃ", "c",
		"ｄ", "d",
		"ｅ", "e",
		"ｆ", "f",
		"ｇ", "g",
		"ｈ", "h",
		"ｉ", "i",
		"ｊ", "j",
		"ｋ", "k",
		"ｌ", "l",
		"ｍ", "m",
		"ｎ", "n",
		"ｏ", "o",
		"ｐ", "p",
		"ｑ", "q",
		"ｒ", "r",
		"ｓ", "s",
		"ｔ", "t",
		"ｕ", "u",
		"ｖ", "v",
		"ｗ", "w",
		"ｘ", "x",
		"ｙ", "y",
		"ｚ", "z",
	).Replace(text)

	return text
}
