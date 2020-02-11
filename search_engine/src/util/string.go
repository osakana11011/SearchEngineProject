package util

import (
    "strings"
    "unicode"
)

// CutStringData はstrをmaxLenの長さで丸める
func CutStringData(str string, maxLen int, cutString string) string {
    if len(strings.Split(str, "")) < maxLen {
        maxLen = len(strings.Split(str, ""))
        cutString = ""
    }
    str = string([]rune(str)[:maxLen]) + cutString
    return str
}

// Normalize は検索文字列の正規化を行う
func Normalize(str string) string {
    // 検索文字列の前後空白を削除
    str = strings.Trim(str, " ")

    // 改行コード、タブコードの削除
    str = strings.Replace(str, "\n", "", -1)
    str = strings.Replace(str, "\t", "", -1)

    // 2つ以上の空白は1つ分の空白に置き換える
    str = strings.Replace(str, "  ", " ", -1)

    // シングルクオーテーションを削除
    str = strings.Replace(str, "'", "", -1)

    // 全角文字 => 半角文字
    str = zenkakuToHankaku(str)

    // 大文字 => 小文字
    str = strings.ToLower(str)

    return str
}

func zenkakuToHankaku(str string) string {
     return strings.ToLowerSpecial(alphanumConv, str)
}

var alphanumConv = unicode.SpecialCase{
    unicode.CaseRange {
        Lo: 0xff01, // '！'
        Hi: 0xff0c, // '，'
        Delta: [unicode.MaxCase]rune {
            0,
            0x0021 - 0xff01, // '!' - '！'
            0,
        },
    },
    unicode.CaseRange {
        Lo: 0xff0e, // '．'
        Hi: 0xff5d, // '｝'
        Delta: [unicode.MaxCase]rune {
            0,
            0x002e - 0xff0e, // '.' - '．'
            0,
        },
    },
}
