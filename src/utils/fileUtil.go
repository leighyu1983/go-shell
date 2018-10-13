package util

import (
	"mime/multipart"
	"io"
	"os"
	"fmt"
)

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

func appendToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	
	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt([]byte(content), n)  
	defer f.Close()  
	return err
}

func CreatePathIfNotExist(path string) (error) {
	exist, err := PathExists(path)
	if(err != nil) {
		return err
	}

    if(exist) {
		return nil
	}

	return os.Mkdir(path, os.ModePerm)
}

func CreateFile(file multipart.File, path string, filename string) (error){

	CreatePathIfNotExist(path)

	out, _ := os.Create(path + filename);
	defer out.Close()

	fmt.Printf("[util.SaveFile] prepare to write into file '%s'\n", path + filename)
	_, err := io.Copy(out, file)
    if err != nil {
		//fmt.Println("=====ccccccccccccc=====" + path)
		//fmt.Println(fmt.Printf("%s", err))
		//fmt.Println("=====dddddddddddd=====" + path)
		panic(err)
	}
	fmt.Printf("[util.SaveFile] successfully write into file '%s'\n", path + filename)
	return nil
}