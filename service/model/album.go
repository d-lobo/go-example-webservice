package model

import "github.com/kamva/mgm/v3"

type Album struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string  `json:"title"`
	Artist           string  `json:"artist"`
	Price            float64 `json:"price"`
}

type AlbumDTO struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (a *Album) Dto() *AlbumDTO {
	return &AlbumDTO{
		Artist: a.Artist,
		Title:  a.Title,
		Price:  a.Price,
	}
}

func (a *AlbumDTO) Model() *Album {
	return &Album{
		Artist: a.Artist,
		Title:  a.Title,
		Price:  a.Price,
	}
}
