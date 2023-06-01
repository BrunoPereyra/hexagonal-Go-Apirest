package processimage

import (
	"context"
	"fmt"
	"hexagonal/config"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func Processimage(fileHeader *multipart.FileHeader, PostImageChanel chan string, errChanel chan error) {

	fmt.Println(fileHeader, "~~~~~~~~~~~~~~~~")
	CLOUDINARY_URL := config.CLOUDINARY_URL()
	if CLOUDINARY_URL == "" {
		CLOUDINARY_URL = "cloudinary://843737794533519:CVj4VI4oIy6vCehUuRqOc27Iq-I@depcty8j1"
	}

	if fileHeader != nil {
		file, _ := fileHeader.Open()

		ctx := context.Background()
		cldService, errcloudinary := cloudinary.NewFromURL(CLOUDINARY_URL)
		if errcloudinary != nil {
			errChanel <- errcloudinary
		}
		resp, errcldService := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		if errcldService != nil || !strings.HasPrefix(resp.SecureURL, "https://") {
			errChanel <- errcldService
		}
		PostImageChanel <- resp.SecureURL
	} else {
		PostImageChanel <- ""
	}
}
