package file

import (
	"io"
	"os"
)

/*
File name    : file.go
Author       : miaoyc
Create date  : 2022/9/3 1:10 上午
Description  : 文件目录相关
*/

/*
FileExists
判断文件是否存在
如果文件存在返回true,否则返回false
*/
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// FileSize 获取文件大小
func FileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func DeleteFile(fileName string) {
	os.Remove(fileName)
}

// CopyFile 拷贝文件
func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm) //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 创建目录
func CreateDir(path string) error {
	exist, _ := PathExists(path)
	if !exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// BatchCreateDir 批量创建目录
func BatchCreateDir(pathList ...string) {
	for _, path := range pathList {
		CreateDir(path)
	}
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
