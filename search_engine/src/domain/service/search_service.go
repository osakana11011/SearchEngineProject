package service

import (
    "fmt"
    "search_engine_project/search_engine/src/domain/model/entity"
    "search_engine_project/search_engine/src/domain/model/valueobject"
    "search_engine_project/search_engine/src/domain/repository"
)

// SearchService は文書に関する様々な処理を呼び出す為の窓口。
type SearchService interface {
    Search(title string) ([]entity.Document, error)
}

type searchService struct {
    tokenRepo repository.TokenRepository
    documentRepo repository.DocumentRepository
    invertedDataRepo repository.InvertedDataRepository
}

// NewSearchService はSearchServiceを扱うインスタンスを提供する。
func NewSearchService(
    tokenRepo repository.TokenRepository,
    documentRepo repository.DocumentRepository,
    invertedDataRepo repository.InvertedDataRepository,
    ) SearchService {

    return &searchService{
        tokenRepo: tokenRepo,
        documentRepo: documentRepo,
        invertedDataRepo: invertedDataRepo,
    }
}

func (ss *searchService) Search(q string) ([]entity.Document, error) {
    query := valueobject.NewQueryFromString(q)

    tokens, err := ss.tokenRepo.GetByTokenNames(query.QueryStrings)
    if err != nil {
        return []entity.Document{}, err
    }

    var invertedLists [][]entity.InvertedData
    for _, token := range tokens {
        invertedList, _ := ss.invertedDataRepo.GetByToken(token)
        invertedLists = append(invertedLists, invertedList)
    }

    if len(invertedLists) == 0 {
        return []entity.Document{}, nil
    }

    documentIDs := ranking(invertedLists)

    documents, _ := ss.documentRepo.GetByIDs(documentIDs)

    return documents, nil
}

func ranking(invertedLists [][]entity.InvertedData) []uint {
    // 1番目
    // 1番目の転置データに入っている文書IDは全て検索結果候補
    var candidateDocumentIDs []uint
    for _, invertedData := range invertedLists[0] {
        candidateDocumentIDs = append(candidateDocumentIDs, invertedData.DocumentID)
    }

    // 2〜N番目
    // 2〜N番目については、それまでに候補に入っていて、かつ転置リストにも含まれていることが条件
    for _, invertedList := range invertedLists[1:] {
        var tmpCandidateDocumentIDs []uint
        for _, invertedData := range invertedList {
            for _, candidateDocumentID := range candidateDocumentIDs {
                if invertedData.DocumentID == candidateDocumentID {
                    tmpCandidateDocumentIDs = append(tmpCandidateDocumentIDs, invertedData.DocumentID)
                    break
                }
            }
        }
        candidateDocumentIDs = tmpCandidateDocumentIDs
    }

    if len(candidateDocumentIDs) == 0 {
        return []uint{}
    }

    fmt.Println(candidateDocumentIDs)

    return candidateDocumentIDs[:10]
}
