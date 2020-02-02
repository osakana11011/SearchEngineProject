package entity

import (
    "strings"
    "strconv"
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

// GetInvertedList は転置リストを返す。
func GetInvertedList(documentID uint, document Document, tokens []Token) []InvertedData {
    // トークン名 => ID の変換テーブル
    token2ID := getToken2IdTable(tokens)
    tokenOffsets := generateTokenOffsets(document.UnUniqueTokens)
    tokenCounts := generateTokenCounts(document.UnUniqueTokens)
    allTokenCounts := len(document.UnUniqueTokens)

    var invertedList []InvertedData
    for tokenName, tokenCounts := range tokenCounts {
        tf := (float64)(tokenCounts) / (float64)(allTokenCounts)
        offsets := strings.Join(tokenOffsets[tokenName][:], ",")

        invertedData := InvertedData{DocumentID: documentID, TokenID: token2ID[tokenName], TF: tf, Offsets: offsets}

        invertedList = append(invertedList, invertedData)
    }

    return invertedList
}

func getToken2IdTable(tokens []Token) map[string]uint {
    token2ID := make(map[string]uint)
    for _, token := range tokens {
        token2ID[token.Name] = token.ID
    }

    return token2ID
}

func generateTokenCounts(tokens []string) map[string]int {
    tokenCounts := make(map[string]int)

    for _, tokenName := range tokens {
        _, isExist := tokenCounts[tokenName]
        if isExist {
            tokenCounts[tokenName]++
        } else {
            tokenCounts[tokenName] = 1
        }
    }

    return tokenCounts
}

func generateTokenOffsets(tokens []string) map[string][]string {
    tokenOffsets := make(map[string][]string)

    for i, tokenName := range tokens {
        _, isExist := tokenOffsets[tokenName]
        if !isExist {
            tokenOffsets[tokenName] = []string{}
        }
        tokenOffsets[tokenName] = append(tokenOffsets[tokenName], strconv.Itoa(i))
    }

    return tokenOffsets
}
