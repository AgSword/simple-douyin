package favoriteDB

import "testing"

// import "github.com/karry-almond/packages/model"
import "log"

func TestInit(t *testing.T) {
	Init()
}

func TestNewFavorite(t *testing.T) {
	Init()
	NewFavorite(1, 1)
}

func TestCancelFavorite(t *testing.T) {
	Init()
	CancelFavorite(1, 1)
}

func TestGetFavoriteList(t *testing.T) {
	Init()
	status, videoList, err := GetFavoriteList(1)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
}
