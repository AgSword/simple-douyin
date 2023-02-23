package dal

import (
	"github.com/AgSword/simpleDouyin/dal/mysql"
	"github.com/AgSword/simpleDouyin/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()

}
