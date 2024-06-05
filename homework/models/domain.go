package models

//доменная модель: представляет информацию о файле
type FileInfo struct {
	Name string
	Size int64
	FileType string
}

//доменная модель: представляет содержимое файла
type FileData struct {
	Data []byte
}

//доменная модель: представляет имя файла
type FileName struct {
	Name string
}