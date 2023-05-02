package biz

import (
	"context"
	"crypto/md5"
	"fmt"
	pb "imgcropper/api/imgcropper/service"
	"imgcropper/internal/conf"
	"imgcropper/pkg/filesystem/ftpdriver"
	"imgcropper/pkg/imagehelper"
	"path"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type CropImg struct {
	url   string
	width int64
}
type ImgCropLog struct {
	Source    string
	StorePath string
	Md5Url    string
	Width     uint64
	FileType  string
}

type CropImgRepo interface {
	GetMd5Url(ctx context.Context, md5url string) bool
	GetFileSize(ctx context.Context, md5url string, width int64) (*ImgCropLog, error)
	CreateFileLog(ctx context.Context, cropLog *ImgCropLog) error
	StoreImg(ctx context.Context, filename string, imagedata []byte) (url, storepath string, err error)
	GetImgFromDisk(ctx context.Context, storePath string) (imagedata []byte, err error)

	StoreImgWithFtp(ctx context.Context, ftp *ftpdriver.FtpInfo, filename string, imagedata []byte) (url, storepath string, err error)
	GetImgFromDiskWithFtp(ctx context.Context, ftp *ftpdriver.FtpInfo, storePath string) (imagedata []byte, err error)

	GetFileSizeFromCache(ctx context.Context, md5url string, width int64) (*ImgCropLog, error)
	SetFileSizeFromCache(ctx context.Context, cropLog *ImgCropLog) error
	SetMd5UrlFromCache(ctx context.Context, md5url string) error
	GetMd5UrlFromCache(ctx context.Context, md5url string) bool
}

type CropImgUsecase struct {
	fileconf *conf.FileSystem
	repo     CropImgRepo
	log      *log.Helper
}

func NewCropImgUsecase(repo CropImgRepo, fileconf *conf.FileSystem, logger log.Logger) *CropImgUsecase {
	return &CropImgUsecase{
		fileconf: fileconf,
		repo:     repo,
		log:      log.NewHelper(log.With(logger, "module", "usecase/imgcrop")),
	}
}

func (uc *CropImgUsecase) CropImgBiz(ctx context.Context, url string, width int64) (*pb.CropImgReply, error) {
	//fmt.Println(url, width)
	md5filename := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	baseName := path.Base(url)
	ext := path.Ext(baseName)                 // 输出 .jpg
	name := strings.TrimSuffix(baseName, ext) // 输出 name
	newname := name + "_w" + "_ " + strconv.Itoa(int(width)) + ext

	oname := name + ext
	ftpinfo, ftperr := ftpdriver.NewFtpInfo(uc.fileconf.Ftp.Host, uc.fileconf.Ftp.Username, uc.fileconf.Ftp.Password, uc.fileconf.Ftp.Root, uc.fileconf.Ftp.Url, uint(uc.fileconf.Ftp.Port), uc.fileconf.Ftp.Dir)
	if ftperr != nil {
		uc.log.Error("ftp连接失败")
		uc.log.Error(ftperr)
		noftpbyte, noftpfiletype, err := imagehelper.ResizeImgToByte(url, width)
		if err != nil {
			return nil, err
		}
		return &pb.CropImgReply{Imgdata: noftpbyte, Imgname: newname, Imagetype: noftpfiletype}, nil

	}
	defer ftpinfo.Conn.Quit()

	//uc.log.Debug(md5filename)
	if uc.repo.GetMd5Url(ctx, md5filename) {

		//数据库是否记录过该缩略图
		imgcroplog, err := uc.repo.GetFileSize(ctx, md5filename, width)
		var obyte []byte
		//没有对应的缩略图
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				uc.log.Error(err)
				return nil, errors.New(500, "dberror", "数据库出错了")
			}
			//记录不存在

			//uc.log.Debug("没缩略图")
		} else {
			//uc.log.Debug("有缩略图")
			//有对应缩略图
			storebyte, err := uc.repo.GetImgFromDiskWithFtp(ctx, ftpinfo, imgcroplog.StorePath)
			if err != nil {
				uc.log.Error(err)
			}
			if storebyte != nil {
				return &pb.CropImgReply{
					Imgdata: storebyte, Imgname: newname, Imagetype: imgcroplog.FileType,
				}, nil
			}
		}

		//uc.log.Debug("开始创建缩略图")
		//原图是否在本地有副本
		originimgcroplog, err := uc.repo.GetFileSize(ctx, md5filename, 0)
		if err != nil {
			//uc.log.Error(err)
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New(500, "dberror", "数据库出错了")
			}
			//原图不存在创建原图
			// newimg, filetype, err := imagehelper.GetImgFromUrl(url)
			// if err != nil {
			// 	uc.log.Error(err)
			// 	return nil, err
			// }

			// newbyte, err := imagehelper.TransferImgToByte(newimg, filetype)
			// obyte = newbyte

			obyte, filetype, err := imagehelper.GetImgFromUrlToBytes(url)
			if err == nil {
				_, ostorepath, err := uc.repo.StoreImgWithFtp(ctx, ftpinfo, oname, obyte)
				originimgcroplog = &ImgCropLog{
					Source:    url,
					StorePath: ostorepath,
					Md5Url:    md5filename,
					Width:     0,
					FileType:  filetype,
				}
				//上传成功则保存原图记录到数据库
				if err == nil {
					uc.repo.CreateFileLog(ctx, originimgcroplog)
				} else {

					uc.log.Error(err)
				}
			}

			uc.log.Debug("原图服务器不存在")
		} else {

			uc.log.Debug("原图服务器存在")
			diskbyte, err := uc.repo.GetImgFromDiskWithFtp(ctx, ftpinfo, originimgcroplog.StorePath)
			if err != nil {
				uc.log.Error("从服务器获取图片流失败")
				uc.log.Error(err)
				//从服务器获取图片流失败
				//存放原图到服务器
				diskbyte, _, err := imagehelper.GetImgFromUrlToBytes(url)
				if err == nil {
					//获取流没问题保存多次去服务器
					uc.repo.StoreImgWithFtp(ctx, ftpinfo, oname, diskbyte)
				} else {
					uc.log.Error(err)
				}
				obyte = diskbyte
			} else {
				obyte = diskbyte
			}
		}

		resizedimgbyte, filetype, err := imagehelper.ResizeImgToByteFromBytes(obyte, originimgcroplog.FileType, width)
		if err == nil {
			url, storepath, err := uc.repo.StoreImgWithFtp(ctx, ftpinfo, newname, resizedimgbyte)
			if err == nil {
				//保存图片到db
				uc.repo.CreateFileLog(ctx, &ImgCropLog{
					Source:    url,
					StorePath: storepath,
					Md5Url:    md5filename,
					Width:     uint64(width),
					FileType:  filetype,
				})
			}

		} else {
			uc.log.Error(err)
		}
		return &pb.CropImgReply{Imgdata: resizedimgbyte, Imgname: newname, Imagetype: filetype}, nil

	} else {
		// uc.log.Error("errerrerrerrerrerrerrerrerrerrerrerr")
		imgio, filetype, err := imagehelper.GetImgFromUrlToBytes(url)
		if err != nil {
			uc.log.Error(err)
			return nil, errors.New(500, "GETPICERROR", "图片错误")
		}

		if err == nil {
			_, ostorepath, err := uc.repo.StoreImgWithFtp(ctx, ftpinfo, oname, imgio)
			//上传成功则保存原图记录到数据库
			if err == nil {
				uc.repo.CreateFileLog(ctx, &ImgCropLog{
					Source:    url,
					StorePath: ostorepath,
					Md5Url:    md5filename,
					Width:     0,
					FileType:  filetype,
				})
			}
		} else {
			uc.log.Error(err)
		}
		if width <= 0 {
			return &pb.CropImgReply{Imgdata: imgio, Imgname: newname, Imagetype: filetype}, nil
		}
		//压缩图片
		resizedimgbyte, filetype, err := imagehelper.ResizeImgToByteFromBytes(imgio, filetype, width)
		if err == nil {
			url, storepath, err := uc.repo.StoreImgWithFtp(ctx, ftpinfo, newname, resizedimgbyte)
			if err == nil {
				//保存图片到db
				uc.repo.CreateFileLog(ctx, &ImgCropLog{
					Source:    url,
					StorePath: storepath,
					Md5Url:    md5filename,
					Width:     uint64(width),
					FileType:  filetype,
				})
			} else {

				uc.log.Error(err)
			}

		} else {

			uc.log.Error(err)
		}
		return &pb.CropImgReply{Imgdata: resizedimgbyte, Imgname: newname, Imagetype: filetype}, nil

	}

}
