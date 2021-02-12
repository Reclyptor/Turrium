package env

import (
	"log"
	"os"
)

var (
	AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_ACCESS_KEY_SECRET = os.Getenv("AWS_ACCESS_KEY_SECRET")
	AWS_REGION = os.Getenv("AWS_REGION")
	MONGO_HOST = os.Getenv("MONGO_HOST")
	MONGO_USERNAME = os.Getenv("MONGO_USERNAME")
	MONGO_PASSWORD = os.Getenv("MONGO_PASSWORD")
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
	MONGO_ALBUM_COLLECTION = os.Getenv("MONGO_ALBUM_COLLECTION")
	MONGO_IMAGE_COLLECTION = os.Getenv("MONGO_IMAGE_COLLECTION")
	MONGO_VIDEO_COLLECTION = os.Getenv("MONGO_VIDEO_COLLECTION")
	MONGO_USER_COLLECTION = os.Getenv("MONGO_USER_COLLECTION")
	OAUTH_CLIENT_ID = os.Getenv("OAUTH_CLIENT_ID")
	OAUTH_CLIENT_SECRET = os.Getenv("OAUTH_CLIENT_SECRET")
	OAUTH_CLIENT_ISSUER = os.Getenv("OAUTH_CLIENT_ISSUER")
	TURRIUM_UI_URL = os.Getenv("TURRIUM_UI_URL")
)

func Verify() {
	if AWS_ACCESS_KEY_ID == "" {
		log.Fatal("Environment variable not found: AWS_ACCESS_KEY_ID")
	}
	if AWS_ACCESS_KEY_SECRET == "" {
		log.Fatal("Environment variable not found: AWS_ACCESS_KEY_SECRET")
	}
	if AWS_REGION == "" {
		log.Fatal("Environment variable not found: AWS_REGION")
	}
	if MONGO_HOST == "" {
		log.Fatal("Environment variable not found: MONGO_HOST")
	}
	if MONGO_USERNAME == "" {
		log.Fatal("Environment variable not found: MONGO_USERNAME")
	}
	if MONGO_PASSWORD == "" {
		log.Fatal("Environment variable not found: MONGO_PASSWORD")
	}
	if MONGO_DATABASE == "" {
		log.Fatal("Environment variable not found: MONGO_DATABASE")
	}
	if MONGO_ALBUM_COLLECTION == "" {
		log.Fatal("Environment variable not found: MONGO_ALBUM_COLLECTION")
	}
	if MONGO_IMAGE_COLLECTION == "" {
		log.Fatal("Environment variable not found: MONGO_IMAGE_COLLECTION")
	}
	if MONGO_VIDEO_COLLECTION == "" {
		log.Fatal("Environment variable not found: MONGO_VIDEO_COLLECTION")
	}
	if MONGO_USER_COLLECTION == "" {
		log.Fatal("Environment variable not found: MONGO_USER_COLLECTION")
	}
	if OAUTH_CLIENT_ID == "" {
		log.Fatal("Environment variable not found: OAUTH_CLIENT_ID")
	}
	if OAUTH_CLIENT_SECRET == "" {
		log.Fatal("Environment variable not found: OAUTH_CLIENT_SECRET")
	}
	if OAUTH_CLIENT_ISSUER == "" {
		log.Fatal("Environment variable not found: OAUTH_CLIENT_ISSUER")
	}
	if TURRIUM_UI_URL == "" {
		log.Fatal("Environment variable not found: TURRIUM_UI_URL")
	}
}