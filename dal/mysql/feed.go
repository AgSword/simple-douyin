package mysql

import (
	"context"
	"time"
)

const TableNameFeed = "video"

type Video struct {
	ID            int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId        int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Title         string    `gorm:"column:title" db:"title" json:"title" form:"title"`
	PlayUrl       string    `gorm:"column:play_url" db:"play_url" json:"play_url" form:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url" db:"cover_url" json:"cover_url" form:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count" db:"favorite_count" json:"favorite_count" form:"favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count" db:"comment_count" json:"comment_count" form:"comment_count"`
	CreatedAt     time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}

// TableName User's table name
func (*Video) TableName() string {
	return TableNameFeed
}

func GetUserFeeds(ctx context.Context, latestTime *int64) ([]*Video, error) {
	res := make([]*Video, 0)

	if err := DB.WithContext(ctx).
		Where("created_at < ?", latestTime).
		Order("created_at asc").
		//Limit(config.FeedListLength).
		Find(&res).
		Error; err != nil {
		return nil, err	
	}
	return res, nil

}
