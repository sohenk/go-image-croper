package biz

import (
	"context"
	"crypto/md5"
	"fmt"
	pb "imgcropper/api/imgcropper/service"
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
//ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
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
}

type CropImgUsecase struct {
	repo CropImgRepo
	log  *log.Helper
}

func NewCropImgUsecase(repo CropImgRepo, logger log.Logger) *CropImgUsecase {
	return &CropImgUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CropImgUsecase) CropImgBiz(ctx context.Context, url string, width int64) (*pb.CropImgReply, error) {
	//fmt.Println(url, width)

	md5filename := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	baseName := path.Base(url)
	ext := path.Ext(baseName)                 // 输出 .jpg
	name := strings.TrimSuffix(baseName, ext) // 输出 name
	newname := name + "_w" + "_ " + strconv.Itoa(int(width)) + ext

	oname := name + ext

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

			uc.log.Debug("没缩略图")
		} else {
			uc.log.Debug("有缩略图")
			//有对应缩略图
			storebyte, err := uc.repo.GetImgFromDisk(ctx, imgcroplog.StorePath)
			if err != nil {
				uc.log.Error(err)
			}
			if storebyte != nil {
				return &pb.CropImgReply{
					Imgdata: storebyte, Imgname: newname, Imagetype: imgcroplog.FileType,
				}, nil
			}
		}

		uc.log.Debug("开始创建缩略图")
		//原图是否在本地有副本
		originimgcroplog, err := uc.repo.GetFileSize(ctx, md5filename, 0)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				uc.log.Error(err)
				return nil, errors.New(500, "dberror", "数据库出错了")
			}
			//原图不存在创建原图
			newimg, filetype, err := imagehelper.GetImgFromUrl(url)
			if err != nil {
				return nil, err
			}
			newbyte, err := imagehelper.TransferImgToByte(newimg, filetype)
			obyte = newbyte
			if err == nil {
				_, ostorepath, err := uc.repo.StoreImg(ctx, oname, newbyte)
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
				}
			}

			uc.log.Debug("原图服务器不存在")
		} else {

			uc.log.Debug("原图服务器存在")
			diskbyte, err := uc.repo.GetImgFromDisk(ctx, originimgcroplog.StorePath)
			if err != nil {
				uc.log.Error("从服务器获取图片流失败")
				uc.log.Error(err)
				//从服务器获取图片流失败
				//存放原图到服务器
				newimg, filetype, err := imagehelper.GetImgFromUrl(url)
				if err != nil {
					return nil, err
				}
				diskbyte, err := imagehelper.TransferImgToByte(newimg, filetype)

				if err == nil {
					//获取流没问题保存多次去服务器
					uc.repo.StoreImg(ctx, oname, diskbyte)
				}
				obyte = diskbyte
			} else {
				obyte = diskbyte
			}
		}
		newimg, filetype, err := imagehelper.ByteToImage(obyte)
		if err != nil {
			return nil, err
		}
		//压缩图片
		resizedimg := imagehelper.ResizeImg(newimg, width)
		//上传压缩图片到服务器
		resizedimgbyte, err := imagehelper.TransferImgToByte(resizedimg, filetype)
		if err == nil {
			url, storepath, err := uc.repo.StoreImg(ctx, newname, resizedimgbyte)
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

		}
		return &pb.CropImgReply{Imgdata: resizedimgbyte, Imgname: newname, Imagetype: filetype}, nil

	} else {
		//存放原图到服务器
		newimg, filetype, err := imagehelper.GetImgFromUrl(url)
		if err != nil {
			return nil, err
		}
		newbyte, err := imagehelper.TransferImgToByte(newimg, filetype)
		if err == nil {
			_, ostorepath, err := uc.repo.StoreImg(ctx, oname, newbyte)
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
		}
		//压缩图片
		resizedimg := imagehelper.ResizeImg(newimg, width)
		//上传压缩图片到服务器
		resizedimgbyte, err := imagehelper.TransferImgToByte(resizedimg, filetype)
		if err == nil {
			url, storepath, err := uc.repo.StoreImg(ctx, newname, resizedimgbyte)
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

		}
		return &pb.CropImgReply{Imgdata: resizedimgbyte, Imgname: newname, Imagetype: filetype}, nil

	}

}
