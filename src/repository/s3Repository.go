package repository

import (
	"bd2-backend/src/config"
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	envConfig     s3config
)

type s3config struct {
	accessKey string
	secretKey string
	hostname    string
	bucket    string
	useSSL    bool
}

func init() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	envConfig.accessKey = config.AwsS3AccessKeyId
	envConfig.secretKey = config.AwsS3SecretKey
	envConfig.hostname = config.AwsS3Hostname
	envConfig.bucket = config.AwsS3Bucket
	envConfig.useSSL = config.AwsS3UseSSL

	InfoLogger = log.New(log.Writer(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(log.Writer(), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}



func createConnection() (*minio.Client, error) {
	client, err := minio.New(envConfig.hostname, &minio.Options{
		Creds:  credentials.NewStaticV4(envConfig.accessKey, envConfig.secretKey, ""),
		Secure: envConfig.useSSL,
	})
	if err != nil {
		ErrorLogger.Println("Error creating MinIO client:", err)
		return nil, err
	}
	return client, nil
}

func PutFile(fileHandler *multipart.FileHeader, file multipart.File, path string) (string, error) {
	client, err := createConnection()
	if err != nil {
		return "", err
	}

	var size int64 = fileHandler.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)


	_, err = client.PutObject(context.Background(), envConfig.bucket, path, fileBytes, size, minio.PutObjectOptions{
		ContentType: fileType,
	})
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s/%s", envConfig.hostname, path)
	return url, nil
}
