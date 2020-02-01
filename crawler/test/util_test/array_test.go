package util_test

import (
    "testing"
    "strconv"
    "search_engine_project/crawler/src/util"
)

func TestUniqueArray(t *testing.T) {
    testCase := []struct {
        given []string
        expected []string
    } {
        {
            // 重複行データは削除される
            given:    []string{"1", "1", "2", "3", "4", "4", "5"},
            expected: []string{"1", "2", "3", "4", "5"},
        },
        {
            // ユニーク配列にする時、2つ目以降のデータが無視される
            given:    []string{"1", "2", "3", "1", "4", "1", "5", "1"},
            expected: []string{"1", "2", "3", "4", "5"},
        },
        {
            // 全角文字でも判定可能
            given:    []string{"あ", "い", "い", "う", "え", "お", "う"},
            expected: []string{"あ", "い", "う", "え", "お"},
        },
        {
            // 空配列でも大丈夫
            given:    []string{},
            expected: []string{},
        },
    }

    for i, data := range testCase {
        t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
            res := util.UniqArray(data.given)

            if !eqArray(res, data.expected) {
                t.Fatalf("\n[RemoveEmoji失敗]\nexpected: %s\nactual : %s", data.expected, res)
            }
        })
    }
}

func eqArray(actual []string, expected []string) bool {
    if len(actual) != len(expected) {
        return false
    }

    for i := 0; i < len(actual); i++ {
        if actual[i] != expected[i] {
            return false
        }
    }

    return true
}
