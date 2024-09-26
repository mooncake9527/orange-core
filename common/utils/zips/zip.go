package zips

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 解压
func Unzip(zipFile string, destDir string) ([]string, error) {
	zipReader, err := zip.OpenReader(zipFile)
	var paths []string
	if err != nil {
		return []string{}, err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		if strings.Index(f.Name, "..") > -1 {
			return []string{}, fmt.Errorf("%s 文件名不合法", f.Name)
		}
		fPath := filepath.Join(destDir, f.Name)
		paths = append(paths, fPath)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fPath, os.ModePerm)
		} else {
			_, err := doFile(fPath, f)
			if err != nil {
				return nil, err
			}
		}
	}
	return paths, nil
}

func doFile(fPath string, f *zip.File) ([]string, error) {
	if err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
		return []string{}, err
	}

	inFile, err := f.Open()
	if err != nil {
		return []string{}, err
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return []string{}, err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		return []string{}, err
	}
	return nil, nil
}
