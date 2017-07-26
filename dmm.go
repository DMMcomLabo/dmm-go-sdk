package dmm

import (
	"github.com/dmmlabo/dmm-go-sdk/api"
)

// Client for DMM services
//
// DMMサービス接続のためのClient
type Client struct {
	AffiliateID string
	ApiID       string

	// services
	Actress *api.ActressService
	Author  *api.AuthorService
	Floor   *api.FloorService
	Genre   *api.GenreService
	Maker   *api.MakerService
	Product *api.ProductService
	Series  *api.SeriesService
}

// New creates client
//
// Clientの新規作成
func New(affiliateID, apiID string) *Client {
	c := &Client{
		ApiID:       apiID,
		AffiliateID: affiliateID,
	}
	c.Actress = api.NewActressService(affiliateID, apiID)
	c.Author = api.NewAuthorService(affiliateID, apiID)
	c.Floor = api.NewFloorService(affiliateID, apiID)
	c.Genre = api.NewGenreService(affiliateID, apiID)
	c.Maker = api.NewMakerService(affiliateID, apiID)
	c.Product = api.NewProductService(affiliateID, apiID)
	c.Series = api.NewSeriesService(affiliateID, apiID)
	return c
}
