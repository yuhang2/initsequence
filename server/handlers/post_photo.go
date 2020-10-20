package handlers

import (
	"github.com/yuhang2/initsequence/external/aws"
)

func UploadPhoto() {
	s3.S3Client.Upload("whocare")
}
