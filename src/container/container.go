package container

import (
	"github.com/edukmx/nuitee/config"
	"github.com/edukmx/nuitee/httpx"
	"github.com/edukmx/nuitee/internal/app/hotel"
	"github.com/edukmx/nuitee/internal/domain/nationality/service"
	"github.com/edukmx/nuitee/internal/infra/client"
	"github.com/edukmx/nuitee/internal/infra/repository"
	"github.com/edukmx/nuitee/internal/ui/handler"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {

	container := dig.New()
	_ = container.Provide(config.NewConfig)
	_ = container.Provide(httpx.NewServer)
	_ = container.Provide(handler.NewHotelsRatesHandler)
	_ = container.Provide(client.NewHotelBedsAdapter)
	_ = container.Provide(hotel.NewFindHotelsRatesHandler)
	_ = container.Provide(repository.NewNationalityRepository)
	_ = container.Provide(service.NewFindByCode)
	return container
}
