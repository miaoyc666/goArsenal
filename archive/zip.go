package archive

import (
	"os"

	"github.com/alexmullins/zip"
)

// ZipBytesWithPassword 将字节流压缩为加密的 ZIP 文件
//
//	fileBytes - 文件的字节流
//	zipPath   - 输出 zip 文件路径
//	fileName  - zip 内部文件名，如 "data.csv", "document.pdf" 等
//	password  - zip 密码
func ZipBytesWithPassword(fileBytes []byte, zipPath, fileName, password string) error {
	// 1. 创建 zip 文件
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 2. zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 3. 创建加密的 ZIP 文件项
	fw, err := zipWriter.Encrypt(fileName, password)
	if err != nil {
		return err
	}

	// 4. 将字节写入 ZIP
	_, err = fw.Write(fileBytes)
	return err
}
