package service

import (
	"errors"
	"example/webservice/service/model"
	"github.com/kamva/mgm/v3"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := mgm.SetDefaultConfig(nil, "test", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Panicf("bla: %v", err)
	}
	log.Info("connected to mongo!")
}

func CreateNewAlbum(aDTO *model.AlbumDTO) (*model.AlbumDTO, error) {
	//a := newAlbum(aDTO)
	a, err := store(aDTO.Model())
	if err != nil {
		log.Errorf("cant store album in storage: %v", a, err)
		return nil, err
	}
	log.Infof("stored album: %v", a)
	return a.Dto(), nil
}

func GetAllAlbums() []model.Album {
	a := &model.Album{}
	res := []model.Album{}
	mgm.Coll(a).SimpleFind(&res, bson.M{})
	return res
}

func GetAlbumById(id string) (*model.AlbumDTO, error) {
	if id == "" {
		return nil, errors.New("provided id is empty")
	}
	a, err := getById(id)
	if err != nil {
		return nil, err
	}
	return a.Dto(), nil
}

func store(a *model.Album) (*model.Album, error) {
	err := mgm.Coll(a).Create(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func getById(id string) (*model.Album, error) {
	a := &model.Album{}
	err := mgm.Coll(a).FindByID(id, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
