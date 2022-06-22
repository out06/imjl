package Util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {

		return false
	}
	return s.IsDir()

}

func IsFile(path string) bool {
	return !IsDir(path)
}

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}
// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth string, suffix []string) (files []string, err error) {
	files = make([]string, 0, 30)

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		//    return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		for _, suf := range suffix {
			suf = strings.ToUpper(suf) //忽略后缀匹配的大小写
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suf) {
				files = append(files, filename)
			}
		}

		return nil
	})

	return files, err
}

func CopyFile() {
	lf := "/Users/imjl/Downloads/Obsidian-0.14.15-universal.dmg"
	rf := "/Users/imjl/Downloads/test.dmg"

	l, err := os.OpenFile(lf, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}

	r, err := os.OpenFile(rf, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
	}

	defer l.Close()
	defer r.Close()

	lbuffer := bufio.NewReader(l)
	rbuffer := bufio.NewWriter(r)

	n, err := io.Copy(rbuffer, lbuffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	err = rbuffer.Flush()
	if err != nil {
		fmt.Println(err)
	}
}