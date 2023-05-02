package server

import (
	"bytes"
	"crypto/tls"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fogleman/gg"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

func responseErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	// 拿到error并转换成kratos Error实体
	se := errors.FromError(err)
	// 通过Request Header的Accept中提取出对应的编码器
	// body, err := codec.Marshal(se)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	wb := getDefaultRemoteErrorPic()
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型w.Header().Set("content-type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "filename=error.jpg")
	w.Header().Set("content-type", "image/jpeg")
	// 设置HTTP Status Code
	w.WriteHeader(int(se.Code))
	// log.Info("r", r)
	w.Write(wb.Bytes())

	// data := v.(*pb.CropImgReply)

	// w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型w.Header().Set("content-type", "application/octet-stream")
	// w.Header().Set("content-type", "image/"+data.Imagetype)
	// w.Header().Set("Content-Disposition", "filename="+data.Imgname)
	// w.Write(data.Imgdata)
	// return nil

}
func drawNonePic() *bytes.Buffer {
	wb := new(bytes.Buffer)
	// m := image.NewRGBA(image.Rect(0, 0, 600, 400))
	// blue := color.RGBA{230, 230, 230, 1}
	// draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	// draw.Dra
	dc := gg.NewContext(600, 400)
	dc.SetRGB(128, 128, 128)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	// dc.Clear()
	// if err := dc.LoadFontFace("方正楷体简体.ttf", 120); err != nil {
	// 	panic(err)
	// }
	dc.DrawPoint(100, 100, 100)
	dc.DrawString("Error ! Picture Error.", 230, 400/2)
	// fmt.Println(se.Message)

	jpeg.Encode(wb, dc.Image(), nil)
	return wb
}
func getDefaultLocalErrorPic() *bytes.Buffer {
	path := "C:\\Users\\DELL\\Desktop\\nopic.png"
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 6) // 读写方式打开
	if err != nil {
		// 如果有错误返回错误内容
		log.Error("打不开图片")
		return drawNonePic()
	}
	defer fp.Close()
	bs, err := ioutil.ReadAll(fp)
	if err != nil {

		log.Error("读图片有错")
		return drawNonePic()
	}
	reader := bytes.NewBuffer(bs)

	return reader
}

func getDefaultRemoteErrorPic() *bytes.Buffer {
	path := "https://res.zhrct.com/nopic.png"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(path)
	if err != nil {
		return drawNonePic()
	}
	defer resp.Body.Close()
	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	return buf
}
