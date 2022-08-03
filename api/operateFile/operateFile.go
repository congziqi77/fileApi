package operatefile

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
)

//文件信息 如果文件不存在返回文件不存在信息

func FileInfo(filepath string) (fs.FileInfo, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("file does not exist")
		}
		return nil, err
	}
	return fileInfo, nil
}

//删除文件
func FileDel(filepath string) error {
	err := os.Remove(filepath)
	return err
}

//复制文件

func FileCopy(filepath string, newFileName string) (string, error) {
	originalFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	newFile, err := os.Create(newFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
	return newFileName, nil
}
