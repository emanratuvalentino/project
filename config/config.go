package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME                string
	DB_PASSWORD                string
	DB_PORT                    string
	DB_HOST                    string
	DB_NAME                    string
	DEFAULT_SCHEMA             string
	PUBLIC_KEY                 string
	MINIO_SERVER               string
	MINIO_BUCKET_NAME          string
	MINIO_ACCESS_KEY_ID        string
	MINIO_SECRET_ACCESS_KEY_ID string
	REFERENSI_SERVER           string
	PROFILE_PNS_API            string
	DISPATCHER_API            string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.json", &configuration)
	return configuration
}
