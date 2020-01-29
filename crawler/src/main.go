package main

import (
    "os"
    "fmt"
    "search_engine_project/crawler/src/domain/model/entity"
    "search_engine_project/crawler/src/domain/repository"
    "search_engine_project/crawler/src/domain/service"
    "search_engine_project/crawler/src/usecase"
    "search_engine_project/crawler/src/infrastructure/persistance/datastore"
    "github.com/joho/godotenv"
    "go.uber.org/dig"
    _ "github.com/go-sql-driver/mysql"
)

func init () {
    // 環境変数ENVを参照し、それに応じた環境変数ファイルを取り込む
    env := os.Getenv("ENV")
    if env != "" {
        if err := godotenv.Load(fmt.Sprintf("./.envfiles/%s.env", env)); err != nil {
            panic(err)
        }
    }

    // 自動マイグレーション
    if err := datastore.MigrateAll(); err != nil {
        panic(err)
    }
}

func main() {
    c := dig.New()

    c.Provide(datastore.NewGormDBConnection)

    c.Provide(datastore.NewCrawlWaitingRepository)
    c.Provide(datastore.NewDocumentRepository)
    c.Provide(datastore.NewDomainRepository)
    c.Provide(datastore.NewInvertedDataRepository)
    c.Provide(datastore.NewTokenRepository)

    c.Provide(service.NewCrawlWaitingService)
    c.Provide(service.NewCrawlService)
    c.Provide(usecase.NewCrawlUsecase)

    c.Invoke(func(crawlWaitingRepository repository.CrawlWaitingRepository) {
        crawlWaiting := entity.CrawlWaiting{URL: "https://ja.wikipedia.org/wiki/Google", IsPriority: true}
        crawlWaitingRepository.Insert(crawlWaiting)
    })

    c.Invoke(func(crawlUsecase usecase.CrawlUsecase) {
        crawlUsecase.ExecCrawlService()
    })
}
