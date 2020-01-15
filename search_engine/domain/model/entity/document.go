package entity

// Document はWeb文書1つに対応するデータ構造。
type Document struct {
	ID    int64   // 文書ID
	Title string  // タイトル
	URL   string  // URL
}
