package entity

import (
    "fmt"
    "regexp"
    "github.com/jinzhu/gorm"
)

const (
    acceptDomain = "ja.wikipedia.org"
    denyExtensions = ".svg|.jpg|.png|.gif"
)

// CrawlWaiting はクロール待ち情報1つ分に対応するデータ構造
type CrawlWaiting struct {
    gorm.Model
    URL        string     `gorm:"type:varchar(2083)"`
    IsPriority bool       `gorm:"type:boolean;default:0;index"`
}

// IsValid はセットされているクロールするかどうが判定する
func (d *CrawlWaiting) IsValid() bool {
    // ドメインと拡張子を見て有効性を判断する
    return isAcceptDomain(d.URL) && isAcceptExtension(d.URL)
}

func isAcceptDomain(url string) bool {
    domainRegexp := regexp.MustCompile(fmt.Sprintf(`^https://%s/`, acceptDomain))
    if domainRegexp.MatchString(url) {
        return true
    }
    return false
}

func isAcceptExtension(url string) bool {
    extensionsRegexp := regexp.MustCompile(fmt.Sprintf(`[%s]$`, denyExtensions))
    if !extensionsRegexp.MatchString(url) {
        return true
    }
    return false
}
