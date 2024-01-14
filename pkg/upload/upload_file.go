package upload

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		files, exists := form.File["file"]
		if !exists || len(files) == 0 {
			c.Set("dataFile", []string{"-"})
			return next(c)
		}

		var uploadedFiles []string
		for _, file := range files {
			allowedExts := map[string]bool{".txt": true, ".pdf": true}
			ext := filepath.Ext(file.Filename)
			if !allowedExts[ext] {
				return c.JSON(http.StatusBadRequest, "File not allowed. File with .txt or .pdf extension only")
			}

			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()

			tempFile, err := ioutil.TempFile("storages", "file-*"+ext)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer tempFile.Close()

			_, err = io.Copy(tempFile, src)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			data := tempFile.Name()
			filename := data[8+len("file-"):]
			uploadedFiles = append(uploadedFiles, filename)
		}

		c.Set("dataFile", uploadedFiles)

		return next(c)
	}
}
