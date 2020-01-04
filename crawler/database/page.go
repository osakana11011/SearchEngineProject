package database

// RegistPage ...
func RegistPage(title string, url string) error {
	// DB接続
	db, connectionErr := ConnectDB()
	if connectionErr != nil {
		return connectionErr
	}
	defer db.Close()

    // タイトルとURLを登録する
	_, registErr := db.Exec("INSERT INTO pages(title, url, created_at, updated_at) VALUES(?, ?, NOW(), NOW())", title, url)
	if registErr != nil {
		return registErr
	}

	return nil
}

// IsRegistedPage ...
func IsRegistedPage(url string) (bool, error) {
	// DB接続
	db, connectionErr := ConnectDB()
	if connectionErr != nil {
		return true, connectionErr
	}
	defer db.Close()

    // URLで検索掛けて、1件でもヒットしたら「既に登録されている」判定
	var counts int
	queryErr := db.QueryRow("SELECT COUNT(*) FROM pages WHERE url = ? LIMIT 1", url).Scan(&counts)
	if queryErr != nil {
		return true, queryErr
	}

	return (counts > 0), nil
}

// IsRegistedRecently ...
func IsRegistedPageRecently(url string) (bool, error) {
	// DB接続
	db, connectionErr := ConnectDB()
	if connectionErr != nil {
		return true, connectionErr
	}
	defer db.Close()

    // URLと現在の日時で検索を掛けて、1件でも引っ掛かったら最近登録されたということ
	var counts int
	queryErr := db.QueryRow("SELECT COUNT(*) FROM pages WHERE url = ? AND created_at >= ( NOW( ) - INTERVAL 1 DAY ) LIMIT 1", url).Scan(&counts)
	if queryErr != nil {
		return true, queryErr
	}

	return (counts > 0), nil
}
