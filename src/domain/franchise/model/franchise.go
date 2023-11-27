package model

import locationTypes "github.com/flexuxs/clubHubApp/src/domain/location/model"

type Franchise struct {
	Id       string                 `json:"id"`
	Name     string                 `json:"name"`
	URL      string                 `json:"url"`
	Location locationTypes.Location `json:"location"`
}
