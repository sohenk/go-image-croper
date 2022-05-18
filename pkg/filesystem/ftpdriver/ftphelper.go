package ftpdriver

import (
	"bytes"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jlaffaye/ftp"
	"io/ioutil"
	"strconv"
	"time"
)

type FtpInfo struct {
	Host     string
	UserName string
	PassWord string
	Root     string
	Url      string
	Port     uint
	Dir      string
	Conn     *ftp.ServerConn
	//Connect  *ftp.ServerConn
	//Passive  bool
	//Ssl      bool
	//TimeOut  uint
}

func NewFtpInfo(Host string, UserName string, PassWord string, Root string, Url string, Port uint, Dir string) (*FtpInfo, error) {
	if Port == 0 {
		Port = 21
	}

	ftp := &FtpInfo{
		Host:     Host,
		UserName: UserName,
		PassWord: PassWord,
		Root:     Root,
		Url:      Url,
		Port:     Port,
		Dir:      Dir,
		//Passive:  Passive,
		//Ssl:      Ssl,
		//TimeOut:  TimeOut,
	}
	conn, err := ftp.Login()
	if err != nil {
		return nil, err
	}
	ftp.Conn = conn
	//ftp.Connect ,err =FtpLogin()
	return ftp, nil
}

func (f *FtpInfo) Login() (*ftp.ServerConn, error) {
	url := f.Host + ":" + strconv.Itoa(int(f.Port))
	c, err := ftp.Dial(url, ftp.DialWithTimeout(1*time.Second))
	if err != nil {
		return nil, err
	}

	err = c.Login(f.UserName, f.PassWord)
	if err != nil {
		return nil, err
	}

	return c, nil
	// Do something with the FTP conn

}
func (f *FtpInfo) Logout() error {
	f.Reconnet()
	if err := f.Conn.Quit(); err != nil {
		return err
	}
	return nil
}
func (f *FtpInfo) ReadFile(filepath string) ([]byte, error) {
	f.Reconnet()
	dir := f.Root + filepath
	r, err := f.Conn.Retr(dir)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
func (f *FtpInfo) Store(
	byteBuffer []byte,
//byteBuffer *bytes.Buffer,
	fileName string) (url, storePath string, err error) {

	f.Reconnet()
	err = f.EnsureFtpDirExist(f.Dir)
	if err != nil {
		return "", "", err
	}
	storePath = f.Dir + "/" + fileName
	reader := bytes.NewReader(byteBuffer)
	err = f.Conn.Stor(f.Root+storePath, reader)
	if err != nil {
		return "", "", err
	}
	url = f.Url + "/" + f.Dir + "/" + fileName
	return url, storePath, nil
}

func (f *FtpInfo) EnsureFtpDirExist(dir string) error {
	f.Reconnet()
	err := f.Conn.ChangeDir(f.Root + dir)
	if err != nil {
		er := f.Conn.MakeDir(f.Root + dir)

		if er != nil {
			return er
		}

		return nil
	}

	return nil

}

func (f *FtpInfo) Reconnet() {
	if err := f.Conn.NoOp(); err != nil {
		log.Error(err)
		log.Warn("ftp重新连接")
		conn, err := f.Login()
		if err == nil {
			f.Conn = conn
		}
	}
}
