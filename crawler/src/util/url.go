package util

import (
    "regexp"
)

// NormalizeURL はURLの正規化を行う
func NormalizeURL(url string) string {
    // URLフラグメントの削除(同じページが複数登録されるのを防ぐ為)
    sharpRegexp := regexp.MustCompile(`#.*`)
    url = sharpRegexp.ReplaceAllString(url, "")

    return url
}
