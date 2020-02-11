package valueobject

// Document はWeb文書1つに対応する構造体。
type Document struct {
    title string  // タイトル
    url   string  // URL
}

// NewDocument は新しくDocumentを生成して返す。
func NewDocument(title string, url string) Document {
    return Document{title: title, url: url}
}

// Title はtitleの値を取得する為のメソッド。
func (d *Document) Title() string {
    return d.title
}

// URL はurlの値を取得する為のメソッド。
func (d *Document) URL () string {
    return d.url
}
