package server

import (
	nethttp "net/http"
	pb "newkratos/api/imgcropper/service"
)

func responseEncoder(w nethttp.ResponseWriter, r *nethttp.Request, v interface{}) error {
	// 通过Request Header的Accept中提取出对应的编码器
	// 如果找不到则忽略报错，并使用默认json编码器
	//codec, _ := http.CodecForRequest(r, "Accept")

	data := v.(*pb.CropImgReply)

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型w.Header().Set("content-type", "application/octet-stream")
	w.Header().Set("content-type", "image/"+data.Imagetype)
	w.Header().Set("Content-Disposition", "filename="+data.Imgname)
	w.Write(data.Imgdata)
	return nil
}
