package database

// RegistWord ...
func RegistWord(word string, partOfSpeech string) error {
	// DB接続
	db, connectionErr := ConnectDB()
	if connectionErr != nil {
		return connectionErr
	}
	defer db.Close()

	// タイトルとURLを登録する
	_, registErr := db.Exec("INSERT INTO words(word, part_of_speech, created_at, updated_at) VALUES(?, ?, NOW(), NOW())", word, partOfSpeech)
	if registErr != nil {
		return registErr
	}

	return nil
}

// IsRegistedWord ...
func IsRegistedWord(word string) (bool, error) {
	// DB接続
	db, connectionErr := ConnectDB()
	if connectionErr != nil {
		return true, connectionErr
	}
	defer db.Close()

    // URLで検索掛けて、1件でもヒットしたら「既に登録されている」判定
	var counts int
	queryErr := db.QueryRow("SELECT COUNT(*) FROM words WHERE word = ? LIMIT 1", word).Scan(&counts)
	if queryErr != nil {
		return true, queryErr
	}

	return (counts > 0), nil
}
