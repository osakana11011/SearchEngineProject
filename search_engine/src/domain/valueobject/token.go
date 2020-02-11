package valueobject

// Token は文書中に出現する文字の塊の単位。
type Token struct {
    id  int64  // トークンID
    str string // トークン文字列
}

// NewToken は新しいToken構造体を返す
func NewToken(id int64, str string) Token {
    return Token{id: id, str: str}
}

// ID はトークンIDを返す。
func (t *Token) ID() int64 {
    return t.id
}

// Str はトークン文字列を返す。
func (t *Token) Str() string {
    return t.str
}
