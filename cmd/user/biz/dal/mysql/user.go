package mysql

import (
	"context"
	// "log"
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID            int64          `gorm:"column:id" json:"id"`
	Name          string         `gorm:"column:name" json:"name"`
	Password      string         `gorm:"column:password" json:"password"`
	FollowCount   int64          `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int64          `gorm:"column:follower_count" json:"follower_count"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;not null" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

func GetUserByUsername(ctx context.Context, userName string) ([]*User, error) {
	// log.Println(DB==nil)
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserByIds(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CreateUser(ctx context.Context, users *User) error {
	println(ctx == nil)
	println(DB == nil)
	return DB.WithContext(ctx).Create(users).Error
}

func UpdateUserPassword(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Model(user).Where("id", user.ID).Update("password", user.Password).Error
}
