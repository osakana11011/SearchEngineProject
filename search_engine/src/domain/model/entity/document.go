package entity

import (
    "github.com/jinzhu/gorm"
)

// Document は文書情報1つに対応するデータ構造
type Document struct {
    gorm.Model
    Title          string   `gorm:"type:text"`
    URL            string   `gorm:"type:varchar(2083)"`
    Description    string   `gorm:"type:text"`
    DomainID       uint     `gorm:"type:int;index"`
    Domain         Domain
}
