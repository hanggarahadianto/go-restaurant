package config

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cldName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cldKey := os.Getenv("CLOUDINARY_API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}

func EnvCloudUploadFolderRestaurant() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}
	return os.Getenv("CLOUDINARY_MAIN_FOLDER")
}
func EnvCloudUploadFolderMenu() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}
	return os.Getenv("CLOUDINARY_MENU_FOLDER")
}
