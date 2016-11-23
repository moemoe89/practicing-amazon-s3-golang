package services

import (
	"github.com/minio/minio-go"
	"os"
)

func AwsUpload(path,filename string) (error){

	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	s3Client, err := minio.New("s3.amazonaws.com", accessKeyID, secretAccessKey, true)
	if err != nil {
		return err
	}

	object, err := os.Open(path +"/"+ filename)
	if err != nil {
		return err
	}
	defer object.Close()

	_, err = s3Client.PutObject(os.Getenv("AWS_BUCKET"), filename, object, "image/png")
	if err != nil {
		return err
	}

	return nil
}