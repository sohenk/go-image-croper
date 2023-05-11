package service

import (
	"context"
	"fmt"
	"imgcropper/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"

	pb "imgcropper/api/imgcropper/service"
)

type ImgcropperService struct {
	pb.UnimplementedImgcropperServer
	uc  *biz.CropImgUsecase
	log *log.Helper
}

func NewImgcropperService(
	uc *biz.CropImgUsecase,
	logger log.Logger,
) *ImgcropperService {
	return &ImgcropperService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/imagecropper")),
	}
}

func (s *ImgcropperService) CropImg(ctx context.Context, req *pb.CropImgRequest) (*pb.CropImgReply, error) {
	if req.Url == "" {
		return nil, errors.Newf(400, "No Url", "Url must be not empty")
	}
	width := req.GetWidth()
	if width < 1 {
		width = 0
	}
	sg := &singleflight.Group{}
	return s.SingleCropImg(ctx, sg, req)
}
func (s *ImgcropperService) SingleCropImg(ctx context.Context, sg *singleflight.Group, req *pb.CropImgRequest) (*pb.CropImgReply, error) {
	width := req.GetWidth()
	key := fmt.Sprintf("go_imagecroper:ByRouteDetail:url:%s:width:%d", req.Url, width)
	v, err, _ := sg.Do(key, func() (interface{}, error) {
		return s.uc.CropImgBiz(ctx, req.GetUrl(), width)
	})
	return v.(*pb.CropImgReply), err
}
