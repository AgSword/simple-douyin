package service

import (
	"bytes"
	"fmt"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"os"
	"github.com/AgSword/simpleDouyin/kitex_gen/publish"
	"github.com/AgSword/simpleDouyin/cmd/publish/model"
	user "github.com/AgSword/simpleDouyin/dal/mysql"
	
	"github.com/AgSword/simpleDouyin/cmd/publish/model/db"
	"github.com/AgSword/simpleDouyin/cmd/publish/pkg/minio"
	"strings"
	"time"
)

//"bytes"
//"context"
//"fmt"
//"image"
//"image/jpeg"
//"os"
//"strings"
//
//"github.com/a76yyyy/tiktok/dal/db"
//"github.com/a76yyyy/tiktok/kitex_gen/publish"
//"github.com/a76yyyy/tiktok/pkg/minio"
//"github.com/a76yyyy/tiktok/pkg/ttviper"
//"github.com/gofrs/uuid"
//ffmpeg "github.com/u2takey/ffmpeg-go"

func PublishAction(req *publish.PublishActionRequest, uid int) (err error) {
	// 用minio获取url
	MinioVideoBucketName := minio.MinioVideoBucketName
	videoData := []byte(req.Data)
	reader := bytes.NewReader(videoData)

	u2, err := uuid.NewV4()
	if err != nil {
		return err
	}
	fileName := u2.String() + "." + "mp4"

	// 上传视频
	err = minio.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		return err
	}
	// 获取视频链接
	url, err := minio.GetFileUrl(MinioVideoBucketName, fileName, 0)
	playUrl := strings.Split(url.String(), "?")[0]
	if err != nil {
		return err
	}

	u3, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// 获取封面
	coverPath := u3.String() + "." + "jpg"
	coverData, err := readFrameAsJpeg(playUrl)
	if err != nil {
		return err
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadFile(MinioVideoBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return err
	}

	// 获取封面链接
	coverUrl, err := minio.GetFileUrl(MinioVideoBucketName, coverPath, 0)
	if err != nil {
		return err
	}

	CoverUrl := strings.Split(coverUrl.String(), "?")[0]

	videoModel := &model.Video{
		UserId:        int64(uid),
		Title:         req.Title,
		PlayUrl:       playUrl,
		CoverUrl:      CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	tx := db.Db

	err = tx.Model(&model.Video{}).Create(&videoModel).Error
	if err != nil {
		return err
	}
	return nil
}

// ReadFrameAsJpeg 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}

func PublishList(req *publish.PublishListRequest) ([]*publish.Video, error) {
	var tbVideoList []*model.Video
	tx := db.Db
	err := tx.Where("user_id = ?", req.UserId).Find(&tbVideoList).Error
	//err := tx.Model(&model.Video{UserId: req.UserId}).Find(&tbVideoList).Error
	if err != nil {
		return nil, err
	}

	var list []*publish.Video
	for _, video := range tbVideoList {
		var user user.User
		// if err := tx.Model(&model.User{ID: video.UserId}).Find(&user); err != nil {
		// 	return nil, err.Error
		// }
		// res := &User{}
		if err := tx.Where("id = ?", video.UserId).Find(&user).Error; err != nil {
			return nil, err
		}
		publishUser:=&publish.User{Id: user.ID, Name: user.Name,FollowCount: &user.FollowCount,
            FollowerCount: &user.FollowerCount,IsFollow: false}
		respVideo := &publish.Video{
			Id:            video.ID,
			Author:        publishUser,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.Title,
		}
		list = append(list, respVideo)

	}
	return list, nil
}
