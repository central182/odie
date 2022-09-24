package dependency_cli

import (
	"github.com/central182/odie/internal/adapter/inbound/cli"
	odapi_resty "github.com/central182/odie/internal/adapter/outbound/common/odapi/resty"
	dictionary_service_odapi "github.com/central182/odie/internal/adapter/outbound/dictionary/dictionary_service/odapi"
	"github.com/central182/odie/internal/domain/application"
)

func InitInitApplication(appId, appKey string) cli.InitApplication {
	return func() application.Application {
		return application.New(
			dictionary_service_odapi.New(
				odapi_resty.New(appId, appKey),
			),
		)
	}
}
