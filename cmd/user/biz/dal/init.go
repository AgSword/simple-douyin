package dal

import (
	"github.com/AgSword/simple-douyin/cmd/user/biz/dal/mysql"
	"github.com/AgSword/simple-douyin/cmd/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
