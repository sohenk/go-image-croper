package imagehelper

import (
	"bytes"
	"crypto/tls"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"imgcropper/pkg/resizegif"
	"io"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nfnt/resize"
)

func GetImgFromUrlToBytes(url string) (b []byte, filetype string, er error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	// retrieve a byte slice from bytes.Buffer
	data := buf.Bytes()
	_, filetype, err = image.Decode(bytes.NewBuffer(data))
	if err != nil {
		log.Error(err)
		return nil, "", errors.New(500, "image format error", "file format error")
	}

	return data, filetype, nil
}

//Turn ioreader to image
func IoReaderToImage(reader io.Reader) (image.Image, string, error) {
	newimg, filetype, err := image.Decode(reader)
	if err != nil {
		return nil, "", errors.New(500, "image format error", err.Error())
	}

	return newimg, filetype, nil
}

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

//ioreadertoBytes
func IoReaderToBytes(reader io.Reader) ([]byte, error) {
	buf := &bytes.Buffer{}

	buf.ReadFrom(reader)
	return buf.Bytes(), nil

}

//imagetobytes
func ImageToBytes(img image.Image, filetype string) ([]byte, error) {
	w := new(bytes.Buffer)
	switch filetype {
	case "png":
		err := png.Encode(w, img)
		if err != nil {
			return nil, err
		}
	case "gif":
		err := gif.Encode(w, img, nil)
		if err != nil {
			return nil, err
		}
	case "jpeg", "jpg":
		err := jpeg.Encode(w, img, nil)
		if err != nil {
			return nil, err
		}
	default:
		// not sure how you got here but what are we going to do with you?
		// fmt.Println("Unknown image type: ", filetype)
		return nil, errors.New(500, "Picture Invalid", "图片格式有误")
		//io.Copy(w, file)
	}
	// fmt.Println("filetype", filetype)
	return w.Bytes(), nil
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

//bytes to ioreader
func BytesToIoReader(img []byte) io.Reader {
	return bytes.NewReader(img)

}

//resize img from bytes to bytes
func ResizeImgToByteFromBytes(img []byte, filetype string, width int64) ([]byte, string, error) {
	reader := bytes.NewReader(img)

	newimg, filetype, err := image.Decode(reader)
	if err != nil {
		log.Error(err)
		return nil, "", errors.New(500, "Picture Invalid", "图片格式有误")
	}
	thumbnailSize := int(width)
	var newImage image.Image
	if filetype != "gif" {
		newImage = resize.Resize(uint(thumbnailSize), 0, newimg, resize.Lanczos3)
	}
	w := new(bytes.Buffer)
	switch filetype {
	case "png":
		err = png.Encode(w, newImage)
	case "gif":
		gifs, err := resizegif.Resize(BytesToIoReader(img), thumbnailSize, 0)
		if err != nil {
			log.Error(err)
			return nil, "", errors.New(500, "Picture Invalid", "图片转换有误")
		}
		err = gif.EncodeAll(w, gifs)
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
