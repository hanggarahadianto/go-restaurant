package utils

import (
	"context"
	"mime/multipart"
	"restaurant/config"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadtoMenuFolder(file multipart.File, filePath string) (string, error) {
	ctx := context.Background()
	cld, err := config.SetupCloudinary()
	if err != nil {
		return "no background context", err
	}

	// create upload params
	uploadParams := uploader.UploadParams{
		PublicID:     filePath,
		ResourceType: "image",
		Folder:       config.EnvCloudUploadFolderMenu(),
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
