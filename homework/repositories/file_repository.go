package storage

import (
	"grpc_serv/models"
	"os"
	"path/filepath"
)

type Repository interface {
	GetStreamBytes(filename string) (*models.FileData, error)
	ListFiles() ([]string, error)
	GetFileInfo(filename string) (*models.FileInfo, error)
	UploadFile(string, []byte) error
}

type FileRepository struct {
	Repository
}

func NewFileService() *FileRepository {
	return &FileRepository{
		Repository: NewFileRepository(),
	}
}

type LocalFileRepository struct {
	Root string
}

func NewFileRepository() *LocalFileRepository {
	return &LocalFileRepository{Root: "./repositories/storage"}//устанавливаем каталог для хранения файлов
}
// получаем список имён файлов в нашем каталоге
func (repo *LocalFileRepository) ListFiles() ([]string, error) {
	files, err := os.ReadDir(repo.Root)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, nil
}
// получаем информацию о файле
func (repo *LocalFileRepository) GetFileInfo(name string) (*models.FileInfo, error) {
	file, err := os.Open(filepath.Join(repo.Root, name))
	if err != nil {
		return &models.FileInfo{}, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return &models.FileInfo{}, err
	}

	return &models.FileInfo{Name: name, Size: fileInfo.Size(), FileType: filepath.Ext(name)}, nil
}

// считываем содержимое файла
func (repo *LocalFileRepository) GetStreamBytes(filename string) (*models.FileData, error) {
	filePath := filepath.Join(repo.Root, filename)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &models.FileData{Data: data}, nil
}

// загружаем файл в наш каталог
func (repo *LocalFileRepository) UploadFile(filename string, content []byte) error {
	filePath := filepath.Join(repo.Root, filename)
	_ = os.WriteFile(filePath, content, 0644)
	return nil
}
