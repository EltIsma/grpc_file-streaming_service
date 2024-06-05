package service

import (
	"grpc_serv/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type Service interface {
	GetStreamBytes(filename string)  (*models.FileData, error) 
	ListFiles() ([]string, error)
	GetFileInfo(filename string) (*models.FileInfo, error)
	UploadFile(string, []byte) error
}

type Usercase struct {
	files Service
}


func NewService(files Service) *Usercase {
	return &Usercase{
		files: files,
	}
}

func (u *Usercase) ListFiles() ([]string, error) {
	return u.files.ListFiles()
}
func (u *Usercase) GetFileInfo(name string) (*models.FileInfo, error) {
	return u.files.GetFileInfo(name)
}

func (u *Usercase) GetStreamBytes(filename string) (*models.FileData, error) {
	return u.files.GetStreamBytes(filename)
}

func (u *Usercase) UploadFile(filename string,content []byte) error{
	return u.files.UploadFile(filename, content)
}

// функция для проверки имени на корректность
func ValidateFileName(name string) error {
	if name == "" {
		return status.Error(codes.InvalidArgument, "Invalid filename")
	}
	return nil
}

// функция для проверки содержимого на корректность
func ValidateFileContent(content []byte) error {
	if len(content) == 0 {
		return status.Error(codes.InvalidArgument, "Invalid file content")
	}
	return nil
}