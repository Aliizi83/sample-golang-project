package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
	"github.com/google/uuid"
)

type FileService struct {
	*BaseService[models.File, dto.CreateFile, dto.UpdateFile, dto.File]
}

func NewFileService(cfg *config.Config, fileRepository repositories.FileRepository) *FileService {
	return &FileService{
		BaseService: NewBaseService[models.File, dto.CreateFile, dto.UpdateFile, dto.File](cfg, fileRepository),
	}
}

func (s *FileService) Create(ctx context.Context, req dto.CreateFile) (dto.File, error) {
	var fileResponse dto.File
	fileName, err := s.uploadAndSaveFile(req.File, req.Directory)
	if err != nil {
		return fileResponse, err
	}

	req.Name = fileName
	req.MimeType = req.File.Header.Get("Content-Type")

	return s.BaseService.Create(ctx, req)
}

func (s *FileService) Delete(ctx context.Context, id int) error {
	file, err := s.GetById(ctx, id)
	if err != nil {
		s.logger.Error(err, logging.IO, logging.RemoveFile, "could not find file while deleting", nil)
		return err
	}

	err = os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))
	if err != nil {
		s.logger.Error(err, logging.IO, logging.RemoveFile, err.Error(), nil)
	}

	return s.BaseService.Delete(ctx, id)
}

func (s *FileService) uploadAndSaveFile(file *multipart.FileHeader, directory string) (string, error) {
	randFileName := uuid.New()

	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := file.Filename
	fileNameArr := strings.Split(fileName, ".")
	fileExt := fileNameArr[len(fileNameArr)-1]
	fileName = fmt.Sprintf("%s.%s", randFileName, fileExt)
	dst := fmt.Sprintf("%s/%s", directory, fileName)

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}
	return fileName, nil

}
