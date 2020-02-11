package entity

import (
    "github.com/jinzhu/gorm"
)

// Domain はドメイン情報1つ分に対応するデータ構造
type Domain struct {
    gorm.Model
    Name string `gorm:"type:varchar(255);unique_index;not null"`
}
