package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64          `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId    int64          `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	VoideoId  int64          `gorm:"column:voideo_id" db:"voideo_id" json:"voideo_id" form:"voideo_id"`
	Contents  string         `gorm:"column:contents" db:"contents" json:"contents" form:"contents"`
	CreatedAt time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}

type Favorite struct {
	ID        int64          `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId    int64          `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	VideoId   int64          `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	CreatedAt time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}

type Relation struct {
	ID        int64          `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId    int64          `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	ToUserId  int64          `gorm:"column:to_user_id" db:"to_user_id" json:"to_user_id" form:"to_user_id"`
	CreatedAt time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}

type User struct {
	ID            int64          `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name          string         `gorm:"column:name" db:"name" json:"name" form:"name"`
	Password      string         `gorm:"column:password" db:"password" json:"password" form:"password"`
	FollowCount   int64          `gorm:"column:follow_count" db:"follow_count" json:"follow_count" form:"follow_count"`
	FollowerCount int64          `gorm:"column:follower_count" db:"follower_count" json:"follower_count" form:"follower_count"`
	CreatedAt     time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}

type Video struct {
	ID            int64          `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId        int64          `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Title         string         `gorm:"column:title" db:"title" json:"title" form:"title"`
	PlayUrl       string         `gorm:"column:play_url" db:"play_url" json:"play_url" form:"play_url"`
	CoverUrl      string         `gorm:"column:cover_url" db:"cover_url" json:"cover_url" form:"cover_url"`
	FavoriteCount int64          `gorm:"column:favorite_count" db:"favorite_count" json:"favorite_count" form:"favorite_count"`
	CommentCount  int64          `gorm:"column:comment_count" db:"comment_count" json:"comment_count" form:"comment_count"`
	CreatedAt     time.Time      `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}
