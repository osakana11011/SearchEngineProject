package util_test

import (
    "testing"
    "strconv"

    "search_engine_project/search_engine/src/util"
)

func TestNormalize(t *testing.T) {
    testCase := []struct {
        given    string
        expected string
    } {
        {"０１２３４５６７８９", "0123456789"},
        {"ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ", "abcdefghijklmnopqrstuvwxyz"},
        {"！＂＃＄％＆＇（）＊＋，．／：；＜＝＞？＠［＼］＾＿｀｛｜｝", "!\"#$%&'()*+,./:;<=>?@[\\]^_`{|}"},
        {"あいうえお", "あいうえお"},
        {"アイウエオ", "アイウエオ"},
        {"日本", "日本"},
        {"       ", ""},
        {"    abc    ", "abc"},
    }

    for i, data := range testCase {
        t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
            res := util.Normalize(data.given)
            if res != data.expected {
                t.Fatalf("\n[Normalize失敗]\nexpected: %s\nactual  : %s", data.expected, res)
            }
        })
    }
}
