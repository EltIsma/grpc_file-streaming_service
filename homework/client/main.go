package main

import (
	_ "bufio"
	"context"
	"fmt"
	fs "grpc_serv/service/file"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := fs.NewFileStreamClient(conn)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Available Commands:")
		fmt.Println("list - Get the list of files")
		fmt.Println("info <file_name> - Get information about a specific file")
		fmt.Println("download <file_name> - Download a file from the server")
		fmt.Println("upload <file_path> - Upload a file to the server")
		return
	}

	command := args[1]
	switch command {
	case "list":
		listFiles(client)
	case "info":
		if len(args) < 3 {
			fmt.Println("Please provide a file name for 'info' command")
			return
		}
		fileName := args[2]
		getFileInfo(client, fileName)
	case "download":
		if len(args) < 3 {
			fmt.Println("Please provide a file name for 'download' command")
			return
		}
		fileName := args[2]
		downloadFile(client, fileName)
	case "upload":
		if len(args) < 3 {
			fmt.Println("Please provide a file path for 'upload' command")
			return
		}
		filePath := args[2]
		uploadFile(client, filePath)
	default:
		fmt.Println("Invalid command. Available commands: list, info, download, upload")
	}
}

func listFiles(client fs.FileStreamClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.ListFiles(ctx, &fs.Empty{})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("All files in our storage:")
	for _, name := range resp.Names {
		fmt.Println(name)
	}
}

func getFileInfo(client fs.FileStreamClient, fileName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	respInfo, err := client.GetFileInfo(ctx, &fs.FileName{Name: fileName})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("File Info: %+v\n", respInfo)
}

func downloadFile(client fs.FileStreamClient, fileName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := client.SendFileName(ctx, &fs.FileName{Name: "kuku.txt"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer f.Close()
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		if _, err := f.Write(chunk.Data); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
	fmt.Println("File downloaded successfully")
}

func uploadFile(client fs.FileStreamClient, filePath string) {
	filename := filepath.Base(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()
	stream1, err := client.UploadFile(ctx)
	if err != nil {
		log.Fatalf("Failed to start upload stream: %v", err)
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("could not read file: %v", err)
		return
	}
	err = stream1.Send(&fs.UploadFileRequest{
		Filename: filename,
		Content:  content,
	})
	if err != nil {
		log.Printf("could not send file content: %v", err)
		return
	}
	response, err := stream1.CloseAndRecv()
	if err != nil {
		fmt.Printf("could not receive response: %v", err)
		return
	}

	fmt.Printf("File %s uploaded successfully. Server message: %s\n", response.Filename, response.Message)

}
