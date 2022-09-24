package odapi_resty

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/central182/odie/internal/adapter/outbound/common/odapi"
)

func New(appId, appKey string) odapi.Client {
	return client{
		restyClient: resty.New().
			SetBaseURL("https://od-api.oxforddictionaries.com/api/v2").
			SetHeader("app_id", appId).
			SetHeader("app_key", appKey),
	}
}

type client struct {
	restyClient *resty.Client
}

func (c client) GetEntries(req odapi.GetEntriesRequest) (odapi.GetEntriesResponse, error) {
	restyResponse, err := c.restyClient.R().
		SetResult(&odapi.GetEntriesResponse{}).
		SetError(&odapi.ErrorDescription{}).
		Get(fmt.Sprintf("entries/%s/%s", req.SourceLang, req.WordId))
	if err != nil {
		return odapi.GetEntriesResponse{}, err
	}

	if restyResponse.IsSuccess() {
		resp, ok := restyResponse.Result().(*odapi.GetEntriesResponse)
		if !ok {
			return odapi.GetEntriesResponse{}, errors.New("couldn't parse response as GetEntriesResponse")
		}
		return *resp, nil
	}

	if restyResponse.IsError() && restyResponse.StatusCode() == http.StatusNotFound {
		return odapi.GetEntriesResponse{}, odapi.ErrEntryNotFound
	}

	if restyResponse.IsError() {
		status := restyResponse.Status()

		errd, ok := restyResponse.Error().(*odapi.ErrorDescription)
		if ok && errd.Error != "" {
			return odapi.GetEntriesResponse{}, fmt.Errorf("%s: %s", status, errd.Error)
		}

		if body := restyResponse.String(); body != "" {
			return odapi.GetEntriesResponse{}, fmt.Errorf("%s: %s", status, body)
		}

		return odapi.GetEntriesResponse{}, errors.New(status)
	}

	return odapi.GetEntriesResponse{}, errors.New("failed for unknown reason")
}
