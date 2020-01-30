package entity

import (
	"regexp"
	"github.com/jinzhu/gorm"
)

// Domain はドメイン情報1つ分に対応するデータ構造
type Domain struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);unique_index;not null"`
}

// GetDomainByURL はURL文字列からドメイン名を取り出す(httpsのみ)
func GetDomainByURL(url string) Domain {
	re := regexp.MustCompile(`https://(\S+?)/`)
	result := re.FindAllStringSubmatch(url, -1)
	domainName := result[0][1]

	return Domain{Name: domainName}
}
