package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Validate file extension
		ext := filepath.Ext(file.Filename)
		if ext != ".png" && ext != ".jpg" {
			return c.JSON(http.StatusBadRequest, "Invalid file format. Only PNG and JPG are allowed.")
		}

		// Validate file size
		if file.Size > 100*1024 { // 100KB
			return c.JSON(http.StatusBadRequest, "File size exceeds the limit of 100KB.")
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		var ctx = context.Background()
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		// Add your Cloudinary credentials ...
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

		// Upload file to Cloudinary ...
		resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "HomeTest"})

		if err != nil {
			fmt.Println(err.Error())
		}

		c.Set("dataFile", resp.SecureURL)
		return next(c)

		// tempFile, err := os.Create(filepath.Join("uploads", "image-*"+ext))
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, err)
		// }
		// defer tempFile.Close()

		// if _, err = io.Copy(tempFile, src); err != nil {
		// 	return c.JSON(http.StatusBadRequest, err)
		// }

		// data := tempFile.Name()

		// c.Set("dataFile", data)
		// return next(c)
	}
}
