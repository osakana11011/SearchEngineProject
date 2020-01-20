package util_test

import (
    "testing"
    "strconv"

    "search_engine_project/crawler/src/util"
)

func TestRemoveEmoji(t *testing.T) {
    testCase := []struct {
        given    string
        expected string
    } {
        {"test", "test"},
        {"0123456789", "0123456789"},
        {"✌️✌️", ""},
        {"あいうえお", "あいうえお"},
        {"アイウエオ", "アイウエオ"},
        {"今日はいい天気☀️", "今日はいい天気"},
        {"✋hoge⛅️ほげ😱", "hogeほげ"},
        {"!#$%&'()", "!$%&'()"},
    }

    for i, data := range testCase {
        t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
            res := util.RemoveEmoji(data.given)
            if res != data.expected {
                t.Fatalf("\n[RemoveEmoji失敗]\nexpected: %s\nactual : %s", data.expected, res)
            }
        })
    }
}
