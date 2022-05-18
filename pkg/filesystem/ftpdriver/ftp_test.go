package ftpdriver

import (
	"testing"
)

func TestFtpInfo_ReadFile(t *testing.T) {
	ftp, err := NewFtpInfo("172.19.80.7", "res", "kaifa9394", "/home/res/", "https://res.zhrct.com/", 21, "imgcropper")

	if err != nil {
		t.Error(err)
		return
	}
	defer ftp.Logout()
	//读文件

	//filepath:="/home/res/wechatrct/wechat/8/addon/chargefunction/images/b8f5db1d7c999e3c05f91cd9fa28035d.jpg"
	//file, err := ftp.ReadFile(c, filepath)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fialename := "fsdfs.txt"
	//
	//data := []byte("Hello AlwaysBeta")
	//reader := bytes.NewReader(data)
	//err = ftp.Conn.Stor(ftp.Dir+"/"+fialename, reader)
	//if err != nil {
	//
	//	//return  err
	//}
	//t.Log(err)

	byt, err := ftp.ReadFile("imgcropper/test.txt")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(byt)

	t.Log(string(byt))
	//url, fp, err := ftp.Store(data, fialename)
	//
	//t.Log("文件路径", ftp.Conn, url, fp)
	////t.Log(createfilepath)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
}
