package storage

import (
	"bytes"
	"grpc_serv/models"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)


type MyTestSuite struct {
	suite.Suite
	service Repository
	root    string
}

func (suite *MyTestSuite) SetupSuite() {
	suite.root = "./test_data"
	suite.service = &LocalFileRepository{Root: suite.root}
	file1, err := os.Create(filepath.Join(suite.root, "file1.txt"))
	if err != nil {
		panic(err)
	}
	defer file1.Close()
	_, err = file1.WriteString("0123456789")
	if err != nil {
		panic(err)
	}
	file2, err := os.Create(filepath.Join(suite.root, "file2.txt"))
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	_, err = file2.WriteString("file content")
	if err != nil {
		panic(err)
	}
	file3, err := os.Create(filepath.Join(suite.root, "file3.txt"))
	if err != nil {
		panic(err)
	}
	defer file3.Close()
}

func (suite *MyTestSuite) TestLocalFileRepository_ListFiles() {
	expectedFiles := []string{"file1.txt", "file2.txt", "file3.txt"}

	files, err := suite.service.ListFiles()

	if err != nil {
		suite.T().Errorf("Error occurred while listing files: %v", err)
	}
	if !equalSlice(files, expectedFiles) {
		suite.T().Errorf("Listed files did not match the expected files: %v", err)
	}
}

func (suite *MyTestSuite) TestLocalFileRepository_GetFileInfo() {

	expectedFileInfo := &models.FileInfo{Name: "file1.txt", Size: 10, FileType: ".txt"}

	fileInfo, err := suite.service.GetFileInfo("file1.txt")

	if err != nil {
		suite.T().Errorf("Error occurred while getting file info: %v", err)
	}
	if *fileInfo != *expectedFileInfo {
		suite.T().Errorf("Retrieved file info did not match the expected info: %v", err)
	}
}

func (suite *MyTestSuite) TestLocalFileRepository_GetStreamBytes() {
	_, err := suite.service.GetStreamBytes("nonexistent_file.txt")
	suite.Assert().Error(err)

	fileData := []byte("file content")
	err = os.WriteFile(filepath.Join(suite.root, "file2.txt"), fileData, 0644)
	suite.Assert().NoError(err)

	receivedFileData, err := suite.service.GetStreamBytes("file2.txt")
	suite.Assert().NoError(err)
	suite.Assert().Equal(fileData, receivedFileData.Data)
}

func (suite *MyTestSuite) TearDownSuite() {
	err := os.Remove(filepath.Join(suite.root, "file1.txt"))
	if err != nil {
		suite.T().Logf("Error occurred while removing files: %v", err)
	}
	err = os.Remove(filepath.Join(suite.root, "file2.txt"))
	if err != nil {
		suite.T().Logf("Error occurred while removing files: %v", err)
	}
	err = os.Remove(filepath.Join(suite.root, "file3.txt"))
	if err != nil {
		suite.T().Logf("Error occurred while removing files: %v", err)
	}
}

func (suite *MyTestSuite) TestUploadFile() {
    filename := "file3.txt"
    content := []byte("Hello, World!")
    err := suite.service.UploadFile(filename, content)
    if err != nil {
		suite.T().Logf("Failed to upload file: %v", err)
    }

    uploadedContent, err := os.ReadFile(filepath.Join(suite.root, filename))
    if err != nil {
		suite.T().Logf("Failed to read uploaded file: %v",err)
    }

    if !bytes.Equal(uploadedContent, content) {
		suite.T().Logf("Uploaded file content does not match expected content: expected %s, got %s", content, uploadedContent)
    }
}

func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}

func equalSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}


