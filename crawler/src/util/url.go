package util

import (
    "regexp"
)

// NormalizeURL はURLの正規化を行う
func NormalizeURL(url string, domainName string) string {
    // URLフラグメントの削除(同じページが複数登録されるのを防ぐ為)
    sharpRegexp := regexp.MustCompile(`#.*`)
    url = sharpRegexp.ReplaceAllString(url, "")

    noProtcolRegexp := regexp.MustCompile(`^//.*`)
    noDomainRegexp := regexp.MustCompile(`^/.*`)

    if noProtcolRegexp.MatchString(url) {
        return "https:" + url
    }
    if noDomainRegexp.MatchString(url) {
        return "https://" + domainName + url
    }

    return url
}
