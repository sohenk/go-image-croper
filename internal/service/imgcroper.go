package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"newkratos/internal/biz"

	pb "newkratos/api/imgcropper/service"
)

type ImgcroperService struct {
	pb.UnimplementedImgcroperServer
	uc  *biz.CropImgUsecase
	log *log.Helper
}

func NewImgcroperService(
	uc *biz.CropImgUsecase,
	logger log.Logger,
) *ImgcroperService {
	return &ImgcroperService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/imagecropper")),
	}
}

func (s *ImgcroperService) CropImg(ctx context.Context, req *pb.CropImgRequest) (*pb.CropImgReply, error) {
	if req.Url == "" {
		return nil, errors.Newf(400, "No Url", "Url must be not empty")
	}
	width := req.GetWidth()
	if width < 1 {
		width = 960
	}
	return s.uc.CropImgBiz(ctx, req.GetUrl(), width)
}
