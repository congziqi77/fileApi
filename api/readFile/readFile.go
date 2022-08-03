package main

import (
	"os"
)

//读取整个文件

func ReadAllFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}



