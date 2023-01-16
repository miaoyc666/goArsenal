package file

import (
	"bufio"
	"io"
	"os"
	"sync"
)

/*
File name    : file.go
Author       : miaoyc
Create date  : 2022/9/3 1:10 上午
Description  : 文件目录相关
*/

var (
	lock sync.Mutex
)

/*
IsExists
判断文件是否存在
如果文件存在返回true,否则返回false
*/
func IsExists(filePath string) bool {
	_, err := os.Stat(filePath) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
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

// GetSize 获取文件大小
func GetSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
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

// SaveFile 保存文件
func SaveFile(filePath, content string) error {
	lock.Lock()
	defer lock.Unlock()
	fileHandle, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	// NewWriter 默认缓冲区大小是 4096
	// 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
	buf := bufio.NewWriterSize(fileHandle, len(content))
	buf.WriteString(content)
	err = buf.Flush()
	if err != nil {
		return err
	}
	return nil
}

// DeleteFile 删除文件
func DeleteFile(filePath string) {
	os.Remove(filePath)
}
