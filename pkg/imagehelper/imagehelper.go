package imagehelper

import (
	"bytes"
	"crypto/tls"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
)

func GetImgFromUrl(url string) (newimg image.Image, newimgtype string, er error) {
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
	return img, filetype, nil
}

func GetRemoteImgByte(url string) ([]byte, string, error) {
	img, filetype, err := GetImgFromUrl(url)
	if err != nil {
		return nil, "", err
	}
	by, err := TransferImgToByte(img, filetype)
	if err != nil {
		return nil, "", err
	}
	return by, filetype, err
}

func ResizeImg(img image.Image, width int64) (newimg image.Image) {

	thumbnailSize := int(width)
	var newImage image.Image

	newImage = resize.Resize(uint(thumbnailSize), 0, img, resize.Lanczos3)
	return newImage

}

func TransferImgToByte(newImage image.Image, filetype string) (bytedata []byte, err error) {
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
		return nil, errors.New(500, "Picture Invalid", "图片格式有误")
	}
	// fmt.Println("filetype", filetype)
	return w.Bytes(), nil
}

func ByteToImage(imgbyte []byte) (image.Image, string, error) {
	reader := bytes.NewReader(imgbyte)
	newimg, filetype, err := image.Decode(reader)
	if err != nil {

		log.Error(err)
		return nil, "", errors.New(500, "Picture Invalid", "图片格式有误")
	}
	return newimg, filetype, nil
}

func ResizeImgToByte(url string, width int64) ([]byte, string, error) {

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
