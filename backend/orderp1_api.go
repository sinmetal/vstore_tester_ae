package backend

import (
	"context"
	"net/http"

	"github.com/favclip/ucon"
	"github.com/favclip/ucon/swagger"
)

func setupOrderP1(swPlugin *swagger.Plugin) {
	api := &OrderP1API{}
	tag := swPlugin.AddTag(&swagger.Tag{Name: "OrderP1", Description: "OrderP1 list"})
	var hInfo *swagger.HandlerInfo

	hInfo = swagger.NewHandlerInfo(api.Post)
	ucon.Handle(http.MethodPost, "/orderP1", hInfo)
	hInfo.Description, hInfo.Tags = "post to orderP1", []string{tag.Name}

	hInfo = swagger.NewHandlerInfo(api.Get)
	ucon.Handle(http.MethodGet, "/orderP1", hInfo)
	hInfo.Description, hInfo.Tags = "get to orderP1", []string{tag.Name}
}

type OrderP1API struct{}

type OrderP1APIPostRequest struct {
}

type OrderP1APIPostResponse struct {
}

func (api *OrderP1API) Post(c context.Context, form *OrderP1APIPostRequest) (*OrderP1APIPostResponse, error) {
	return nil, nil
}

type OrderP1APIGetRequest struct {
}

type OrderP1APIGetResponse struct {
}

func (api *OrderP1API) Get(c context.Context, form *OrderP1APIGetRequest) (*OrderP1APIGetResponse, error) {
	return &OrderP1APIGetResponse{}, nil
}
