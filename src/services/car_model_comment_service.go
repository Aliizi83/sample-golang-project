package services

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

func NewCarModelCommentService(cfg *config.Config, repository repositories.CarModelCommentRepository) *CarModelCommentService {
	return &CarModelCommentService{
		BaseService: NewBaseService[models.CarModelComment, dto.CreateCarModelComment, dto.UpdateCarModelComment, dto.CarModelComment](cfg, repository),
	}
}

type CarModelCommentService struct {
	*BaseService[models.CarModelComment, dto.CreateCarModelComment, dto.UpdateCarModelComment, dto.CarModelComment]
}

func (s *CarModelCommentService) Create(ctx context.Context, req dto.CreateCarModelComment) (dto.CarModelComment, error) {
	req.UserId = s.GetAuthUserId(ctx)
	return s.BaseService.Create(ctx, req)
}
