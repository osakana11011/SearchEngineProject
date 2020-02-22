package service

import (
    "search_engine_project/search_engine/src/domain/model/entity"
    "search_engine_project/search_engine/src/domain/model/valueobject"
    "search_engine_project/search_engine/src/domain/repository"
)

// SearchService は文書に関する様々な処理を呼び出す為の窓口。
type SearchService interface {
    Search(q string, page int) (entity.SearchResult, error)
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

func (ss *searchService) Search(q string, page int) (entity.SearchResult, error) {
    query := valueobject.NewQueryFromString(q)

    tokens, err := ss.tokenRepo.GetByTokenNames(query.QueryStrings)
    if err != nil {
        return entity.SearchResult{}, err
    }

    var invertedLists [][]entity.InvertedData
    for _, token := range tokens {
        invertedList, _ := ss.invertedDataRepo.GetByToken(token)
        invertedLists = append(invertedLists, invertedList)
    }

    if len(invertedLists) == 0 {
        return entity.SearchResult{}, nil
    }

    documentsN, documentIDs := ranking(invertedLists, page)

    documents, _ := ss.documentRepo.GetByIDs(documentIDs)

    return entity.NewSearchResult(q, documents, documentsN, page, 10), nil
}

func ranking(invertedLists [][]entity.InvertedData, page int) (int, []uint) {
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

    documentsN := len(candidateDocumentIDs)

    if len(candidateDocumentIDs) < 10 {
        return documentsN, candidateDocumentIDs
    }

    return documentsN, candidateDocumentIDs[((page-1)*10):(page*10)]
}
