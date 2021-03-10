package storage

import (
	"code-be-docudigital/config"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() *minio.Client {

	// Initialize minio client object.
	configuration := config.GetConfig()

	minioClient, err := minio.New(configuration.MINIO_SERVER, &minio.Options{
		Creds:  credentials.NewStaticV4(configuration.MINIO_ACCESS_KEY_ID, configuration.MINIO_SECRET_ACCESS_KEY_ID, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return minioClient
}
