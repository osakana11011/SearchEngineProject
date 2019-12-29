package database

// Page ...
type Page struct {
	Title string
	Url string
}

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

// GetPages ...
func GetPages(q string) ([]Page, error) {
	// DB接続
	db, connectionErr := ConnectDB()
	if connectionErr != nil {
		return nil, connectionErr
	}
	defer db.Close()

	stmt, prepareErr := db.Prepare("SELECT title, url FROM pages WHERE title LIKE ? LIMIT 10")
	if prepareErr != nil {
		return nil, prepareErr
	}
	defer stmt.Close()

	rows, queryErr := stmt.Query("%" + q + "%")
    if queryErr != nil {
        return nil, queryErr
	}
	defer rows.Close()

	var pages []Page
	for rows.Next() {
        page := Page{}
        scanErr := rows.Scan(&page.Title, &page.Url)
        if scanErr != nil {
            return nil, scanErr
        }
        pages = append(pages, page)
	}

	return pages, nil
}

// IsRegisted ...
func IsRegisted(url string) (bool, error) {
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
func IsRegistedRecently(url string) (bool, error) {
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
