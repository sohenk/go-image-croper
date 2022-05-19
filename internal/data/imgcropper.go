package data

import (
	"context"
	"imgcropper/internal/biz"
	"imgcropper/pkg/filesystem/ftpdriver"

	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type ImgCropLog struct {
	gorm.Model
	Source    string `gorm:"size:300"`
	StorePath string `gorm:"size:255"`
	Md5Url    string `gorm:"size:255;index"`
	Width     uint64 `gorm:"index"`
	FileType  string `gorm:"size:20"`
}

type imgCropRepo struct {
	data *Data
	log  *log.Helper
}

func NewImgCropRepo(data *Data, logger log.Logger) biz.CropImgRepo {
	return &imgCropRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (i *imgCropRepo) GetMd5Url(ctx context.Context, md5url string) bool {
	//TODO implement me
	imgcroplog := new(ImgCropLog)
	rowsaffected := i.data.db.Where("md5_url = ?", md5url).Take(&imgcroplog).RowsAffected
	if rowsaffected > 0 {
		return true
	}
	return false
}

func (i *imgCropRepo) GetFileSize(ctx context.Context, md5url string, width int64) (*biz.ImgCropLog, error) {
	imgcroplog := new(ImgCropLog)
	result := i.data.db.Where("md5_url = ? and width = ?", md5url, width).Take(&imgcroplog)
	if result.RowsAffected > 0 {
		return &biz.ImgCropLog{
			Md5Url:    imgcroplog.Md5Url,
			Source:    imgcroplog.Source,
			Width:     imgcroplog.Width,
			StorePath: imgcroplog.StorePath,
			FileType:  imgcroplog.FileType,
		}, nil
	}
	return nil, result.Error
}

func (i *imgCropRepo) CreateFileLog(ctx context.Context, cropLog *biz.ImgCropLog) error {

	newcl := &ImgCropLog{
		Md5Url:    cropLog.Md5Url,
		Source:    cropLog.Source,
		Width:     cropLog.Width,
		StorePath: cropLog.StorePath,
		FileType:  cropLog.FileType,
	}
	result := i.data.db.Create(newcl)
	if result.Error != nil {
		i.log.Errorf("创建记录失败%s", result.Error)
		return errors.New(500, "错误", "创建用户失败")
	}
	return nil
}
func (i *imgCropRepo) StoreImg(ctx context.Context, filename string, imagedata []byte) (url, storepath string, err error) {

	url, storepath, err = i.data.ftp.Store(imagedata, filename)

	return url, storepath, err
}

func (i *imgCropRepo) GetImgFromDisk(ctx context.Context, storePath string) (imagedata []byte, err error) {
	return i.data.ftp.ReadFile(storePath)
}

func (i *imgCropRepo) StoreImgWithFtp(ctx context.Context, ftp *ftpdriver.FtpInfo, filename string, imagedata []byte) (url, storepath string, err error) {

	url, storepath, err = ftp.Store(imagedata, filename)

	return url, storepath, err
}

func (i *imgCropRepo) GetImgFromDiskWithFtp(ctx context.Context, ftp *ftpdriver.FtpInfo, storePath string) (imagedata []byte, err error) {
	return ftp.ReadFile(storePath)
}
