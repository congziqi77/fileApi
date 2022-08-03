package main

import (
	"bufio"
	"io"
	"os"

	"github.com/congziqi77/fileApi/api/operatefile"
)

//写文件
func WriteFile(fileName, content string) error {
	file, err := checkFileExist(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, content)
	if err != nil {
		return err
	}
	return nil
}

//缓存写文件
func writeFileByBuf(filename string, content []byte) error {
	file, err := checkFileExist(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	//为这个文件创建一个bufferd writer
	bufferWriter := bufio.NewWriter(file)

	_, err = bufferWriter.Write(content)
	if err != nil {
		return err
	}
	bufferWriter.Flush()
	return nil
}

func checkFileExist(fileName string) (*os.File, error) {
	_, err := operatefile.FileInfo(fileName)
	if err != nil {
		return nil, err
	}
	var file *os.File
	if os.IsNotExist(err) {
		file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
	} else {
		file, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
	}
	return file, nil
}
