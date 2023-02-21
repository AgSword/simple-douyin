package dal

import (
	"github.com/AgSword/simpleDouyin/cmd/feed/biz/dal/mysql"
	"github.com/AgSword/simpleDouyin/cmd/feed/biz/dal/redis"
	//github.com/AgSword/simpleDouyin
)

func Init() {
	redis.Init()
	mysql.Init()
}
