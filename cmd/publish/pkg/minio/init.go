package minio

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client
	//Config               = ttviper.ConfigInit("TIKTOK_MINIO", "minioConfig")
	MinioEndpoint        = "47.100.224.26:9000"
	MinioAccessKeyId     = "minioadmin"
	MinioSecretAccessKey = "minioadmin"
	MinioUseSSL          = false
	MinioVideoBucketName = "tiktok-video"
)

//# Endpoint: 配置为运行minio服务的主机地址
//Minio:
//Endpoint: localhost:9000
//AccessKeyId: tiktokMinio
//SecretAccessKey: tiktokMinio
//UseSSL: false
//VideoBucketName : tiktok-video

// Minio 对象存储初始化
func init() {
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
		Secure: MinioUseSSL,
	})
	if err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
	// fmt.Println(client)
	klog.Debug("minio client init successfully")
	minioClient = client
	if err := CreateBucket(MinioVideoBucketName); err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
}
