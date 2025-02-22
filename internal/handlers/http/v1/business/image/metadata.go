package image

import (
	"mime"
	"mime/multipart"

	"github.com/timohahaa/gw/internal/errors"
)

func getFileMeta(header *multipart.FileHeader) (int64, string, string, error) {
	var (
		fileSize = header.Size
		fileExt  string
		mimeType = header.Header.Get("content-type")
	)

	if mimeType == "" {
		return 0, "", "", errors.Get(errors.ContentTypeHeaderRequired)
	}

	ext, err := mime.ExtensionsByType(mimeType)
	if err != nil {
		return 0, "", "", err
	}
	if len(ext) == 0 {
		return 0, "", "", errors.Get(errors.UsupportedFileExtension)
	}

	fileExt = getFileExt(ext)

	return fileSize, fileExt, mimeType, nil
}

// len is checked before call
func getFileExt(candidates []string) string {
	first := candidates[0]
	switch first {
	case ".jpe", "jpeg":
		first = ".jpg"
	}

	return first
}
