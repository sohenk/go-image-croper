package data

import (
	"gorm.io/gorm"
	"imgcropper/internal/conf"
	"imgcropper/pkg/filesystem/ftpdriver"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDataBase, NewFtpClient, NewImgCropRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	ftp *ftpdriver.FtpInfo
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, ftp *ftpdriver.FtpInfo) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  db,
		ftp: ftp,
		log: log.NewHelper(log.With(logger, "module", "data/imgcrop")),
	}, cleanup, nil
}
func NewFtpClient(c *conf.FileSystem) (*ftpdriver.FtpInfo, error) {
	filesytem := c.Ftp
	ftpinfo, err := ftpdriver.NewFtpInfo(filesytem.Host, filesytem.Username, filesytem.Password, filesytem.Root, filesytem.Url, uint(filesytem.Port), filesytem.Dir)
	if err != nil {
		return nil, err
	}
	return ftpinfo, nil
}

// NewDataBase 初始化数据库
func NewDataBase(c *conf.Data) (*gorm.DB, error) {
	dsn := c.Database.Source
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 设置连接池
	// 空闲
	sqlDb.SetMaxIdleConns(50)
	// 打开
	sqlDb.SetMaxOpenConns(100)
	// 超时
	sqlDb.SetConnMaxLifetime(time.Second * 30)
	err = DBAutoMigrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DBAutoMigrate 数据库模型自动迁移
func DBAutoMigrate(db *gorm.DB) error {
	//fmt.Println("自动迁移")
	// 在这里让GORM知道那些结构体是我们的数据模型，GORM将完成自动建表
	err := db.AutoMigrate(
		&ImgCropLog{},
	)
	if err != nil {
		return err
	}

	return nil
}
