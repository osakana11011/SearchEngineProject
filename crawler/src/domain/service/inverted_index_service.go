package service

// import (
//     "search_engine_project/crawler/src/domain/model/entity"
//     "search_engine_project/crawler/src/infrastructure/persistance/datastore"
// )

// // InvertedIndexService は転置インデックスに関するサービスを提供する
// type InvertedIndexService interface {
//     Regist(*entity.InvertedIndex) error
// }

// type invertedIndexService struct {}

// // NewInvertedIndexService 転置インデックスサービスを利用する為の実体を返す
// func NewInvertedIndexService() InvertedIndexService {
//     return &invertedIndexService{}
// }

// // Regist トークン+転置リストをDBに登録する
// func (x *invertedIndexService) Regist(invertedIndex *entity.InvertedIndex) error {
//     tokenRepository := datastore.NewTokenRepository()
//     invertedListRepository := datastore.NewInvertedListRepository()

//     // トークンと転置リストをバルクインサート
//     if err := tokenRepository.BulkInsert(invertedIndex.TokenDictionary); err != nil {
//         return err
//     }
//     if err := invertedListRepository.BulkInsert(invertedIndex.InvertedList); err != nil {
//         return err
//     }

//     return nil
// }
