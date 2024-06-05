package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	_ "time"

	"grpc_serv/config"
	"grpc_serv/models"
	storage "grpc_serv/repositories"
	"grpc_serv/service"
	"grpc_serv/service/file"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fileServer struct {
	serv service.Usercase
	file.UnimplementedFileStreamServer
}
//ListFiles возвращает список всех файлов на сервере
func (fs *fileServer) ListFiles(ctx context.Context, in *file.Empty) (*file.FileList, error) {
	files, err := fs.serv.ListFiles()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot read directory: %v", err)
	}
	return &file.FileList{Names: files}, nil

}
// SendFileName получив имя файла от клиента, отправлет поток байтов клиенту
func (fs *fileServer) SendFileName(req *file.FileName, stream file.FileStream_SendFileNameServer) error {
	fn := models.FileName{Name: req.GetName()}
	err := service.ValidateFileName(fn.Name)
	if err != nil {
		return err
	}
	files, err := fs.serv.GetStreamBytes(fn.Name)
	if err != nil {
		return status.Errorf(codes.NotFound, "File not found: %v", err)
	}
	for _, chunk := range files.Data {
		if err := stream.Send(&file.FileByte{Data: []byte{chunk}}); err != nil {
			return err
		}
		//time.Sleep(1 * time.Second)
	}
	// if err := stream.Send(&file.FileByte{Data: files.Data}); err != nil {
	// 	return err
	// }

	return nil
}
//Возвращает информацию о файле
func (fs *fileServer) GetFileInfo(ctx context.Context, in *file.FileName) (*file.FileInfo, error) {
	fn := models.FileName{Name: in.GetName()}
	err := service.ValidateFileName(fn.Name)
	if err != nil {
		return nil, err
	}
	fileInfo, err := fs.serv.GetFileInfo(fn.Name)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "File not found: %v", err)
	}

	return &file.FileInfo{
		Name: fileInfo.Name,
		Size: fileInfo.Size,
		Type: fileInfo.FileType,
	}, nil

}

func (fs *fileServer) UploadFile(stream file.FileStream_UploadFileServer) error {
	fmt.Println("Receiving file...")

	var filename string
	var content []byte

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if filename == "" {
			filename = req.Filename
		}

		content = append(content, req.Content...)
	}
	err := service.ValidateFileName(filename) // Проверяем, что имя файла не пустое
	if err != nil {
		return nil
	}
	err = service.ValidateFileContent(content) // Проверяем, что содержимое файла файла не пустое
	if err != nil {
		return nil
	}

	err = fs.serv.UploadFile(filename, content)
	if err != nil {
		return status.Errorf(codes.Internal, "could not write file: %v", err)
	}

	return stream.SendAndClose(&file.UploadFileResponse{
		Filename: filename,
		Message:  "File uploaded successfully",
	})
}

func main() {
	cfg, err := config.Load("./config.yaml") // загружаем настройки для запуска сервера
	if err != nil {
		 log.Fatalf("failed to get config: %v", err)
	}

	lis, err := net.Listen(cfg.Network, ":" +cfg.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	repo := storage.NewFileService()
	service := service.NewService(repo)
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	file.RegisterFileStreamServer(s, &fileServer{serv: *service})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
