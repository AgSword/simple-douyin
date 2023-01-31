package mysql

import (
	"context"
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID             int32     `gorm:"column:id;primaryKey" json:"id"`                // 主键id
	UserName       string    `gorm:"column:user_name" json:"user_name"`             // 用户名
	Password       string    `gorm:"column:password" json:"password"`               // 密码
	FollowingCount int32     `gorm:"column:following_count" json:"following_count"` // 关注数
	FollowerCount  int32     `gorm:"column:follower_count" json:"follower_count"`   // 被关注数
	UpdateAt       time.Time `gorm:"column:update_at" json:"update_at"`             // 更新时间
	CreateAt       time.Time `gorm:"column:create_at" json:"create_at"`             // 创建时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

func GetUserByUsername(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserByIds(ctx context.Context, userIDs []int32) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}
