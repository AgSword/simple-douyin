package relationDB

import "testing"

// import "github.com/karry-almond/packages/model"
import "log"

func TestInit(t *testing.T) {
	Init()
}

func TestNewRelation(t *testing.T) {
	Init()
	NewRelation(1, 2)
}

func TestCancelFavorite(t *testing.T) {
	Init()
	CancelRelation(1, 2)
}

func TestGetFollowList(t *testing.T) {
	Init()
	status, videoList, err := GetFollowList(1)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
}
func TestGetFollowerList(t *testing.T) {
	Init()
	status, videoList, err := GetFollowList(1)
	log.Println(status)
	log.Println(videoList)
	log.Println(err)
}
