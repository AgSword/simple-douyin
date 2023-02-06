package dal

import (
	"github.com/AgSword/simpleDouyin/cmd/user/biz/dal/mysql"
	"github.com/AgSword/simpleDouyin/cmd/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
