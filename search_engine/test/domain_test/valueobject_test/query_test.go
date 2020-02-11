package valueobject_test

import (
    "testing"
    "strconv"

    "search_engine_project/search_engine/src/domain/valueobject"
)

func TestNewQuerySuccess(t *testing.T) {
    testCase := []struct {
        q string
        expectedQueryString string
        expectedTokens []string
    } {
        {"今日はいい天気です。", "今日はいい天気です。", []string{"今日", "天気"}},
        {"隣の客はよく柿食う客だ。", "隣の客はよく柿食う客だ。", []string{"隣", "客", "柿", "客"}},
        {"ヒカキンの兄はセイキン", "ヒカキンの兄はセイキン", []string{"ヒカキン", "兄", "セイキン"}},
        {"今日はとてもいい天気😀😱", "今日はとてもいい天気", []string{"今日", "天気"}},
        {" テスト ", "テスト", []string{"テスト"}},
    }

    for i, data := range testCase {
        t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
            // Queryの生成
            res, err := valueobject.NewQuery(data.q)
            if err != nil {
                t.Fatal("\n[Query生成失敗]")
            }

            // QueryStringが正しいか検証。
            if res.QueryString() != data.expectedQueryString {
                t.Fatalf("[QueryStringの不一致] expected: %s, actual: %s", data.q, res.QueryString())
            }

            // 生成されたQueryが正しいか確かめる。
            if len(data.expectedTokens) != len(res.Tokens()) {
                t.Fatalf("[Token数の不一致] expected: %v, actual: %v", data.expectedTokens, res.Tokens())
            }

            // 生成されたTokenが全て正しいか確かめる。
            for i := 0; i < len(data.expectedTokens); i++ {
                if data.expectedTokens[i] != res.Tokens()[i] {
                    t.Fatalf("[Token不一致] expected: %v, actual: %v", data.expectedTokens, res.Tokens())
                }
            }
        })
    }
}


