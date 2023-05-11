package data

import (
	"context"
	"fmt"
	"imgcropper/internal/biz"
	"imgcropper/pkg/filesystem/ftpdriver"
	"imgcropper/pkg/rediscache"

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
	if i.GetMd5UrlFromCache(ctx, md5url) {
		return true
	}
	imgcroplog := new(ImgCropLog)
	rowsaffected := i.data.db.Where("md5_url = ?", md5url).Take(&imgcroplog).RowsAffected
	return rowsaffected > 0
}
func (i *imgCropRepo) GetMd5UrlFromCache(ctx context.Context, md5url string) bool {
	key := fmt.Sprintf("go_imagecropper:ismd5url:%s", md5url)
	cacheresult, err := rediscache.Get[int](ctx, i.data.Rdb, key)
	if err == nil {
		return false
	}
	if cacheresult == 1 {
		return true
	}

	return false
}
func (i *imgCropRepo) SetMd5UrlFromCache(ctx context.Context, md5url string) error {
	key := fmt.Sprintf("go_imagecropper:ismd5url:%s", md5url)
	err := rediscache.Set(ctx, i.data.Rdb, key, 1, 0)

	return err
}

func (i *imgCropRepo) GetFileSize(ctx context.Context, md5url string, width int64) (*biz.ImgCropLog, error) {
	cacheresult, err := i.GetFileSizeFromCache(ctx, md5url, width)
	// i.log.Debug("cacheresultxxx:", cacheresult)
	if err == nil {
		return cacheresult, nil
	}

	imgcroplog := new(ImgCropLog)
	result := i.data.db.Where("md5_url = ? and width = ?", md5url, width).Take(&imgcroplog)
	if result.RowsAffected > 0 {
		rr := &biz.ImgCropLog{
			Md5Url:    imgcroplog.Md5Url,
			Source:    imgcroplog.Source,
			Width:     imgcroplog.Width,
			StorePath: imgcroplog.StorePath,
			FileType:  imgcroplog.FileType,
		}
		i.SetMd5UrlFromCache(ctx, imgcroplog.Md5Url)
		i.SetFileSizeFromCache(ctx, rr)
		return rr, nil
	}
	return nil, result.Error
}
func (i *imgCropRepo) GetFileSizeFromCache(ctx context.Context, md5url string, width int64) (*biz.ImgCropLog, error) {
	key := fmt.Sprintf("go_imagecropper:md5url:%s:width:%d", md5url, width)
	cacheresult, err := rediscache.Get[biz.ImgCropLog](ctx, i.data.Rdb, key)
	if err != nil {
		return nil, err
	}
	i.log.Debug("cacheresult:xxxxx:", cacheresult)
	return &cacheresult, nil

}
func (i *imgCropRepo) SetFileSizeFromCache(ctx context.Context, cropLog *biz.ImgCropLog) error {
	key := fmt.Sprintf("go_imagecropper:md5url:%s:width:%d", cropLog.Md5Url, cropLog.Width)
	err := rediscache.Set(ctx, i.data.Rdb, key, cropLog, 0)
	return err
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
	i.SetFileSizeFromCache(ctx, cropLog)
	i.SetMd5UrlFromCache(ctx, cropLog.Md5Url)
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
