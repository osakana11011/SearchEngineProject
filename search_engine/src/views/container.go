package views

import (
	"search_engine_project/search_engine/src/usecase"
    "search_engine_project/search_engine/src/domain/service"
	"search_engine_project/search_engine/src/infrastructure/persistance/datastore"

	"go.uber.org/dig"
)

func getNewContainer() *dig.Container {
	c := dig.New()

    c.Provide(datastore.NewGormDBConnection)
    c.Provide(datastore.NewTokenRepository)
    c.Provide(datastore.NewDocumentRepository)
    c.Provide(datastore.NewInvertedDataRepository)

    c.Provide(service.NewSearchService)
	c.Provide(usecase.NewSearchUseCase)

	return c
}
