package datastore

import (
    "fmt"
    "math"
    "strings"

    "search_engine_project/crawler/src/domain/repository"
)

// TokenRepository はトークンに関するDB操作を行う構造体
type TokenRepository struct {}

// NewTokenRepository はトークンのDB操作を提供するハンドラを返す
func NewTokenRepository() repository.TokenRepository {
    return &TokenRepository{}
}

// BulkInsert はトークン情報をバルクインサートする
func (r *TokenRepository) BulkInsert(tokens []string) error {
    // 登録する単語が0の時はそのまま返す
    if len(tokens) == 0 {
        return nil
    }

    // DB接続
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    // バルクインサートする回数
    bulkNum := (int)(math.Ceil(float64(len(tokens)) / 100.0))
    for i := 0; i < bulkNum; i++ {
        // バルクインサート用のSQLを構築
        bulkInsertSQL := "INSERT IGNORE INTO tokens (token, created_at, updated_at) VALUES "
        to := i * 100
        from := (int)(math.Min((float64)(i*100+100), (float64)(len(tokens))))
        tokensMass := tokens[to:from]
        for _, token := range tokensMass {
            token = strings.NewReplacer("'", "",).Replace(token)
            bulkInsertSQL += fmt.Sprintf("('%s', NOW(), NOW()), ", token)
        }
        bulkInsertSQL = bulkInsertSQL[:len(bulkInsertSQL)-2]

        // 登録処理
        _, err := db.Exec(bulkInsertSQL)
        if err != nil {
            return err
        }
    }

    return nil
}

// GetIDs は知りたいトークンの配列を取り、一度のクエリでIDを取得する
func (r *TokenRepository) GetIDs(tokens []string) (map[string]int64, error) {
    // tokensが無い場合は、クエリが組めないので空の連想配列を返す
    if (len(tokens) == 0) {
        return map[string]int64{}, nil
    }

    // DB接続
    db, err := connectDB()
    if err != nil {
        return map[string]int64{}, err
    }
    defer db.Close()

    // クエリを投げる
    rows, err := db.Query("SELECT id, token FROM tokens WHERE token IN('" + strings.Join(tokens, "','") + "')")
    if err != nil {
        return map[string]int64{}, err
    }
    defer rows.Close()

    // トークンからIDを引っ張れるように変換していく
    tokenLookUpTable := make(map[string]int64)
    for rows.Next() {
        var tokenID int64
        var token string
        if err := rows.Scan(&tokenID, &token); err != nil {
            return map[string]int64{}, err
        }
        tokenLookUpTable[token] = tokenID
    }

    return tokenLookUpTable, nil
}

// GetID はトークン名からトークンIDを求める
func (r *TokenRepository) GetID(token string) (int64, error) {
    // DB接続
    db, err := connectDB()
    if err != nil {
        return 0, err
    }
    defer db.Close()

    // IDの取得
    var id int64
    db.QueryRow("SELECT id FROM tokens WHERE token = ? LIMIT 1", token).Scan(&id)

    return id, nil
}

// GetCounts はトークン名を取ってそのデータに合致する数を返す
func (r *TokenRepository) GetCounts(token string) (int, error) {
    // DB接続
    db, err := connectDB()
    if err != nil {
        return 0, err
    }
    defer db.Close()

    // 件数の取得
    var counts int
    db.QueryRow("SELECT COUNT(id) FROM tokens WHERE token = ?", token).Scan(&counts)

    return counts, nil
}
