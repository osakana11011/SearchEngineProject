package entity

import (
    "github.com/jinzhu/gorm"
)

// CrawlWaiting はクロール待ち情報1つ分に対応するデータ構造
type CrawlWaiting struct {
    gorm.Model
    URL        string     `gorm:"type:varchar(2083)"`
    IsPriority bool       `gorm:"type:boolean;default:0;index"`
}

// IsValid はセットされているクロールするかどうが判定する
func (d *CrawlWaiting) IsValid() bool {
    return true
}
