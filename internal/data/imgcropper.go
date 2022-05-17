package data

import (
	"imgcropper/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type imgCropRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewImgCropRepo(data *Data, logger log.Logger) biz.CropImgRepo {
	return &imgCropRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
