package entity

import (
	"github.com/jinzhu/gorm"
)

// Token はトークン情報1つに対応するデータ構造
type Token struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);unique_index;not null"`
}
