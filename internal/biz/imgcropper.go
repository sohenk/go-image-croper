package biz

import (
	"bytes"
	"context"
	"crypto/tls"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	pb "imgcropper/api/imgcropper/service"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nfnt/resize"
)

var (
// ErrUserNotFound is user not found.
//ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type CropImg struct {
	url   string
	width int64
}

type CropImgRepo interface {
	//CropImg(context.Context, *CropImg) (*CropImg, error)
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

	baseName := path.Base(url)
	ext := path.Ext(baseName)                 // 输出 .jpg
	name := strings.TrimSuffix(baseName, ext) // 输出 name
	newname := name + "_w" + strconv.Itoa(int(width)) + ext

	newimgByte, filetype, err := ResizeImg(url, width)
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}
	return &pb.CropImgReply{Imgdata: newimgByte, Imgname: newname, Imagetype: filetype}, nil
}
func ResizeImg(url string, width int64) ([]byte, string, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	//file, err := os.Open(url)
	//defer file.Close()
	if err != nil {
		log.Error(err)
		return nil, "", errors.New(404, "not found", "file not found")
	}
	img, filetype, err := image.Decode(resp.Body)
	if err != nil {

		log.Error(err)
		return nil, "", errors.New(500, "image format error", "file format error")
	}
	thumbnailSize := int(width)
	var newImage image.Image

	newImage = resize.Resize(uint(thumbnailSize), 0, img, resize.Lanczos3)
	w := new(bytes.Buffer)
	switch filetype {
	case "png":
		err = png.Encode(w, newImage)
	case "gif":
		err = gif.Encode(w, newImage, nil)
	case "jpeg", "jpg":
		err = jpeg.Encode(w, newImage, nil)
	//case "bmp":
	//	err = bmp.Encode(w, newImage)
	//case "tiff":
	//	err = tiff.Encode(w, newImage, nil)
	default:
		// not sure how you got here but what are we going to do with you?
		// fmt.Println("Unknown image type: ", filetype)
		err = errors.New(500, "Picture Invalid", "图片格式有误")
		//io.Copy(w, file)
	}
	if err != nil {

		log.Error(err)
		return nil, "", errors.New(500, "Picture Invalid", "图片格式有误")
	}
	// fmt.Println("filetype", filetype)
	return w.Bytes(), filetype, nil
}
