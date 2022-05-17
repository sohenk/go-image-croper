package service

import (
	"context"
	"imgcropper/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

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
		width = 960
	}
	return s.uc.CropImgBiz(ctx, req.GetUrl(), width)
}
