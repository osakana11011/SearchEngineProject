package entity

import (
    "github.com/jinzhu/gorm"
)

// InvertedData は転置データ1つ分に対応するデータ構造
type InvertedData struct {
    gorm.Model
    DocumentID uint     `gorm:"type:int;index;not null"`
    Document   Document
    TokenID    uint     `gorm:"type:int;index;not null"`
    Token      Token
    TF         float64  `gorm:"type:float"`
    Offsets    string   `gorm:"type:text"`
}
