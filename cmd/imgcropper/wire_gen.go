// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"imgcropper/internal/biz"
	"imgcropper/internal/conf"
	"imgcropper/internal/data"
	"imgcropper/internal/server"
	"imgcropper/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, fileSystem *conf.FileSystem, logger log.Logger) (*kratos.App, func(), error) {
	db, err := data.NewDataBase(confData)
	if err != nil {
		return nil, nil, err
	}
	ftpInfo, err := data.NewFtpClient(fileSystem)
	if err != nil {
		return nil, nil, err
	}
	client := data.NewRedisClient(confData, logger)
	dataData, cleanup, err := data.NewData(confData, logger, db, ftpInfo, client)
	if err != nil {
		return nil, nil, err
	}
	cropImgRepo := data.NewImgCropRepo(dataData, logger)
	cropImgUsecase := biz.NewCropImgUsecase(cropImgRepo, fileSystem, logger)
	imgcropperService := service.NewImgcropperService(cropImgUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, imgcropperService, logger)
	grpcServer := server.NewGRPCServer(confServer, imgcropperService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
