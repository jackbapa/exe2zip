package deal

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)


func Depress(zipfile , destPath string)  {
	//目标文件夹不存在则创建
	if _, err := os.Stat(destPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(destPath, 0755)
		}
	}
	reader , _ := zip.OpenReader(zipfile)
	defer reader.Close()

	for _, file := range reader.File{
		if file.FileInfo().IsDir(){
			os.MkdirAll(destPath+"\\"+file.Name,0777)
		}
	}
	for _, file := range reader.File{
		if ! file.FileInfo().IsDir(){
			filename := destPath + "\\" + file.Name
			fmt.Println(filename)
			src ,_:= file.Open()
			defer src.Close()
			w,_:=os.Create(filename)
			io.Copy(w,src)

	}





	}

}
