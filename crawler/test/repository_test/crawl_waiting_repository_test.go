package repository_test

import (
	"testing"
	"fmt"
	"github.com/joho/godotenv"
	"search_engine_project/crawler/src/domain/model/entity"
	"search_engine_project/crawler/src/infrastructure/persistance/datastore"
	_ "github.com/go-sql-driver/mysql"
)

// func TestInsert(t *testing.T) {
// 	// dotenvファイルを取り込む
// 	if err := godotenv.Load("./.env"); err != nil {
// 		panic(err)
// 	}
// 	datastore.MigrateAll()
// 	crawlWaitingRepository := datastore.NewCrawlWaitingRepository()

// 	crawlWaiting := entity.CrawlWaiting{URL: "https://hoge1.com", IsPriority: false, IsCrawled: false}
// 	if err := crawlWaitingRepository.Insert(crawlWaiting); err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestBulkInsert(t *testing.T) {
	// dotenvファイルを取り込む
	if err := godotenv.Load("./.env"); err != nil {
		panic(err)
	}
	datastore.MigrateAll()
	crawlWaitingRepository := datastore.NewCrawlWaitingRepository()

	var hoge []entity.CrawlWaiting
	hoge = append(hoge, entity.CrawlWaiting{URL: "https://hoge1.com", IsPriority: false})
	hoge = append(hoge, entity.CrawlWaiting{URL: "https://hoge2.com", IsPriority: false})
	hoge = append(hoge, entity.CrawlWaiting{URL: "https://hoge3.com", IsPriority: true})
	hoge = append(hoge, entity.CrawlWaiting{URL: "https://hoge4.com", IsPriority: false})
	hoge = append(hoge, entity.CrawlWaiting{URL: "https://hoge5.com", IsPriority: true})

	fmt.Println(hoge)
	crawlWaitingRepository.BulkInsert(hoge)
}

// func TestGetTopPriority(t *testing.T) {
// 	// dotenvファイルを取り込む
// 	if err := godotenv.Load("./.env"); err != nil {
// 		panic(err)
// 	}

// 	crawlWaitingRepository := datastore.NewCrawlWaitingRepository()

// 	crawlWaiting, err := crawlWaitingRepository.GetTopPriority()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(crawlWaiting)
// }